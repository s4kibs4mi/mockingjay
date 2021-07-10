package services

//go:generate mockgen -source=user_service.go -destination mocks/mock_user_service.go -package mocks UserService

type UserService interface {
	CreateUser(name, email string) (*User, error)
	UpdateUser(ID, name, email string) (*User, error)
	GetUser(ID string) (*User, error)
	DeleteUser(ID string) error
}

type User struct {
	ID    string
	Name  string
	Email string
}
