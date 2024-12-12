package todo

import (
	"context"
	"errors"
	"fmt"
	"my-first-api/internal/db"
	"strings"
)

type Item struct {
	Task   string
	Status string
}

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Add(todo string) error {
	items, err := s.GetAll()
	if err != nil {
		return fmt.Errorf("could not get all items: %w", err)
	}

	for _, item := range items {
		if item.Task == todo {
			return errors.New("task already exists")
		}
	}

	if err := s.db.InsertItem(context.Background(), db.Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	}); err != nil {
		return fmt.Errorf("could not insert item: %w", err)
	}
	return nil
}

func (s *Service) Search(query string) ([]string, error) {
	items, err := s.GetAll()
	if err != nil {
		return nil, fmt.Errorf("could not get all items: %w", err)
	}

	var results []string
	for _, todo := range items {
		if strings.Contains(strings.ToLower(todo.Task), strings.ToLower(query)) {
			results = append(results, todo.Task)
		}
	}
	return results, nil
}

func (s *Service) GetAll() ([]Item, error) {
	var results []Item
	items, err := s.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not get all items: %w", err)
	}
	for _, item := range items {
		results = append(results, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}
	return results, nil
}
