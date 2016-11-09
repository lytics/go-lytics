package lytics

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Given a set of bytes, convert to json map[string]interface{}, then
// convert
func FlattenJson(jsonBytes []byte) (map[string][]string, error) {
	jsonMap := make(map[string]interface{})
	if err := json.Unmarshal(jsonBytes, &jsonMap); err != nil {
		return nil, err
	}
	return FlattenJsonMap(jsonMap)
}

// Attempt to flatten A map of map[string]interface{} into url Values
func FlattenJsonMapIntoQs(uv url.Values, jsonMap map[string]interface{}, pre string) error {
	for k, v := range jsonMap {
		if err := flattenJsonValue(uv, v, pre+k); err != nil {
			return err
		}
	}
	return nil
}

// The map m should be a map that resulted from json.Unmarshal(), or the behavior is
// undefined (that's bad).
func FlattenJsonMap(m map[string]interface{}) (map[string][]string, error) {
	outMap := make(map[string][]string, len(m))
	if err := flattenJson("", outMap, m); err != nil {
		return nil, err
	}
	return outMap, nil
}

func flattenJson(prefix string, into map[string][]string, toFlatten interface{}) error {
	typ := reflect.TypeOf(toFlatten)
	kind := typ.Kind()
	rVal := reflect.ValueOf(toFlatten)

	switch {
	case isScalar(kind):
		into[prefix] = []string{scalarToString(&rVal)}
		return nil
	case kind == reflect.Slice:
		sl := toFlatten.([]interface{})

		containsOnlyScalars := true

		for _, elem := range sl {
			if !isScalar(reflect.ValueOf(elem).Type().Kind()) {
				containsOnlyScalars = false
			}
		}

		if containsOnlyScalars {
			into[prefix] = scalarSliceToStrings(&rVal)
		} else {
			for i := 0; i < len(sl); i++ {
				recursePrefix := strcat(prefix, "[", strconv.Itoa(i), "]")
				err := flattenJson(recursePrefix, into, sl[i])
				if err != nil {
					return err
				}
			}
		}

	case kind == reflect.Map:
		for _, mapKey := range rVal.MapKeys() {
			mapVal := rVal.MapIndex(mapKey)

			var recursePrefix string
			if prefix == "" {
				recursePrefix = mapKey.String()
			} else {
				recursePrefix = strcat(prefix, ".", mapKey.String())
			}

			if mapVal.IsNil() {
				into[recursePrefix] = []string{}
			} else {
				if err := flattenJson(recursePrefix, into, mapVal.Interface()); err != nil {
					return err
				}
			}

		}
	default:
		panic(fmt.Sprintf("Unexpected type %s kind %s", rVal.Type(), kind))
	}

	return nil
}

// The input must be a reflect.Value that is a slice whose elements are scalars.
// Only works for scalar types that would be produced by JSON decoding into a
// map[string]interface{}.
func scalarSliceToStrings(rVal *reflect.Value) []string {
	sliceLen := rVal.Len()
	strs := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		sliceElem := rVal.Index(i).Elem()
		strs[i] = scalarToString(&sliceElem)
	}
	return strs
}

// Only works for scalar types that would be produced by JSON decoding into a
// map[string]interface{}.
func scalarToString(rVal *reflect.Value) string {
	switch rVal.Type().Kind() {
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(rVal.Float(), 'f', -1, 64)

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", rVal.Interface())

	case reflect.String:
		return rVal.String()

	case reflect.Bool:
		if rVal.Bool() {
			return "true"
		} else {
			return "false"
		}
	}
	panic(fmt.Sprintf("Unexpected scalar type %s", rVal.Type()))
}

// Only works for types that would be produced by JSON decoding into a
// map[string]interface{}.
func isScalar(k reflect.Kind) bool {
	switch k {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Bool,
		reflect.String, reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func strcat(strs ...string) string {
	return strings.Join(strs, "")
}

func flattenJsonValue(into url.Values, toFlatten interface{}, key string) error {
	if toFlatten == nil {
		// ??
		// into.Set(key,"")
		return nil
	}
	typ := reflect.TypeOf(toFlatten)
	kind := typ.Kind()
	rVal := reflect.ValueOf(toFlatten)

	if isScalar(kind) {
		into[key] = []string{scalarToString(&rVal)}
		return nil
	}
	switch val := toFlatten.(type) {
	// case []interface{}:
	// 	sva := make([]string, 0)
	// 	for _, av := range x {
	// 		if err := flattenJsonValue(n, av, key); err != nil {
	// 			return nil
	// 		}
	// 	}
	// 	if len(sva) > 0 {
	// 		uv[key] = sva
	// 	}
	case map[string]interface{}:
		if len(val) > 0 {
			if err := FlattenJsonMapIntoQs(into, val, key+"."); err != nil {
				return err
			}
		}
	case bool:
		if val == true {
			into.Set(key, "t")
		} else {
			into.Set(key, "f")
		}
	default:
		// what types don't we support?
		//  []interface{}
	}
	return nil
}
