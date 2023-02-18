package JSON

import (
	"encoding/json"
	"fmt"
	"learn-mongo/pkg/models/Todos"
	"reflect"
)

func STRINGIFY(data interface{}) ([]byte, error) {
	reflectVal := reflect.ValueOf(data)

	if reflectVal.Type().Kind() == reflect.Slice {
		result, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		return result, nil
	}

	isTodo, ok := data.(Todos.Todo)

	if !ok {
		return nil, fmt.Errorf("This's should be a struct of a slice...")
	}

	result, _ := json.Marshal(isTodo)

	return result, nil
}
