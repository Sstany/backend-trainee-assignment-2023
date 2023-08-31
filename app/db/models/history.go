package models

import "time"

const (
	TypeAdd    = "add"
	TypeDelete = "delete"
)

type History struct {
	ID          int       `json:"historyId"`
	UserId      int       `json:"userId"`
	Time        time.Time `json:"modificationTime"`
	Type        string    `json:"type"`
	SegmentName string    `json:"segment"`
}
