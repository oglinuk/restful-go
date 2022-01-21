package main

type BookResp struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Published string `json:"published"`
	Genre string `json:"genre"`
	ReadStatus string `json:"readstatus"`
}

type BooksResp struct {
	Books []BookResp `json:"books"`
}

type HeartbeatResp struct {
	IP string `json:"ip"`
}
