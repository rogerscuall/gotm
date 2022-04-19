package tasks

import (
	"encoding/json"
	"time"
)

const (
	timeFormat = "Mon Jan 2 15:04:05 2006"
)

type Task struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	CreationTime   string `json:"creationTime,omitempty"`
	CompletionTime string `json:"completionTime,omitempty"`
}

func New(id, name string) Task {
	t := time.Now().Format(timeFormat)
	return Task{
		ID:           id,
		Name:         name,
		CreationTime: t,
	}
}

func (t *Task) Complete() {
	t.CompletionTime = time.Now().Format(timeFormat)
}

func ToBytes(t Task) []byte {
	b, _ := json.Marshal(t)
	return b
}

func (t *Task) FromBytes(b []byte) error {
	if err:= json.Unmarshal(b, t); err != nil {
		return err
	}
	return nil
}