package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v3"
)

func ReadJSONFile(fileName string, dest interface{}) error {
	// check dest must be pointer
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value")
	}

	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, dest)
}

func ReadYAMLFile(fileName string, dest interface{}) error {
	// check dest must be pointer
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value")
	}

	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(d, dest)
}
