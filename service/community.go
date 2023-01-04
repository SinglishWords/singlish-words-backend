package service

import (
	"singlishwords/dao"
)

var communityDAO = dao.CommunityDAO{}

func CreateCommunityMapAndSet(words []string) (map[string]int64, map[int64]int, error) {
	m := make(map[string]int64)
	c := make(map[int64]int)

	communityMappings, err := communityDAO.MultiSelectByWord(words)
	if err != nil {
		return nil, nil, err
	}

	for _, cm := range communityMappings {
		m[cm.Word] = cm.Community

		_, ok := c[cm.Community]
		if !ok {
			c[cm.Community] = 1
		}
	}

	return m, c, nil
}
