package session

import "encoding/json"

type Result struct {
	Pages int
	Time  float64
	Map   map[any]any
}

// Unpack -> right now I have not the slightest clue if this does
// exactly what I would like it to or not
func (r *Result) Unpack() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
