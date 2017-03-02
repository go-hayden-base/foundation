package foundation

import (
	"testing"
)

func TestContains(t *testing.T) {
	arrStr := []string{"a", "b", "c"}
	arrInt := []int{1, 2, 3}
	strA, errStrA := Contains("b", arrStr)
	strB, errStrB := Contains("d", arrStr)
	intA, errIntA := Contains(3, arrInt)
	intB, errIntB := Contains(5, arrInt)
	_, errA := Contains(1, 2)
	if errA == nil {
		t.Error("type error")
	} else if errStrA != nil || errStrB != nil || errIntA != nil || errIntB != nil {
		t.Error(errStrA, errStrB, errIntA, errIntB)
	} else if !strA || strB || !intA || intB {
		t.Error(strA, strB, intA, intB)
	}
}
