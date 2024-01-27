// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package gen

import (
	"time"
)

type Admin struct {
	Seq   uint64 `json:"seq"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Pw    []byte `json:"pw"`
	Phone string `json:"phone"`
}

type Menu struct {
	Seq         uint64    `json:"seq"`
	Category    string    `json:"category"`
	Price       int32     `json:"price"`
	Cost        int32     `json:"cost"`
	Name        string    `json:"name"`
	NameInitial string    `json:"name_initial"`
	Description string    `json:"description"`
	Barcode     []byte    `json:"barcode"`
	Expire      time.Time `json:"expire"`
	Size        string    `json:"size"`
}
