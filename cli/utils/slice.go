package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"reflect"
)

// StringSliceMakeUnique removes duplicate strings from a slice
func StringSliceMakeUnique(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

// BoolSliceMakeUnique removes duplicate booleans from a slice
func BoolSliceMakeUnique(elements []bool) []bool {
	encountered := map[bool]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []bool{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func HashOfStringSlice(arr []string) string {
	arrBytes := []byte{}
	for _, item := range arr {
		jsonBytes, _ := json.Marshal(item)
		arrBytes = append(arrBytes, jsonBytes...)
	}
	bytes := md5.Sum(arrBytes)
	return hex.EncodeToString(bytes[:16])
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
