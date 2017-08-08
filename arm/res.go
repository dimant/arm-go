package arm

import (
	"../util"
	"strings"
)

type Resource struct {
	Id string
	Name string
	Type string
	Location string
	Tags map[string]interface{}
}

func (res *Resource) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := util.SetField(res, strings.Title(k), v)
		if err != nil {
			return err
		}
	}

	return nil
}