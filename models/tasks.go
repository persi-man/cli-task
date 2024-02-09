package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
type Status string

var (
	StatusDone    Status = "done"
	StatusPending Status = "pending"
)

func LoadTasks() error {

	data, err := ioutil.ReadFile("../assets/tasks.json")
	tasks := []Task{}
	if err != nil {
		if os.IsNotExist(err) {
			// No tasks.json file exists yet, so we start with an empty list
			tasks = []Task{} // Initialize to an empty slice if no file exists
			return nil
		}
		return err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}

	return nil
}
