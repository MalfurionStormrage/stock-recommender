package models

import (
	"time"
)

type APIResponse struct {
	Items    []StockItem `json:"items"`
	NextPage string      `json:"next_page"`
}

type StockItem struct {
	Id         string    `json:"id"`
	Ticker     string    `json:"ticker"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
	Origin     string    `json:"origin"`
	Next_page  string    `json:"next_page"`
}

type StockItemCreateRequest struct {
	Ticker     string    `json:"ticker"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
	Origin     string    `json:"origin"`
	NextPage   string    `json:"next_page"`
}
