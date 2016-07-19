package structs

import "time"

type Event struct {
	Action    string            `json:"action"`   // app:create, release:create, release:promote, etc.
	Status    string            `json:"status"`   // success or error
	Duration  int64             `json:"duration"` // 1000ms
	Data      map[string]string `json:"data"`     // {"rack": "example-rack", "app": "example-app", "id": "R123456789", "message": "unable to load release"}
	Timestamp time.Time         `json:"timestamp"`
}
