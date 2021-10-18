package post

import (
	"api/pkg/entities"
)

type AdminService interface {
	InsertPost(post *entities.Post) (*entities.Post, error)
	FetchPost(ID string) (*entities.Post, error)
	FetchPosts() (*[]entities.Post, error)
	UpdatePost(post *entities.Post) (*entities.Post, error)
	RemovePost(ID string) error
}

type adminService struct {
	repository Repository
}

//NewAdminService is used to create a single instance of the publicService
func NewAdminService(r Repository) AdminService {
	return &adminService{
		repository: r,
	}
}

func (s *adminService) InsertPost(post *entities.Post) (*entities.Post, error) {
	return s.repository.CreatePost(post)
}
func (s *adminService) FetchPost(ID string) (*entities.Post, error) {
	return s.repository.ReadPost(ID)
}
func (s *adminService) FetchPosts() (*[]entities.Post, error) {
	return s.repository.ReadPosts()
}
func (s *adminService) UpdatePost(post *entities.Post) (*entities.Post, error) {
	return s.repository.UpdatePost(post)
}
func (s *adminService) RemovePost(ID string) error {
	return s.repository.DeletePost(ID)
}
