package main

import (
	"fmt"
	"math"
	"singlishwords/model"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/community"
	"gonum.org/v1/gonum/graph/simple"
)

func createAssociationsGraph() (*simple.WeightedDirectedGraph, map[int64]string, error) {
    absent := math.Inf(1)       // the weight returned for absent edges
	g := simple.NewWeightedDirectedGraph(0, absent)

	words, err := associationDAO.GetUniqueWords()
	if err != nil {
		return nil, nil, err
	}

	// Add all nodes, create hash table that maps word to ID
	wordToId := make(map[string]int64)
	idToWord := make(map[int64]string)
	for _, word := range words {
		node := g.NewNode()
		id := node.ID()
		wordToId[word] = id
		idToWord[id] = word
		g.AddNode(node)
	}

	associations, err := associationDAO.GetAll()
	if err != nil {
		return nil, nil, err
	}

	// Add edges
	var source, target string
	var sourceId, targetId int64
	var sourceNode, targetNode graph.Node
	var ok bool
	for _, association := range associations {
		source, target = association.Source, association.Target
		if source == target {
			continue
		}

		sourceId, ok = wordToId[source]
		if !ok {
			return nil, nil, err
		}
		targetId, ok = wordToId[target]
		if !ok {
			return nil, nil, err
		}
		sourceNode, targetNode = g.Node(sourceId), g.Node(targetId)
		if sourceNode == nil || targetNode == nil {
			return nil, nil, err
		}
		e := g.NewWeightedEdge(sourceNode, targetNode, float64(association.Count))
		g.SetWeightedEdge(e)
	}

	return g, idToWord, nil
}

func runLouvain(g graph.Graph) [][]graph.Node {
	var resolution float64 = 1 // resolution heuristics: 
	return community.Modularize(g, resolution, nil).Expanded().Communities()
}

func printCommunities() error {
	communityMappings, err := communityDAO.GetAll()
	if err != nil {
		return err
	}

	var cur int64 = 0
	for _, cm := range communityMappings {
		if cm.Community != cur {
			fmt.Printf("\n---------------------------\n")
			fmt.Printf("Community %d: ", cm.Community)
			cur = cm.Community
		}
		fmt.Printf("%s, ", cm.Word)
	}

	return nil
}

func storeCommunityMappings(communities [][]graph.Node, idToWord map[int64]string) {
	for i, community := range communities {
		for _, node := range community {
			word, ok := idToWord[node.ID()]
			if !ok {
				return
			}
			cm := model.CommunityMapping{Word: word, Community: int64(i)}
			communityDAO.Upsert(&cm)
		}
	}
}

func detectCategories() error {
	g, idToWord, err := createAssociationsGraph()
	if err != nil {
		return err
	}

	communities := runLouvain(g)
	storeCommunityMappings(communities, idToWord)

	err = printCommunities()
	if err != nil {
		return err
	}

	return nil
}