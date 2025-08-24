package main

import (
	"errors"
	"time"
)

type User struct {
	ID         int64             `json:"id"`
	Attributes map[string]string `json:"attributes,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
}

func (s *Store) CreateUser(id int64, attrs map[string]string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[id]; exists {
		return errors.New("user already exists")
	}
	s.users[id] = &User{ID: id, Attributes: attrs, CreatedAt: time.Now()}
	return nil
}

func (s *Store) DeleteUser(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(s.users, id)
	for k := range s.explicitAdds {
		delete(s.explicitAdds[k], id)
	}
	for k := range s.explicitRems {
		delete(s.explicitRems[k], id)
	}
	return nil
}
