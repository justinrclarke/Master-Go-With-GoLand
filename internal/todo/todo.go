package todo

import (
	"errors"
)

type Item struct {
	Task   string
	Status string
}

type Service struct {
	todos []Item
}

func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

func (s *Service) Add(todo string) error {
	for _, t := range s.todos {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}
	s.todos = append(s.todos, Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	})
	return nil
}

func (s *Service) GetAll() []Item {
	return s.todos
}
