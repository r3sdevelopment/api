package post

import (
	"api/pkg/entities"
)

type Service interface {
	InsertPost(post *entities.Post) (*entities.Post, error)
	FetchPosts() (*[]entities.Post, error)
	UpdatePost(post *entities.Post) (*entities.Post, error)
	RemovePost(ID string) error
}

type service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertPost(post *entities.Post) (*entities.Post, error) {
	return s.repository.CreatePost(post)
}
func (s *service) FetchPosts() (*[]entities.Post, error) {
	return s.repository.ReadPost()
}
func (s *service) UpdatePost(post *entities.Post) (*entities.Post, error) {
	return s.repository.UpdatePost(post)
}
func (s *service) RemovePost(ID string) error {
	return s.repository.DeletePost(ID)
}
