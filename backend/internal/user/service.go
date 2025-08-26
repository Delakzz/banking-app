package user

import "errors"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Register(username, password string, role Role) (User, error) {
	user := User{
		Username: username,
		Password: password,
		Role:     role,
	}
	return s.repo.Create(user)
}

func (s *Service) Login(username, password string) (User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return User{}, err
	}

	if user.Password != password {
		return User{}, errors.New("invalid password")
	}

	return user, nil
}
