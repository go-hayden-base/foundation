package foundation

// IEnumerable is enumerable interface
type IEnumerable interface {
	Filter(f func(item interface{}) bool) IEnumerable
	Enumerate(f func(item interface{}, err error, stop *bool))
}

// IContain is interface that has method Contains
type IContain interface {
	Contains(item interface{}) bool
}
