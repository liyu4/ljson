// ljson provide paerses the json string  and stores the result in the value pointed to by &f
// but it is convenient to get map or struct in Go.
package ljson

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
func NewJson(data string) *Js {
	j := new(Js)
	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	if nil != err {
		return j
	}
	j.data = f
	return j
}

// Acoording to the key of the returned data information , return js.data
// if you know json is an object
func (j *Js) Get(key string) *Js {
	m := j.GetMapData()
	if v, ok := m[key]; ok {
		j.data = v
		return j
	}
	j.data = nil
	return j
}

// return map in Go
func (j *Js) GetMapData() map[string]interface{} {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

// GetIndex get []interface or map in Go
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

// you must know this an  []interface so that you can get element
func (j *Js) ArrayIndex(i int) (string, error) {
	num := i - 1
	if i > len((j.data).([]interface{})) {
		return "", errors.New("index out of range list")
	}

	if m, ok := (j.data).([]interface{}); ok {
		v := m[num]
		switch vv := v.(type) {
		case float64:
			return strconv.FormatFloat(vv, 'f', -1, 64), nil
		case string:
			return vv, nil
		default:
			return "", nil
		}
	}

	if _, ok := (j.data).(map[string]interface{}); ok {
		return "", errors.New("json object must be was array")
	}
	return "", errors.New("Unkonw error")
}

// The data must be []interface{}, According to your custom number to return key adn array data
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

// According to the custom of the PATH to fing element
// You can use function this to find recursive map
func (j *Js) GetPath(args ...string) *Js {
	d := j
	for i := range args {
		m := d.GetMapData()

		if val, ok := m[args[i]]; ok {
			d.data = val
		} else {
			d.data = nil
			return d
		}
	}
	return d
}

// String return string
func (j *Js) String() string {
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

// Array return array
func (j *Js) Array() ([]string, error) {
	if a, ok := (j.data).([]interface{}); ok {
		array := make([]string, 0)
		for _, v := range a {
			switch vv := v.(type) {
			case float64:
				array = append(array, strconv.FormatFloat(vv, 'f', -1, 64))
			case string:
				array = append(array, vv)
			}

		}
		return array, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

//for test
func (j *Js) Type() {
	fmt.Println(reflect.TypeOf(j.data))
}
