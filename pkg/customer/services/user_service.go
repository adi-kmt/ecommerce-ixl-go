package user_services

import user_repositories "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/repositories"

type UserService struct {
	repo user_repositories.UserRepository
}

func NewUserService(repo user_repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
