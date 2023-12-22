package models

type Thread struct {
	ID      int    `json:"id"`
	Author  User   `json:"author"`
	Tag     Tag    `json:"tag"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
