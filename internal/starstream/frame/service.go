package frame

import (
	"path/filepath"
	"strings"
)

type Service struct {
	Name, ImplName       string
	FilePath, ImportPath string
	ImportPackages       []string
	Repository           *Repository
}

func NewService(entity *Entity, info *ServiceInfo) *Service {
	s := &Service{}
	s.Name = entity.Name
	s.ImplName = strings.ToLower(entity.Name)
	s.ImportPath = filepath.Join(info.Name, "internal/domain/service")
	s.FilePath = filepath.Join(info.Destination, "internal/domain/service", strings.ToLower(entity.Name)+".go")
	s.ImportPackages = make([]string, 0, 1)
	s.ImportPackages = append(s.ImportPackages, entity.Repository.ImportPath)
	s.Repository = entity.Repository
	return s
}
