package yaml

import (
	"errors"
	"io/ioutil"
)

// LoadYAMLFile is to read yaml file
func LoadYAMLFile(filePath string) ([]byte, error) {
	// Loading jsonfile
	if filePath == "" {
		err := errors.New("nothing JSON file")
		return nil, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
