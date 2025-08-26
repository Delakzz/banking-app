package user

import (
	"banking-app/backend/internal/bank"
	"banking-app/backend/internal/customer"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Repository struct {
	file string
	mu   sync.RWMutex
	data struct {
		Banks     []*bank.Bank         `json:"banks"`
		Customers []*customer.Customer `json:"customers"`
		Users     []User               `json:"users"`
	}
	nextID int64
}

func NewRepository(file string) (*Repository, error) {
	r := &Repository{file: file}
	err := r.load()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Repository) load() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	f, err := os.Open(r.file)
	if err != nil {
		if os.IsNotExist(err) {
			// file not found, init empty
			r.data.Users = []User{}
			r.data.Banks = []*bank.Bank{}
			r.data.Customers = []*customer.Customer{}
			return nil
		}
		return err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&r.data); err != nil {
		return err
	}

	// Ensure Users array exists even if not in file
	if r.data.Users == nil {
		r.data.Users = []User{}
	}

	return nil
}

func (r *Repository) save() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	f, err := os.Create(r.file)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(r.data)
}

func (r *Repository) Create(user User) (User, error) {
	// check for duplicate
	for _, u := range r.data.Users {
		if u.Username == user.Username {
			return User{}, errors.New("username already exists")
		}
	}
	r.nextID++
	user.ID = r.nextID
	r.data.Users = append(r.data.Users, user)

	// Save the updated data to the file
	if err := r.save(); err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *Repository) GetByUsername(username string) (User, error) {
	for _, user := range r.data.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return User{}, errors.New("username not found")
}
