package session

import "encoding/json"

type Result struct {
	Time float64
	Map  map[any]any
}

func NewResult(t float64, data map[any]any) *Result {
	return &Result{
		Time: t,
		Map:  data,
	}
}

// Unpack -> right now I have not the slightest clue if this does
// exactly what I would like it to or not ...
func (r *Result) Unpack() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(b)
}
