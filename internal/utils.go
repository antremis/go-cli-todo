package internal

import "github.com/google/uuid"

func (s *FileStorage) doesIdExist(id uuid.UUID) bool {
	if _, ok := s.Filedata[id]; ok {
		return true
	}
	return false
}
