package main

import (
	"encoding/json"
	"log"
	"net/http"
	"task3/solutions"
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

	r.Get("/task/{taskId}", solveTask)
	r.Get("/tasks", solveAllTasks)

	log.Print("Start server at 3000 port")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func solveTask(w http.ResponseWriter, r *http.Request) {
	task := chi.URLParam(r, "taskId")
	log.Printf("got taks %s", task)
	solver := selectSover(task)
	res, _ := solver.SolveTask()

	data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func solveAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []string{ROTATE, WEIRD_ARRAY, CHECK_SEQ, SEQUENCE}
	var results []solutions.ResolutionResult

	for _, t := range tasks {
		res, _ := selectSover(t).SolveTask()
		results = append(results, res)
	}

	// write result
	raw, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.Write(raw)
}

func selectSover(t string) solutions.TaskSolution {
	rotateSolver := rotate.RotateSolution{
		TaskName:   ROTATE,
		DataSource: TASKS_HOST + ROTATE,
		CheckHost:  SOLUTION_HOST,
	}

	waidArraySolver := weirdarray.WeirdArraySolution{
		TaskName:   WEIRD_ARRAY,
		DataSource: TASKS_HOST + WEIRD_ARRAY,
		CheckHost:  SOLUTION_HOST,
	}

	checkSequenceSolver := checksequence.CheckSequenceSolution{
		TaskName:   CHECK_SEQ,
		DataSource: TASKS_HOST + CHECK_SEQ,
		CheckHost:  SOLUTION_HOST,
	}

	sequenseSolver := sequence.SequenceSolution{
		TaskName:   SEQUENCE,
		DataSource: TASKS_HOST + SEQUENCE,
		CheckHost:  SOLUTION_HOST,
	}

	switch t {
	case ROTATE:
		return &rotateSolver
	case WEIRD_ARRAY:
		return &waidArraySolver
	case CHECK_SEQ:
		return &checkSequenceSolver
	case SEQUENCE:
		return &sequenseSolver
	default:
		return &rotateSolver
	}
}
