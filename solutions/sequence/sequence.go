package sequence

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	. "task3/solutions"
)

type SequenceSolution struct {
	TaskName   string
	DataSource string
	CheckHost  string
}

func (s *SequenceSolution) SolveTask() (ResolutionResult, error) {
	failedResolution := ResolutionResult{0, nil}
	response, error := http.Get(s.DataSource)

	if error != nil {
		log.Fatalf("Fail to get data for %s; error %s", s.TaskName, error)
		return failedResolution, error
	}

	var rawData WeidArrayData
	error = json.NewDecoder(response.Body).Decode(&rawData)

	if error != nil {
		log.Fatalf("Fail to parse data: %q", error)
		return failedResolution, nil
	}

	// parse array
	var results []int
	for _, array := range rawData {
		results = append(results, Solution(array[0]))
	}

	log.Printf("results : %v", results)

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

func Solution(A []int) int {
	var sumValues, sumSeq int

	for i, v := range A {
		sumValues += v
		sumSeq += i + 1
	}

	sumSeq += len(A) + 1 // last element of sequence

	return sumSeq - sumValues
}
