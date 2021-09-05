package post

import (
	database "api/db"
	"api/pkg/entities"
)

type Repository interface {
	CreatePost(post *entities.Post) (*entities.Post, error)
	ReadPost() (*[]entities.Post, error)
	UpdatePost(post *entities.Post) (*entities.Post, error)
	DeletePost(ID string) error
}

type repository struct {
	Collection database.Dbinstance
}

func NewRepo(collection database.Dbinstance) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreatePost(post *entities.Post) (*entities.Post, error) {
	if err := r.Collection.Db.Create(&post).Error; err != nil {
		// Create failed, do something e.g. return, panic etc.
		return nil, err
	}

	return post, nil
}

func (r *repository) ReadPost() (*[]entities.Post, error) {
	var posts []entities.Post

	if err := r.Collection.Db.Find(&posts).Error; err != nil {
		// Create failed, do something e.g. return, panic etc.
		return nil, err
	}

	return &posts, nil
}

func (r *repository) UpdatePost(post *entities.Post) (*entities.Post, error) {
	if err := r.Collection.Db.Model(&post).Updates(&post).Error; err != nil {
		// Create failed, do something e.g. return, panic etc.
		return nil, err
	}

	return post, nil
}

func (r *repository) DeletePost(ID string) error {
	if err := r.Collection.Db.Delete(&entities.Post{}, ID).Error; err != nil {
		// Create failed, do something e.g. return, panic etc.
		return err
	}
	return nil
}
