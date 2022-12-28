package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	checksequence "task3/solutions/check-sequence"
	"task3/solutions/rotate"
	"task3/solutions/sequence"
	weirdarray "task3/solutions/weird-array"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	ROTATE        = "Циклическая ротация"
	WEIRD_ARRAY   = "Чудные вхождения в массив"
	CHECK_SEQ     = "Проверка последовательности"
	SEQUENCE      = "Поиск отсутствующего элемента"
	TASKS_HOST    = "https://kuvaev-ituniversity.vps.elewise.com/tasks/"
	SOLUTION_HOST = "https://kuvaev-ituniversity.vps.elewise.com/tasks/solution"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/task/{taskId}", getTask)
	log.Print("Start server at 3000 port")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func getTask(w http.ResponseWriter, r *http.Request) {
	task := chi.URLParam(r, "taskId")
	log.Printf("got taks %s", task)
	switch task {
	case ROTATE:
		checkRevert(task, w)
	case WEIRD_ARRAY:
		checkWeirdArray(task, w)
	case CHECK_SEQ:
		checkCheckSequence(task, w)
	case SEQUENCE:
		checkSequence(task, w)
	}

}

func checkRevert(task string, w http.ResponseWriter) {
	response, error := http.Get(fmt.Sprintf("%s%s", TASKS_HOST, task))

	if error != nil {
		log.Fatalf("Fail to get task %s; error %s", task, error)
		return
	}

	var rawData RevertData
	err := json.NewDecoder(response.Body).Decode(&rawData)

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
		return
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

		results = append(results, rotate.Solution(array, K))
	}

	resolution := Resolution{
		UserName: "kmisha",
		Task:     task,
		Results: &Results{
			Payload: rawData,
			Results: results,
		},
	}

	// send POST request
	body, _ := json.Marshal(resolution)
	postResponse, err := http.Post(SOLUTION_HOST, "application/json", bytes.NewReader(body))

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
		return
	}

	var resolutionResults ResolutionResult
	json.NewDecoder(postResponse.Body).Decode(&resolutionResults)

	jData, _ := json.Marshal(resolutionResults)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func checkWeirdArray(task string, w http.ResponseWriter) {
	response, error := http.Get(fmt.Sprintf("%s%s", TASKS_HOST, task))

	if error != nil {
		log.Fatalf("Fail to get task %s; error %s", task, error)
		return
	}

	var rawData WeidArrayData
	err := json.NewDecoder(response.Body).Decode(&rawData)

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
		return
	}

	// parse array
	var results []int
	for _, array := range rawData {
		results = append(results, weirdarray.Solution(array[0]))
	}

	resolution := Resolution{
		UserName: "kmisha",
		Task:     task,
		Results: &Results{
			Payload: rawData,
			Results: results,
		},
	}

	// send POST request
	body, _ := json.Marshal(resolution)
	postResponse, err := http.Post(SOLUTION_HOST, "application/json", bytes.NewReader(body))

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
		return
	}

	var resolutionResults ResolutionResult
	json.NewDecoder(postResponse.Body).Decode(&resolutionResults)

	jData, _ := json.Marshal(resolutionResults)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func checkCheckSequence(task string, w http.ResponseWriter) {
	response, error := http.Get(fmt.Sprintf("%s%s", TASKS_HOST, task))

	if error != nil {
		log.Fatalf("Fail to get task %s; error %s", task, error)
		return
	}

	var rawData WeidArrayData
	err := json.NewDecoder(response.Body).Decode(&rawData)

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
		return
	}

	// parse array
	var results []int
	for _, array := range rawData {
		results = append(results, checksequence.Solution(array[0]))
	}

	resolution := Resolution{
		UserName: "kmisha",
		Task:     task,
		Results: &Results{
			Payload: rawData,
			Results: results,
		},
	}

	// send POST request
	body, _ := json.Marshal(resolution)
	postResponse, err := http.Post(SOLUTION_HOST, "application/json", bytes.NewReader(body))

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
		return
	}

	var resolutionResults ResolutionResult
	json.NewDecoder(postResponse.Body).Decode(&resolutionResults)

	jData, _ := json.Marshal(resolutionResults)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func checkSequence(task string, w http.ResponseWriter) {
	response, error := http.Get(fmt.Sprintf("%s%s", TASKS_HOST, task))

	if error != nil {
		log.Fatalf("Fail to get task %s; error %s", task, error)
		return
	}

	var rawData WeidArrayData
	err := json.NewDecoder(response.Body).Decode(&rawData)

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
		return
	}

	// parse array
	var results []int
	for _, array := range rawData {
		results = append(results, sequence.Solution(array[0]))
	}

	log.Printf("results : %v", results)

	resolution := Resolution{
		UserName: "kmisha",
		Task:     task,
		Results: &Results{
			Payload: rawData,
			Results: results,
		},
	}

	// send POST request
	body, _ := json.Marshal(resolution)
	postResponse, err := http.Post(SOLUTION_HOST, "application/json", bytes.NewReader(body))

	if err != nil {
		log.Fatalf("Fail to parse data: %q", err)
		return
	}

	var resolutionResults ResolutionResult
	json.NewDecoder(postResponse.Body).Decode(&resolutionResults)

	jData, _ := json.Marshal(resolutionResults)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

type RevertData [][]interface{}
type WeidArrayData [][][]int

type Results struct {
	Payload any `json:"payload"` // данные полученные для решения задачи
	Results any `json:"results"` // результаты полученные при решении задачи
}

type Resolution struct {
	UserName string   `json:"user_name"` // "имя юзера указанное в тг",
	Task     string   `json:"task"`      // "имя задачи",
	Results  *Results `json:"results"`
}

type Fail struct {
	OriginalResult int
	ExternalResult int
}

type ResolutionResult struct {
	Percent int    `json:"percent"` //: 90,
	Fails   []Fail `json:"fails"`
}
