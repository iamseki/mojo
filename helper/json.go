package helper

import (
	"encoding/json"
	"io/ioutil"
)

// JSONFileToStruct puts data from a arbitrary file.json into target struct
func JSONFileToStruct(filepath string, target interface{}) error {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &target)
	if err != nil {
		return err
	}

	return nil
}
