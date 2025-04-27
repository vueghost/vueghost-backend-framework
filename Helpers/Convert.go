/*
 * Vueghost Tech, FZE © 2018- 2019. Unauthorized copying of this file,
 * via any medium is strictly prohibited proprietary and confidential,
 * can not be copied and/or distributed without the express permission of Vueghost
 */

package Helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func StringToBool(value string) bool {
	r, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return r
}

func StringToFloat(value string) float64 {
	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		res = 0.0
	}
	return res
}

func InterfaceToJsonString(obj interface{}) string {
	r, err := json.Marshal(obj)
	if err != nil {
		r = []byte("")
	}
	return string(r)
}

func InterfaceToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)
	return string(res)
}

func StringArrayToInterface(array []string) []interface{} {
	s := make([]interface{}, len(array))
	for i, v := range array {
		s[i] = v
	}
	return s
}

func ToFloat(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		res = 0.0
	}
	return res, err
}
func ToInt(value interface{}) (res int64, err error) {
	val := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		res = int64(val.Uint())
	case string:

	default:
		err = fmt.Errorf("math: square root of negative number %g", value)
		res = 0
	}

	return
}
func ToJSON(obj interface{}) (string, error) {
	res, err := json.Marshal(obj)
	if err != nil {
		res = []byte("")
	}
	return string(res), err
}
func ToBoolean(str interface{}) (bool, error) {
	return strconv.ParseBool(str.(string))
}

func StringToInt(value string, returnAsString bool) interface{} {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		if returnAsString {
			return "0"
		} else {
			return 0
		}

	}
	if returnAsString {
		return fmt.Sprintf("%v", v)
	} else {
		return v
	}
}
