package question_service

import (
	"encoding/json"
	"fmt"
	"singlishwords/model"
	"singlishwords/service"
	"singlishwords/util"
	"strings"
)

// var ctx = context.Background()
var db = service.GetMySQLDB()
var rdb = service.GetRedisDB()

func CacheGetQuestions(limit int) ([]model.Question, error) {
	questions := []model.Question{}
	exist, err := rdb.Exists("question").Result()

	fmt.Println("Exists: ", exist, err)

	if err != nil {
		return questions, err
	}

	if exist == 1 {
		questions, err = getRandomQuestionsFromRedis(limit)
		if err != nil {
			return nil, err
		}
		return questions, nil
	} else {
		questions, err = getAllQuestionsFromDB()
		if err != nil {
			return questions, err
		}
		err = storeAllQuestionsToRedis(questions)
		if err != nil {
			return questions, err
		}
		questions = randomPickNQuestions(questions, limit)
	}
	return questions, err
}

func getRandomQuestionsFromRedis(limit int) ([]model.Question, error) {
	questions := []model.Question{}
	results, err := rdb.SRandMemberN("question", int64(limit)).Result()
	fmt.Println("get from redis results", results)
	if err != nil {
		return questions, err
	}

	jsonString := "[" + strings.Join(results, ", ") + "]"
	fmt.Println("JsonString", jsonString)

	err = json.Unmarshal([]byte(jsonString), &questions)
	if err != nil {
		return questions, err
	}

	return questions, err
}

func storeAllQuestionsToRedis(questions []model.Question) error {
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

func getAllQuestionsFromDB() ([]model.Question, error) {
	questions := []model.Question{}
	sql := "SELECT * FROM `question`"
	err := db.Select(&questions, sql)
	return questions, err
}

func randomPickNQuestions(questions []model.Question, n int) []model.Question {
	shuffled, ok := util.ShufflePickN(questions, n)
	if !ok {
		return nil
	}

	results := make([]model.Question, len(shuffled))
	for i, qi := range shuffled {
		results[i] = qi.(model.Question)
	}

	return results
}
