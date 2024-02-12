package model

import "github.com/gocql/gocql"

type Candy struct {
	ID         gocql.UUID `json:"id,omitempty"`
	Candyname  string     `json:"candyname"`
	CandyPrice float32    `json:"candy_price"`
}


type UpdatedCandy struct {
	Candyname  string     `json:"candyname"`
	CandyPrice float32    `json:"candy_price"`
}