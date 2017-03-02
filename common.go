package foundation

import (
	"errors"
	"reflect"
)

// Contains return target contains obj or not
// The function use reflect, it is not a high performance function
func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
		return false, nil
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, errors.New("target is not a slice, array or map")
}
