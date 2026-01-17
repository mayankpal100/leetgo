package store

import (
    "encoding/json"
    "os"
)

type State struct {
    Solved map[string]bool `json:"solved"`
}

func LoadState(path string) (State, error) {
    var s State
    file, _ := os.ReadFile(path)
    if file == nil {
        s.Solved = make(map[string]bool)
        return s, nil
    }
    err := json.Unmarshal(file, &s)
    return s, err
}

func SaveState(path string, s State) error {
    data, _ := json.MarshalIndent(s, "", "  ")
    return os.WriteFile(path, data, 0644)
}
