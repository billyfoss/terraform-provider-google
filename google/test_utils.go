package google

import (
	"reflect"
	"strconv"
)

type ResourceDataMock struct {
	FieldsInSchema      map[string]interface{}
	FieldsWithHasChange []string
	id                  string
}

func (d *ResourceDataMock) HasChange(key string) bool {
	exists := false
	for _, val := range d.FieldsWithHasChange {
		if key == val {
			exists = true
		}
	}

	return exists
}

func (d *ResourceDataMock) Get(key string) interface{} {
	v, _ := d.GetOk(key)
	return v
}

func (d *ResourceDataMock) GetOk(key string) (interface{}, bool) {
	v, ok := d.GetOkExists(key)
	if ok && !isEmptyValue(reflect.ValueOf(v)) {
		return v, true
	} else {
		return v, false
	}
}

func (d *ResourceDataMock) GetOkExists(key string) (interface{}, bool) {
	for k, v := range d.FieldsInSchema {
		if key == k {
			return v, true
		}
	}

	return nil, false
}

func (d *ResourceDataMock) Set(key string, value interface{}) error {
	d.FieldsInSchema[key] = value
	return nil
}

func (d *ResourceDataMock) SetId(v string) {
	d.id = v
}

func (d *ResourceDataMock) Id() string {
	return d.id
}

func toBool(attribute string) (bool, error) {
	// Handle the case where an unset value defaults to false
	if attribute == "" {
		return false, nil
	}
	return strconv.ParseBool(attribute)
}
