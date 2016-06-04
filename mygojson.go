package mygojson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Js struct {
	data interface{}
}

/*
 if use interface{} will be return map or array......
*/
func Json(data string) *Js {
	j := new(Js)
	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	if nil != err {
		return j
	}
	j.data = f
	return j
}

//Acoording to the key of the returned data infarmation , return js.data
func (j *Js) Get(key string) *Js {
	m := j.GetData()
	if v, ok := m[key]; ok {
		j.data = v
		return j
	}
	j.data = nil
	return j
}

//return json data
func (j *Js) GetData() map[string]interface{} {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

func (j *Js) GetIndex(i int) *Js {
	num := i - 1
	if m, ok := (j.data).([]interface{}); ok {
		v := m[num]
		j.data = v
		return j
	}

	{
		if m, ok := (j.data).(map[string]interface{}); ok {
			var n = 0
			var data = make(map[string]interface{})
			for i, v := range m {
				if n == num {
					switch vv := v.(type) {
					case float64:
						data[i] = strconv.FormatFloat(vv, 'f', -1, 64)
						j.data = data
						return j
					case string:
						data[i] = vv
						j.data = data
						return j
					case []interface{}:
						j.data = vv
						return j
					}
				}
				n++
			}
		}
	}
	j.data = nil
	return j
}

//must be []interface
func (j *Js) ArrayIndex(i int) string {
	num := i - 1
	if i > len((j.data).([]interface{})) {
		data := errors.New("index out of range list").Error()
		return data
	}
	if m, ok := (j.data).([]interface{}); ok {
		v := m[num]
		switch vv := v.(type) {
		case float64:
			return strconv.FormatFloat(vv, 'f', -1, 64)
		case string:
			return vv
		default:
			return ""
		}
	}

	if _, ok := (j.data).(map[string]interface{}); ok {
		return "error"
	}
	return "error"

}

//the data must be []interface{}, According to your custom number to return key adn array data
func (j *Js) GetKey(key string, i int) *Js {
	num := i - 1
	if i > len((j.data).([]interface{})) {
		j.data = errors.New("index out of range list").Error()
		return j
	}
	if m, ok := (j.data).([]interface{}); ok {
		v := m[num].(map[string]interface{})
		if h, ok := v[key]; ok {
			j.data = h
			return j
		}
	}
	j.data = nil
	return j
}

//According to the custom of the PATH to fing PATH
func (j *Js) GetPath(args ...string) *Js {
	d := j
	for i := range args {
		m := d.GetData()

		if val, ok := m[args[i]]; ok {
			d.data = val
		} else {
			d.data = nil
			return d
		}
	}
	return d
}

func (j *Js) ToString() string {
	if m, ok := j.data.(string); ok {
		return m
	}
	if m, ok := j.data.(float64); ok {
		return strconv.FormatFloat(m, 'f', -1, 64)
	}
	return ""
}

func (j *Js) ToArray() (k, d []string) {
	var key, data []string
	if m, ok := (j.data).([]interface{}); ok {
		for _, value := range m {
			for index, v := range value.(map[string]interface{}) {
				switch vv := v.(type) {
				case float64:
					data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
					key = append(key, index)
				case string:
					data = append(data, vv)
					key = append(key, index)
				}
			}
		}
		return key, data
	}
	if m, ok := (j.data).(map[string]interface{}); ok {
		for index, v := range m {
			switch vv := v.(type) {
			case float64:
				data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
				key = append(key, index)
			case string:
				data = append(data, vv)
				key = append(key, index)
			}
		}
		return key, data
	}
	return nil, nil
}

func (j *Js) Array() ([]interface{}, error) {
	if a, ok := (j.data).([]interface{}); ok {
		return a, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

func (j *Js) StringtoArray() []string {
	var data []string
	for _, v := range j.data.([]interface{}) {
		switch vv := v.(type) {
		case string:
			data = append(data, vv)
		case float64:
			data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
		}
	}
	return data
}

//for test
func (j *Js) Type() {
	fmt.Println(reflect.TypeOf(j.data))
}
