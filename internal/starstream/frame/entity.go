package frame

import (
	"path/filepath"

	"github.com/AkiOuma/starstream/internal/starstream/definition"
)

type Entity struct {
	ImportPath, FilePath string
	ImportPackages       []string
	Name                 string
	Field                []*EntityField
	RootField            *EntityField
	ValueObject          *ValueObject
	Repository           *Repository
	Service              *Service
	Usecase              *Usecase
	Transport            *Transport
}

type EntityField struct {
	Id                      int
	Root                    bool
	PrivateName, PublicName string
	Type                    string
	Getter                  string
}

func NewEntity(def *definition.Entity, info *ServiceInfo) *Entity {
	entity := &Entity{}
	importSet := make(map[string]bool)
	entity.ImportPath = filepath.Join(info.Name, "internal/domain/entity")
	entity.FilePath = filepath.Join(info.Destination, "internal/domain/entity", def.Name+".go")
	entity.Name = GetPublicName(def.Name)
	entity.ImportPackages = make([]string, 0)
	entity.Field = make([]*EntityField, 0, len(def.Field))
	for i, v := range def.Field {
		fieldType := ConvertGoType(v.Type)
		if fieldType == nil {
			continue
		}
		field := &EntityField{
			Id:          i + 1,
			PrivateName: v.Name,
			PublicName:  GetPublicName(v.Name),
			Root:        v.Root,
			Type:        fieldType.TypeName,
			Getter:      "Get" + GetPublicName(v.Name) + "()",
		}
		entity.Field = append(entity.Field, field)
		if v.Root {
			entity.RootField = field
		}
		if len(fieldType.ImportPackage) > 0 && !importSet[fieldType.ImportPackage] {
			entity.ImportPackages = append(entity.ImportPackages, fieldType.ImportPackage)
			importSet[fieldType.ImportPackage] = true
		}
	}
	return entity
}

func (e *Entity) SetValueObject(vo *ValueObject) {
	e.ImportPackages = append(e.ImportPackages, vo.ImportPath)
	e.ValueObject = vo
}

func (e *Entity) SetRepository(repo *Repository) {
	e.Repository = repo
}
