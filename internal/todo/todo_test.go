package todo_test

import (
	"context"
	"github.com/Francismensah/goland-api/internal/db"
	"github.com/Francismensah/goland-api/internal/todo"
	"reflect"
	"testing"
)

type MockDB struct {
	items []db.Item // Fixed: renamed from item to items for clarity
}

func (m *MockDB) InsertItem(_ context.Context, item db.Item) error {
	m.items = append(m.items, item) // Fixed: renamed from item to items
	return nil
}

func (m *MockDB) GetAllItems(_ context.Context) ([]db.Item, error) {
	return m.items, nil // Fixed: renamed from item to items
}

func TestService_Search(t *testing.T) {
	tests := []struct {
		name           string
		toDosToAdd     []string
		query          string
		expectedResult []string
	}{
		{
			name:           "given a todo of shop and a search of sh, I should get shop back",
			toDosToAdd:     []string{"shop"},
			query:          "sh",
			expectedResult: []string{"shop"},
		},
		{
			name:           "still return shop, even if the case doesn't match",
			toDosToAdd:     []string{"Shopping"},
			query:          "sh",
			expectedResult: []string{"Shopping"},
		},
		{
			name:           "spaces",
			toDosToAdd:     []string{"go Shopping"},
			query:          "go",
			expectedResult: []string{"go Shopping"},
		},
		{
			name:           "space at start of word",
			toDosToAdd:     []string{" Space at beginning"},
			query:          "space",
			expectedResult: []string{" Space at beginning"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &MockDB{} // Fixed: properly declared the mock DB
			svc := todo.NewService(mockDB)

			for _, toAdd := range tt.toDosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Error(err)
				}
			}

			got, err := svc.Search(tt.query)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
