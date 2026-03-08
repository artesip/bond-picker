package domain

import "time"

type Event struct {
	ID     UUID       `json:"id"`
	Type   string     `json:"type"`
	Status string     `json:"status"`
	Msg    string     `json:"msg"`
	Start  *time.Time `json:"start"`
	End    *time.Time `json:"end"`
}
