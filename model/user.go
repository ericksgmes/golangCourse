package model

import (
	"errors"
	"github.com/google/uuid"
)

type IUser interface {
	Create(user User) (uuid.UUID, error)
	FindByID(id uuid.UUID) (User, bool)
	DeleteByID(id uuid.UUID) bool
	FindAll() []User
}

type User struct {
	ID   uuid.UUID
	Nome string `json:"nome"`
}

var TotalUsers = make(map[uuid.UUID]User)

var (
	ErrUserExists     = errors.New("usuário já existe")
	ErrUserNotFound   = errors.New("usuário não encontrado")
	ErrInvalidUser    = errors.New("o nome deve ser válido")
)

func Create(user User) (uuid.UUID, error) {
	if user.Nome == "" {
		return uuid.Nil, ErrInvalidUser
	}
	if existsByNome(user.Nome) {
		return uuid.Nil, ErrUserExists
	}
	user.ID = uuid.New()
	TotalUsers[user.ID] = user
	return user.ID, nil
}

func FindByID(id uuid.UUID) (User, error) {
	user, found := TotalUsers[id]
	if !found {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func FindAll() []User {
	users := []User{}
	for _, u := range TotalUsers {
		users = append(users, u)
	}
	return users
}


func DeleteByID(id uuid.UUID) bool {
	if _, found := TotalUsers[id]; found {
		delete(TotalUsers, id)
		return true
	}
	return false
}

func existsByNome(nome string) bool {
	for _, u := range TotalUsers {
		if u.Nome == nome {
			return true
		}
	}
	return false
}
