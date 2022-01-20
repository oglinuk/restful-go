package main

type BooksResp struct {
	Books []struct {
		ID string
		Title string
		Author string
		Published string
		Genre string
		ReadStatus string
	}	`json:"books"`
}

type HeartbeatResp struct {
	IP string `json:"ip"`
}
