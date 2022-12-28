package solutions

type TaskSolution interface {
	SolveTask() (ResolutionResult, error)
}

type Fail struct {
	OriginalResult int
	ExternalResult int
}

type ResolutionResult struct {
	Percent int    `json:"percent"` //: 90,
	Fails   []Fail `json:"fails"`
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
