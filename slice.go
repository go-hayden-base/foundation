package foundation

type StringSlice []string

func (s StringSlice) Contains(item interface{}) bool {
	if strItem, ok := item.(string); ok {
		for _, str := range s {
			if str == strItem {
				return true
			}
		}
	}
	return false
}

type IntSlice []int

func (s IntSlice) Contains(item interface{}) bool {
	if intItem, ok := item.(int); ok {
		for _, i := range s {
			if i == intItem {
				return true
			}
		}
	}
	return false
}

type FloatSlice []float64

func (s FloatSlice) Contains(item interface{}) bool {
	if floatItem, ok := item.(float64); ok {
		for _, f := range s {
			if f == floatItem {
				return true
			}
		}
	}
	return false
}
