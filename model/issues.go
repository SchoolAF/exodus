package model

import "time"

type IssueData struct {
	IssueID      string    `json:"issue_id"`
	UserID       string    `json:"user_id"`
	AuthorName   string    `json:"author_name"`
	Date         time.Time `json:"date"`
	Title        string    `json:"title" valid:"required"`
	Device       string    `json:"device" valid:"required"`
	DeviceParsed string    `json:"device_parsed"`
	Version      string    `json:"Version" valid:"required"`
	Description  string    `json:"description" valid:"required"`
	Edited       bool      `json:"Edited"`
	Status       string    `json:"status"`
	Notify       bool      `json:"allow_notify"`
	Attachment   string    `json:"attachment_url"`
}

type Comments struct {
	IssueID    string    `json:"issue_id" valid:"required"`
	CommentID  string    `json:"comment_id" valid:"required"`
	UserID     string    `json:"user_id" valid:"required"`
	AuthorName string    `json:"author_name"`
	Date       time.Time `json:"date" valid:"required"`
	Content    string    `json:"description" valid:"required"`
	Status     string    `json:"status"`
}
