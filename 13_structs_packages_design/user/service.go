package user

import "errors"

type EmptyNameError struct{}

func (e *EmptyNameError) Error() string {
	return "name is required"
}

type UserRepo interface {
	Save(p Person) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) UserService {
	return UserService{repo: repo}
}

func (s UserService) CreateUser(name string, age int) (Person, error) {
	if name == "" {
		return Person{}, &EmptyNameError{}
	}
	if age <= 0 {
		return Person{}, errors.New("age must be greater than 0")
	}
	u := NewUser(name, age)
	if err := s.repo.Save(u); err != nil {
		return Person{}, err
	}
	return u, nil
}
