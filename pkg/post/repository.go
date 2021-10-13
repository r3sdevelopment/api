package post

import (
	database "api/db"
	"api/pkg/entities"
)

type Repository interface {
	CreatePost(post *entities.Post) (*entities.Post, error)
	ReadPost(id string) (*entities.Post, error)
	ReadPosts() (*[]entities.Post, error)
	UpdatePost(post *entities.Post) (*entities.Post, error)
	DeletePost(ID string) error
}

type repository struct {
	Collection database.DbInstance
}

func NewRepo(collection database.DbInstance) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreatePost(post *entities.Post) (*entities.Post, error) {
	if err := r.Collection.Db.Create(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *repository) ReadPost(ID string) (*entities.Post, error) {
	var post entities.Post

	if err := r.Collection.Db.Where("ID = ?", ID).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *repository) ReadPosts() (*[]entities.Post, error) {
	var posts []entities.Post

	if err := r.Collection.Db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r *repository) UpdatePost(post *entities.Post) (*entities.Post, error) {
	if err := r.Collection.Db.Model(&post).Updates(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *repository) DeletePost(ID string) error {
	post, err := r.ReadPost(ID)

	if err != nil {
		return err
	}

	if err := r.Collection.Db.Where("ID = ?", post.ID).Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
