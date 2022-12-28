package weirdarray

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	. "task3/solutions"
)

type WeirdArraySolution struct {
	TaskName   string
	DataSource string
	CheckHost  string
}

func (s *WeirdArraySolution) SolveTask() (ResolutionResult, error) {
	failedResolution := ResolutionResult{0, nil}
	response, error := http.Get(s.DataSource)

	if error != nil {
		log.Fatalf("Fail to get data for task %s; error %s", s.TaskName, error)
		return failedResolution, nil
	}

	var rawData WeidArrayData
	error = json.NewDecoder(response.Body).Decode(&rawData)

	if error != nil {
		log.Fatalf("Fail to parse data: %q", error)
		return failedResolution, error
	}

	// parse array
	var results []int
	for _, array := range rawData {
		results = append(results, Solution(array[0]))
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
		return failedResolution, nil
	}

	var resolutionResults ResolutionResult
	json.NewDecoder(postResponse.Body).Decode(&resolutionResults)

	return resolutionResults, nil
}

func Solution(A []int) int {
	pairs := make(map[int]int)

	for _, n := range A {
		v, ok := pairs[n]

		if ok {
			pairs[n] = v + 1
		} else {
			pairs[n] = 1
		}

	}

	for n, amount := range pairs {
		if amount%2 != 0 {
			return n
		}
	}

	return 0
}
