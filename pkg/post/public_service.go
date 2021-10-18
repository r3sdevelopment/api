package post

import (
	"api/pkg/entities"
)

type PublicService interface {
	FetchPost(ID string) (*entities.Post, error)
	FetchPosts() (*[]entities.Post, error)
}

type publicService struct {
	repository Repository
}

//NewService is used to create a single instance of the publicService
func NewService(r Repository) PublicService {
	return &publicService{
		repository: r,
	}
}

func (s *publicService) FetchPost(ID string) (*entities.Post, error) {
	return s.repository.ReadPost(ID)
}
func (s *publicService) FetchPosts() (*[]entities.Post, error) {
	return s.repository.ReadPosts()
}

