package xml

import (
	"errors"
	"io/ioutil"
)

//-----------------------------------------------------------------------------
// Util
//-----------------------------------------------------------------------------
func LoadXmlFile(filePath string) ([]byte, error) {
	// Loading jsonfile
	if filePath == "" {
		err := errors.New("Nothing XML File")
		return nil, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
