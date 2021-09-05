package entities

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type PostStatus int

const (
	DRAFT PostStatus = iota
	PUBLISHED
)

type Post struct {
	Base
	Body     string     `json:"body"`
	ImageUrl string     `json:"imageUrl"`
	Title    string     `json:"title"`
	UserId   string     `json:"userId"`
	Status   PostStatus `json:"status" sql:"DEFAULT:0"`
}

func (p Post) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.Body, validation.Required),
	)
}
