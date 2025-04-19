package models

import "time"

// Reception модель приемки товаров
type Reception struct {
	ID       string    `json:"id"`
	DateTime time.Time `json:"dateTime"`
	PVZID    string    `json:"pvzId"`
	Products []Product `json:"products,omitempty"`
	Status   string    `json:"status"` // in_progress, close
}
