package models

import "time"

type Order struct {
	ID, UserID, ProductID, Status string
	Amount                        int
	CreatedAt                     time.Time
}
