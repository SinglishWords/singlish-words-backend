package cache

import (
	"encoding/json"
	"math/rand"
	"singlishwords/dao"
	"singlishwords/log"
	"singlishwords/model"
	"strings"
)

type QuestionCache struct{}

var questionDAO dao.QuestionDAO

func (cache QuestionCache) GetNRandomQuestions(limit int) ([]model.Question, error) {
	questions, err := cache.getNRandomQuestionsFromRedis(limit)
	if err == nil {
		return questions, nil
	}

	log.Logger.Warn("Cache get random questions, miss!")
	// cache miss...
	questions, err = questionDAO.GetAll()
	if err != nil {
		log.Logger.Error("Mysql get random questions fail.")
		return nil, err
	}

	err = cache.storeAllToRedis(questions)
	if err != nil {
		log.Logger.Warn("Store to redis cache error.")
		return nil, err
	}

	return randomPickNQuestions(questions, limit), nil
}

func (QuestionCache) storeAllToRedis(questions []model.Question) error {
	if rdb == nil {
		return notConnectedError{}
	}
	pipe := rdb.Pipeline()
	for _, question := range questions {
		pipe.SAdd("question", question)
	}
	_, err := pipe.Exec()

	if err != nil {
		return err
	}

	return nil
}

func (QuestionCache) getNRandomQuestionsFromRedis(limit int) ([]model.Question, error) {
	if rdb == nil {
		return nil, notConnectedError{}
	}
	results, err := rdb.SRandMemberN("question", int64(limit)).Result()
	if err != nil || len(results) == 0 {
		return nil, cacheMissError{}
	}

	jsonString := "[" + strings.Join(results, ", ") + "]"

	questions := make([]model.Question, len(results))
	err = json.Unmarshal([]byte(jsonString), &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

// randomPickNQuestions
// Shuffle Pick,
// Randomly Pick (at most) n unique values from the given slice
// Time complexity: O(n), where n is the parameter n
// Space complexity: O(1), inplace modify the slices
//
// Important note:
// This function will shuffle the slice.
// But it won't totally shuffle the slice.
// It only pick n random entries, so at most shuffle n*2 entries.
func randomPickNQuestions(questions []model.Question, n int) []model.Question {
	length := len(questions)
	if n > length {
		n = length
	}

	for n > 0 {
		r := rand.Intn(length)
		questions[r], questions[length-1] = questions[length-1], questions[r]
		length--
		n--
	}

	return questions[length:]
}
