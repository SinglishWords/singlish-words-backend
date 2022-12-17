package service

import (
	"singlishwords/dao"
	"singlishwords/log"
	"singlishwords/model"
)

var associationDAO = dao.AssociationDAO{}

func marshal(set map[string]int, associations []model.Association) (*Visualisation, error) {
	nodes := make([]model.Node, len(associations))
	links := make([]model.Link, len(associations))
	ids := make(map[string]int64)

	var i int64 = 0
	for word := range set {
		nodes = append(nodes, model.Node{Id: i, Name: word, SymbolSize: 0, Value: 0, Category: 0})
		ids[word] = i
		i++
	}

	for _, association := range associations {
		links = append(links, model.Link{Source: ids[association.Source], Target: ids[association.Target]})
	}

	return &Visualisation{Nodes: nodes, Links: links, Categories: []model.Category{}}, nil
}

func createSetAndNeighbors(associations []model.Association) (map[string]int, []string) {
	m := make(map[string]int)
	n := make([]string, len(associations)-1)

	for _, a := range associations {
		n = append(n, a.Target)

		_, ok := m[a.Source]
		if !ok {
			m[a.Source] = 1
		}

		_, ok = m[a.Target]
		if !ok {
			m[a.Target] = 1
		}
	}

	return m, n
}

type Visualisation struct {
	Nodes 		[]model.Node `json:"nodes"`
	Links 		[]model.Link `json: "links"`
	Categories 	[]model.Category `json: "categories"`
}

func GetForwardAssociations(word string) (*Visualisation, error) {
	// Get set of words: queried word, and all 1-away neighbors of the queried word
	associations, err := associationDAO.GetAssociationsBySource(word)
	if err != nil {
		return nil, err
	}

	set, neighbors := createSetAndNeighbors(associations)
	log.Logger.Infof("Set of words: %+v", set)
	log.Logger.Infof("First-degree neighbors of '%s': %+v", word, set)

	// Get all associations where source in [...neighbors]
	neighborsAssociations, err := associationDAO.MultiSelectBySource(neighbors)
	if err != nil {
		return nil, err
	}

	validNeighborsAssociations := make([]model.Association, 0, len(neighborsAssociations))
	for _, association := range neighborsAssociations {
		_, ok := set[association.Target]
		if ok {
			validNeighborsAssociations = append(validNeighborsAssociations, association)
		}
	}

	allAssociations := append(associations, validNeighborsAssociations...)
	return marshal(set, allAssociations)
}

func IncrementAssociationCount(q string, associatedWord string, inc int64) error {
	return associationDAO.IncrementAssociationBy(q, associatedWord, inc)
}
