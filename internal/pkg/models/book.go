package models

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
)

type Book struct {
	ID        string
	Title     string
	Author    string
	Published string
	Genre string
	ReadStatus string
}

// NewBook constructor
func NewBook(title, author, published, genre, readStatus string) *Book {
	var buff bytes.Buffer
	b := &Book{"", title, author, published, genre, readStatus}
	gob.NewEncoder(&buff).Encode(b)
	hash := md5.New()
	hash.Write(buff.Bytes())
	b.ID = hex.EncodeToString(hash.Sum(nil))
	return b
}
