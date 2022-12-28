package checksequence

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	. "task3/solutions"
)

type CheckSequenceSolution struct {
	TaskName   string
	DataSource string
	CheckHost  string
}

func (s *CheckSequenceSolution) SolveTask() (ResolutionResult, error) {
	failedResolution := ResolutionResult{0, nil}
	response, error := http.Get(s.DataSource)

	if error != nil {
		log.Fatalf("Fail to get data for %s error %s", s.TaskName, error)
		return failedResolution, error
	}

	var rawData WeidArrayData
	err := json.NewDecoder(response.Body).Decode(&rawData)

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
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

	if err != nil {
		log.Fatalf("Fail to parse data: %q", error)
		return failedResolution, error
	}

	var resolutionResults ResolutionResult
	json.NewDecoder(postResponse.Body).Decode(&resolutionResults)

	return resolutionResults, nil
}

func Solution(A []int) int {
	min := A[0]
	max := A[0]
	sumValues := 0
	sumSeq := 0

	for _, v := range A {
		sumValues += v
		if min > v {
			min = v
		}

		if max < v {
			max = v
		}
	}

	for i := min; i <= max; i++ {
		sumSeq += i
	}

	if sumSeq == sumValues {
		return 1
	}

	return 0
}
