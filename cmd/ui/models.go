package main

type bookResp struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Published string `json:"published"`
	Genre string `json:"genre"`
	ReadStatus string `json:"readstatus"`
}

type booksResp struct {
	Books []bookResp `json:"books"`
}

type heartbeatResp struct {
	IP string `json:"ip"`
}
