package internal

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aquasecurity/table"
	"github.com/google/uuid"
)

func (s *FileStorage) Create(todo TodoCreate) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	if s.doesIdExist(id) {
		return errors.New("Todo already exists")
	}

	s.Filedata[id] = Todo{Title: todo.Title, Completed: false, CreatedAt: time.Now(), CompletedAt: time.Time{}}
	return s.Write()
}

func (s *FileStorage) Update(_id string, todo TodoUpdate) error {
	var id uuid.UUID = uuid.MustParse(_id)
	if !s.doesIdExist(id) {
		return errors.New("Todo does not exist")
	}
	s.Filedata[id] = Todo{Title: todo.Title, Completed: false, CreatedAt: time.Now(), CompletedAt: time.Time{}}
	return s.Write()
}

func (s *FileStorage) Delete(_id string) error {
	var id uuid.UUID = uuid.MustParse(_id)
	if !s.doesIdExist(id) {
		return errors.New("Todo does not exist")
	}
	delete(s.Filedata, id)
	return s.Write()
}

func (s *FileStorage) Get(_id string) {
	var id uuid.UUID = uuid.MustParse(_id)
	if !s.doesIdExist(id) {
		fmt.Printf("Todo does not exist")
		return
	}
	table := table.New(os.Stdout)
	table.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")
	var icon string
	if s.Filedata[id].Completed {icon = "✅"} else {icon = "❌"}

	var CompletedAt string
	if s.Filedata[id].CompletedAt.IsZero() {CompletedAt = ""} else {CompletedAt = s.Filedata[id].CompletedAt.Format(time.RFC1123)}
	
	table.AddRow(id.String(), s.Filedata[id].Title, icon, s.Filedata[id].CreatedAt.Format(time.RFC1123), CompletedAt)
	table.Render()
}

func (s *FileStorage) List() {
	table := table.New(os.Stdout)
	table.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")
	for id, todo := range s.Filedata {
		var icon string
		if todo.Completed {icon = "✅"} else {icon = "❌"}

		var CompletedAt string
		if todo.CompletedAt.IsZero() {
			CompletedAt = ""
		} else {
			CompletedAt = todo.CompletedAt.Format(time.RFC1123)
		}
		
		table.AddRow(id.String(), todo.Title, icon, todo.CreatedAt.Format(time.RFC1123), CompletedAt)
	}
	table.Render()
}

func (s *FileStorage) Complete(_id string) error {
	var id uuid.UUID = uuid.MustParse(_id)
	data, ok := s.Filedata[id]
	if !ok {
		return errors.New("Todo does not exist")
	}
	data.Completed = true
	s.Filedata[id] = data
	return s.Write()
}

func (s *FileStorage) Uncomplete(_id string) error {
	var id uuid.UUID = uuid.MustParse(_id)
	data, ok := s.Filedata[id]
	if !ok {
		return errors.New("Todo does not exist")
	}
	data.Completed = false
	s.Filedata[id] = data
	return s.Write()
}
