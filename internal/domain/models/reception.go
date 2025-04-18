package models

import "time"

// Reception модель приемки товаров
type Reception struct {
	ID       string    `json:"id"`
	DateTime time.Time `json:"dateTime"`
	PVZID    string    `json:"pvzId"`
	Status   string    `json:"status"` // in_progress, close
}
