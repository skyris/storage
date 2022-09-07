package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/skyris/storage/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetById(fileID uuid.UUID) (*file.File, error) {
	file, ok := s.files[fileID]
	if !ok {
		// return nil, errors.New(fmt.Sprintf("file %v not fund", fileID))
		return nil, fmt.Errorf("file %v not fund", fileID)
	}
	return file, nil
}
