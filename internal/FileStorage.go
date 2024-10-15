package internal

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

type FileStorage struct {
	Filename string
	Filedata map[uuid.UUID]Todo
}

func NewFileStorage(filename string) *FileStorage {
	s := &FileStorage{
		Filename: filename,
	}
	s.Read()
	return s
}

func (s *FileStorage) Read() {
	filedata, err := os.ReadFile(s.Filename)
	if err != nil {
		s.Filedata = make(map[uuid.UUID]Todo)
		return
	}
	err = json.Unmarshal(filedata, &s.Filedata)
	if err != nil {
		s.Filedata = make(map[uuid.UUID]Todo)
		return
	}
}

func (s *FileStorage) Write() error {
	filedata, err := json.MarshalIndent(s.Filedata, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Filename, filedata, 0777)
}
