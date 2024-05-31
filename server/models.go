package main

type User struct {
	ID      int32   `json:"id"`
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int64   `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}
