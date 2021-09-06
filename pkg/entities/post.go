package entities

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	UserId   string     `json:"userId" gorm:"primarykey;type:uuid;"`
	Status   PostStatus `json:"status" sql:"DEFAULT:0"`
}

func (p Post) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.UserId, validation.Required, is.UUIDv4),
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.Body, validation.Required),
	)
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("ID", uuid.NewString())

	if p.Validate() != nil {
		err = errors.New("can't save invalid data")
	}
	return nil
}
