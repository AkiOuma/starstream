package generator

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AkiOuma/starstream/internal/starstream/frame"
)

type Generator struct {
	frame *frame.Frame
}

func NewGenerator(frame *frame.Frame) *Generator {
	return &Generator{
		frame: frame,
	}
}

func (g *Generator) Launch() {
	// mkdir destination
	if err := os.MkdirAll(g.frame.Destination, os.ModePerm); err != nil {
		log.Fatalf("ERROR: Generate %s failed because: %v", g.frame.Destination, err)
	}
	// init go mod
	if err := g.initGoModule(); err != nil {
		log.Printf("ERROR: init go module %s failed because: %v", g.frame.ServiceName, err)
	}
	// build internal
	if err := g.buildInternal(); err != nil {
		log.Fatalf("ERROR: build internal failed because: %v", err)
	}
}

func (g *Generator) initGoModule() (err error) {
	cmd := exec.Command("go", "mod", "init", g.frame.ServiceName)
	cmd.Dir = g.frame.Destination
	log.Printf("go: creating new go.mod: module %s", g.frame.ServiceName)
	output, err := cmd.Output()
	if err != nil {
		return
	}
	log.Println(string(output))
	return
}

func (g *Generator) buildInternal() error {
	if err := g.buildDomain(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) buildDomain() error {
	for _, v := range g.frame.Entity {
		// valueobject
		valueObjectFolder := filepath.Dir(v.ValueObject.FilePath)
		if err := os.MkdirAll(valueObjectFolder, os.ModePerm); err != nil {
			log.Fatalf("ERROR: Generate %s failed because: %v", v.FilePath, err)
		}
		if err := g.buildValueObject(v.ValueObject); err != nil {
			log.Fatalf("ERROR: Generate ValueObject %s failed because: %v", v.ValueObject.FilePath, err)
		}
		// domain
		domainFolder := filepath.Dir(v.FilePath)
		if err := os.MkdirAll(domainFolder, os.ModePerm); err != nil {
			log.Fatalf("ERROR: Generate %s failed because: %v", v.FilePath, err)
		}
		if err := g.buildEntity(v); err != nil {
			log.Fatalf("ERROR: Generate entity %s failed because: %v", v.FilePath, err)
		}
		// repository
		repositoryFolder := filepath.Dir(v.Repository.FilePath)
		if err := os.MkdirAll(repositoryFolder, os.ModePerm); err != nil {
			log.Fatalf("ERROR: Generate %s failed because: %v", v.FilePath, err)
		}
		if err := g.buildRepository(v.Repository); err != nil {
			log.Fatalf("ERROR: Generate repository %s failed because: %v", v.Repository.FilePath, err)
		}
		// service
		serviceFolder := filepath.Dir(v.Service.FilePath)
		if err := os.MkdirAll(serviceFolder, os.ModePerm); err != nil {
			log.Fatalf("ERROR: Generate %s failed because: %v", serviceFolder, err)
		}
		if err := g.buildService(v.Service); err != nil {
			log.Fatalf("ERROR: Generate service %s failed because: %v", v.Service.FilePath, err)
		}
		// usecase
		usecaseFolder := filepath.Dir(v.Usecase.FilePath)
		if err := os.MkdirAll(usecaseFolder, os.ModePerm); err != nil {
			log.Fatalf("ERROR: Generate %s failed because: %v", usecaseFolder, err)
		}
		if err := g.buildUsecase(v.Usecase); err != nil {
			log.Fatalf("ERROR: Generate service %s failed because: %v", v.Usecase.FilePath, err)
		}
	}
	return nil
}
