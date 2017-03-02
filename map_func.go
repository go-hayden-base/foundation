package foundation

import (
	"errors"
	"log"
	"reflect"
)

// --- Merge ----

// MapMerge merge map 'from' to map 'to'
// 'from' and 'to' must be reflect.Map kind
// and must have same reflect.Type
func MapMerge(from interface{}, to interface{}) error {
	if from == nil || to == nil {
		return nil
	}
    
	if reflect.TypeOf(from).Kind() != reflect.Map {
		return errors.New("from argument is not a map")
	}
	if reflect.TypeOf(to).Kind() != reflect.Map {
		return errors.New("to argument is not a map")
	}
	if reflect.TypeOf(from) != reflect.TypeOf(to) {
		return errors.New("from and to argument's type is not same")
	}

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if rec := recover(); rec != nil {
			log.Println("foundation.MapMerge发生严重错误:", rec)
		} else {
			log.Println("foundation.MapMerge发生未知的严重错误!")
		}
	}()
	fromValue := reflect.ValueOf(from)
	toValue := reflect.ValueOf(to)
	keys := fromValue.MapKeys()
	for _, kv := range keys {
		v := fromValue.MapIndex(kv)
		if v.IsValid() {
			toValue.SetMapIndex(kv, v)
		}
	}
	return nil
}
