package service

import (
	"singlishwords/dao"
	"singlishwords/model"
)

var communityDAO = dao.CommunityDAO{}

func find(s []model.Category, c int64) int {
	for i, e := range s {
		if e.Name == c {
			return i
		}
	}
	return -1
}

// Returns a list of unique communities that the input words are part of
// and a map of {word -> index of its community in the communities list}
func CreateCommunityListAndMap(words []string) ([]model.Category, map[string]int, error) {
	c := make([]model.Category, 0, 10) // 10 is arbitrary.
	m := make(map[string]int)

	communityMappings, err := communityDAO.MultiSelectByWord(words)
	if err != nil {
		return nil, nil, err
	}

	for _, cm := range communityMappings {
		idx := find(c, cm.Community)
		if idx == -1 {
			c = append(c, model.Category{Name: cm.Community})
			idx = len(c) - 1
		}
		
		m[cm.Word] = idx
	}

	return c, m, nil
}
