package collections

import (
	"reflect"
	"sort"
)

func SortedKeys(m interface{}) []string {
	value := reflect.ValueOf(m)
	if value.Kind() != reflect.Map {
		panic("m must be a map of string to some type")
	}
	if value.Type().Key().Kind() != reflect.String {
		panic("m must be a map of string to some type")
	}
	keyValues := value.MapKeys()
	keyStrings := make([]string, len(keyValues))
	for i := range keyValues {
		keyStrings[i] = keyValues[i].String()
	}
	sort.Strings(keyStrings)
	return keyStrings
}
