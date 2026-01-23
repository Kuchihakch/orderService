package models

import "time"

type Order struct {
	ID, UserID, Status string
	Amount             int
	CreatedAt          time.Time
}
