package main

type ProverbRequest struct {
	Offset int `json:"offset"`
}

type ProverbResponse struct {
	Proverb string `json:"proverb"`
}
