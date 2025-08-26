package user

import "errors"

type Repository struct {
	users  []User
	nextID int64
}

func NewRepository() *Repository {
	return &Repository{users: []User{}}
}

func (r *Repository) Create(user User) (User, error) {
	// check for duplicate
	for _, u := range r.users {
		if u.Username == user.Username {
			return User{}, errors.New("username already exists")
		}
	}
	r.nextID++
	user.ID = r.nextID
	r.users = append(r.users, user)
	return user, nil
}

func (r *Repository) GetByUsername(username string) (User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return User{}, errors.New("username not found")
}
