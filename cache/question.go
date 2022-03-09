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
	// questions, err := cache.getNextRangeQuestionsFromRedis(limit)
	// if err == nil {
	// 	return questions, nil
	// }

	// log.Logger.Warn("Cache get random questions, cache miss!")

	// cache miss...
	log.Logger.Infof("Attempting to get weighted questions directly from db")
	questions, err := questionDAO.GetWeightedQuestions(limit)
	if err != nil {
		log.Logger.Error("Mysql get random questions fail.")
		return nil, err
	}

	for _, s := range questions {
		err := questionDAO.UpdateCount(&s)
		if err != nil {
			log.Logger.Warnf("Error updating count")
		}
	}

	// err = cache.storeAllToRedis(questions)
	// if err != nil {
	// 	log.Logger.Warn("Store to redis cache error.")
	// 	return nil, err
	// }

	// return randomPickNQuestions(questions, limit), nil
	return questions, err
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

	log.Logger.Infof("Get %d from mysql and store to the redis.", len(questions))
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
	log.Logger.Infof("Getting questions from redis, question base: %d, size %d", indexE, size)

	if size == 0 {
		return nil, cacheMissError{}
	}

	indexS := (indexE - int64(limit-1)) % size
	indexE = indexE % size
	// indexS:indexE

	var results []string
	if size < int64(limit) {
		indexS = 0
		indexE = size - 1
		log.Logger.Warnf("Get questions from redis, need oversized questions, so return all from %d to %d", indexS, indexE)
	}

	if indexS <= indexE { // means in normal situation
		results, err = rdb.LRange("questionList", indexS, indexE).Result()
		log.Logger.Infof("Get questions from redis, question range: %d to %d", indexS, indexE)
		if err != nil {
			log.Logger.Error("Wrong when getting questions from redis")
			return nil, err
		}
	} else { // means over the end, and then come back to the start
		pipe := rdb.TxPipeline()
		rRange1 := rdb.LRange("questionList", indexS, size-1)
		rRange2 := rdb.LRange("questionList", 0, indexE)
		_, err := pipe.Exec()

		log.Logger.Infof("Get questions from redis, question range: %d to %d and %d to %d",
			indexS, size-1, 0, indexE)

		if err != nil {
			log.Logger.Error("Wrong when getting questions from redis")
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

	maxi, err := questionDAO.GetMaxCount()
	if err != nil {
		log.Logger.Error("Error getting max count")
	}

	for n > 0 {
		r := getRandomWeightedIndex(questions[:length], maxi)
		questions[r], questions[length-1] = questions[length-1], questions[r]
		length--
		n--
	}

	a := questions[length:]

	for _, s := range a {
		err := questionDAO.UpdateCount(&s)
		if err != nil {
			log.Logger.Warnf("Error updating count")
		}
	}

	return questions[length:]
}

func getRandomWeightedIndex(questions []model.Question, maxi int64) int {
	res := make([]model.Question, 0)

	for _, q := range questions {
		var j int64 = 0
		for j < maxi-q.Count {
			res = append(res, q)
			j++
		}
	}
	all_length := len(res)
	sel := res[rand.Intn(all_length)]
	for i, q := range questions {
		if q == sel {
			return i
		}
	}

	return 0
}
