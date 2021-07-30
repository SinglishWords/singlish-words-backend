package cache

import (
	"encoding/json"
	"math/rand"
	"singlishwords/dao"
	"singlishwords/log"
	"singlishwords/model"
	"strings"
	"time"
)

type QuestionCache struct{}

var questionDAO dao.QuestionDAO

func (cache QuestionCache) GetNRandomQuestions(limit int) ([]model.Question, error) {
	questions, err := cache.getNextRangeQuestionsFromRedis(limit)
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

	//for _, question := range questions {
	//	pipe.SAdd("question", question)
	//}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	for _, question := range questions {
		pipe.LPush("questionList", question)
	}
	_, err := pipe.Exec()

	if err != nil {
		return err
	}

	return nil
}

func (QuestionCache) getNextRangeQuestionsFromRedis(limit int) ([]model.Question, error) {
	if rdb == nil {
		return nil, notConnectedError{}
	}

	pipe := rdb.TxPipeline()
	rIncr := rdb.IncrBy("questionIndex", int64(limit))
	rSize := rdb.LLen("questionList")
	_, err := pipe.Exec()

	if err != nil {
		return nil, cacheMissError{}
	}

	indexE, size := rIncr.Val(), rSize.Val()

	if size == 0 {
		return nil, cacheMissError{}
	}

	indexS := (indexE - int64(limit)) % size
	indexE = indexE % size
	// indexS:indexE

	var results []string
	if indexS <= indexE { // means in normal situation
		results, err = rdb.LRange("questionList", indexS, indexE).Result()
		if err != nil {
			return nil, err
		}
	} else { // means over the end, and then come back to the start
		pipe := rdb.TxPipeline()
		rRange1 := rdb.LRange("questionList", indexE, size)
		rRange2 := rdb.LRange("questionList", 0, indexS)
		_, err := pipe.Exec()

		if err != nil {
			return nil, err
		}
		results = append(rRange1.Val(), rRange2.Val()...)
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
