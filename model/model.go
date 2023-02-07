package model

import "time"

type Comment struct {
	Id           uint64 `json:"Id"`
	PostId       uint64 `json:"PostId"`
	Comment      string `json:"Comment"`
	Author       string `json:"Author"`
	CreationDate time.Time `json:"CreationDate"`
}

func (c Comment) Validate() bool {
	if c.Id == 0 {
		return false
	}

	if c.PostId == 0 {
		return false
	}

	if len(c.Comment) == 0 {
		return false
	}

	if len(c.Author) == 0 {
		return false
	}

	return true
}

type Post struct {
	Id           uint64
	Title        string
	Content      string
	CreationDate time.Time
}
