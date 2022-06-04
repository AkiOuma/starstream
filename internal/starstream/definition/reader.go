package definition

import (
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func ReadDefinition(defDir string) (*Definition, error) {
	var d Definition
	f, err := os.Open(defDir)
	if err != nil {
		return nil, nil
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, nil
	}
	if err := yaml.Unmarshal(data, &d); err != nil {
		return nil, err
	}
	if !filepath.IsAbs(d.Destination) {
		d.Destination, err = filepath.Abs(d.Destination)
		if err != nil {
			return nil, nil
		}
	}
	return &d, nil
}
