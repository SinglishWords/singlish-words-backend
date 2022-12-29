package service

import (
	"singlishwords/dao"
	"singlishwords/log"
	"singlishwords/model"
)

// Queried node have the largest symbolSize for visualization purposes
const queriedNodeSymbolSize = 10000000000

var associationDAO = dao.AssociationDAO{}

func marshalForwardAssociations(set map[string]int, associations []model.Association, queriedWord string) (*Visualisation, error) {
	nodes := make([]model.Node, 0, len(associations))
	links := make([]model.Link, 0, len(associations))
	ids := make(map[string]int64)
	
	words := make([]string, len(set))
	var i int64 = 0
	for word := range set {
		words[i] = word
		i++
	}

	// Given a word x, the value of x is the number of times the forward association ... -> x
	// is thought of when the word x is given
	associationsValue, err := associationDAO.CountForwardAssociations(words)
	if err != nil {
		return nil, err
	}

	// Temporary dirty fix to include the queried word if it is not included in associationDAO.CountForwardAssociations(words)
	found := false
	for _, av := range associationsValue {
		if av.Word == queriedWord {
			found = true
			break
		}
	}
	if !found {
		associationsValue = append(associationsValue, model.AssociationValue{ Word: queriedWord, Count: 0 })
	}

	i = 0
	for _, av := range associationsValue {
		word, value := av.Word, av.Count
		symbolSize := value
		if word == queriedWord {
			symbolSize = queriedNodeSymbolSize
		}

		nodes = append(nodes, model.Node{Id: i, Name: word, SymbolSize: symbolSize, Value: value, Category: 0})
		ids[word] = i
		i++
	}

	for _, association := range associations {
		links = append(links, model.Link{Source: ids[association.Source], Target: ids[association.Target]})
	}

	return &Visualisation{Nodes: nodes, Links: links, Categories: []model.Category{}}, nil
}

func marshalBackwardAssociations(set map[string]int, associations []model.Association, queriedWord string) (*Visualisation, error) {
	nodes := make([]model.Node, 0, len(associations))
	links := make([]model.Link, 0, len(associations))
	ids := make(map[string]int64)

	words := make([]string, len(set))
	var i int64 = 0
	for word := range set {
		words[i] = word
		i++
	}

	// Given a word x, the value of x is the number of times the forward association x -> ...
	// comes up. 
	associationsValue, err := associationDAO.CountBackwardAssociations(words)
	if err != nil {
		return nil, err
	}

	// Temporary dirty fix to include the queried word if it is not included in associationDAO.CountForwardAssociations(words)
	found := false
	for _, av := range associationsValue {
		if av.Word == queriedWord {
			found = true
			break
		}
	}
	if !found {
		associationsValue = append(associationsValue, model.AssociationValue{ Word: queriedWord, Count: 0 })
	}

	i = 0
	for _, av := range associationsValue {
		word, value := av.Word, av.Count
		symbolSize := value
		if word == queriedWord {
			symbolSize = queriedNodeSymbolSize
		}

		nodes = append(nodes, model.Node{Id: i, Name: word, SymbolSize: symbolSize, Value: value, Category: 0})
		ids[word] = i
		i++
	}

	for _, association := range associations {
		// Backward link {target -> source}
		links = append(links, model.Link{Source: ids[association.Target], Target: ids[association.Source]})
	}

	return &Visualisation{Nodes: nodes, Links: links, Categories: []model.Category{}}, nil
}

func createSetAndNeighbors(associations []model.Association) (map[string]int, []string) {
	m := make(map[string]int)
	n := make([]string, 0, len(associations))

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

func createSetAndBackwardNeighbors(associations []model.Association) (map[string]int, []string) {
	m := make(map[string]int)
	n := make([]string, 0, len(associations))

	for _, a := range associations {
		n = append(n, a.Source)

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
	Links 		[]model.Link `json:"links"`
	Categories 	[]model.Category `json:"categories"`
}

// Get all associations where source in [...neighbors]
func getNeighborsForwardAssociations(set map[string]int, neighbors []string) ([]model.Association, error) {
	neighborsAssociations, err := associationDAO.MultiSelectBySource(neighbors)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	validNeighborsAssociations := make([]model.Association, 0, len(neighborsAssociations))
	for _, association := range neighborsAssociations {
		// Is valid association iff target is in the set of nodes
		_, ok := set[association.Target]
		if ok {
			validNeighborsAssociations = append(validNeighborsAssociations, association)
		}
	}

	return validNeighborsAssociations, nil
}

func GetForwardAssociationsVisualisation(word string) (*Visualisation, error) {
	set, allAssociations, err := GetSetAndForwardAssociations(word)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return marshalForwardAssociations(set, allAssociations, word)
}

func GetSetAndForwardAssociations(word string) (map[string]int, []model.Association, error) {
	// Get set of words: queried word, and all 1-away neighbors of the queried word
	associations, err := associationDAO.GetAssociationsBySource(word)
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}

	set, neighbors := createSetAndNeighbors(associations)
	log.Logger.Infof("Set of words: %+v", set)
	log.Logger.Infof("First-degree neighbors of '%s': %+v", word, neighbors)

	neighborsAssociations, err := getNeighborsForwardAssociations(set, neighbors)
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}
	allAssociations := append(associations, neighborsAssociations...)

	return set, allAssociations, nil
}

// Get all associations where target in [...backwardNeighbors]
func getNeighborsBackwardAssociations(set map[string]int, neighbors []string) ([]model.Association, error) {
	neighborsAssociations, err := associationDAO.MultiSelectByTarget(neighbors)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	validNeighborsAssociations := make([]model.Association, 0, len(neighborsAssociations))
	for _, association := range neighborsAssociations {
		_, ok := set[association.Source]
		if ok {
			validNeighborsAssociations = append(validNeighborsAssociations, association)
		}
	}

	return validNeighborsAssociations, nil
}

func GetBackwardAssociationsVisualisation(word string) (*Visualisation, error) {
	set, allAssociations, err := GetSetAndBackwardAssociations(word)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return marshalBackwardAssociations(set, allAssociations, word)
}

func GetSetAndBackwardAssociations(word string) (map[string]int, []model.Association, error) {
	// Get set of words: queried word, and all 1-away backward neighbors of the queried word
	associations, err := associationDAO.GetBackwardAssociationsBySource(word)
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}

	set, backwardNeighbors := createSetAndBackwardNeighbors(associations)
	log.Logger.Infof("Set of words: %+v", set)
	log.Logger.Infof("First-degree backward neighbors of '%s': %+v", word, backwardNeighbors)

	neighborsAssociations, err := getNeighborsBackwardAssociations(set, backwardNeighbors)
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}
	allAssociations := append(associations, neighborsAssociations...)

	return set, allAssociations, nil
}

func IncrementAssociationCount(q string, associatedWord string, inc int64) error {
	return associationDAO.IncrementAssociationBy(q, associatedWord, inc)
}
