package db

import (
	"encoding/json"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

func Mapper(i *tracer.Infos, data interface{}, to interface{}) error {
	jsonResult, err := json.Marshal(data)
	if err != nil {
		i.LogError(err)
		return err
	}

	err = json.Unmarshal(jsonResult, to)
	if err != nil {
		i.LogError(err)
		return err
	}
	return nil
}

