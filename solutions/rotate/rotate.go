package rotate

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	. "task3/solutions"
)

type RotateSolution struct {
	TaskName   string
	DataSource string
	CheckHost  string
}

func (s *RotateSolution) SolveTask() (ResolutionResult, error) {
	failedResolution := ResolutionResult{0, nil}
	response, error := http.Get(s.DataSource)

	if error != nil {
		log.Fatalf("Fail to get data for task %s; error %s", s.TaskName, error)
		return failedResolution, nil
	}

	var rawData RevertData
	error = json.NewDecoder(response.Body).Decode(&rawData)

	if error != nil {
		log.Fatalf("Fail to parse data: %q", error)
		return failedResolution, nil
	}

	var results [][]int
	// parse array and run solution
	for _, data := range rawData {
		rawArray := reflect.ValueOf(data[0])
		rawK := reflect.ValueOf(data[1])

		K := int(rawK.Float())
		array := make([]int, rawArray.Len())

		for i := 0; i < rawArray.Len(); i++ {
			array[i] = int(rawArray.Index(i).Interface().(float64))
		}

		results = append(results, Solution(array, K))
	}

	resolution := Resolution{
		UserName: "kmisha",
		Task:     s.TaskName,
		Results: &Results{
			Payload: rawData,
			Results: results,
		},
	}

	// send POST request
	body, _ := json.Marshal(resolution)
	postResponse, error := http.Post(s.CheckHost, "application/json", bytes.NewReader(body))

	if error != nil {
		log.Fatalf("Fail to parse data: %q", error)
		return failedResolution, error
	}

	var resolutionResults ResolutionResult
	json.NewDecoder(postResponse.Body).Decode(&resolutionResults)

	return resolutionResults, nil
}

func Solution(A []int, K int) []int {
	amount := len(A)
	result := make([]int, amount)

	for i, n := range A {
		idx := (i + K) % amount
		result[idx] = n
	}

	return result
}
