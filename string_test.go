package foundation

import (
	"testing"
)

func TestStrMD5(t *testing.T) {
	origin := "this is origin string"
	md5 := "2086146e78bd0248fec3f88c4489c827"
	if StrMD5(origin) != md5 {
		t.Error(md5)
	}
}

func TestStrRandomMD5(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 10000; i++ {
		s := StrRandomMD5()
		if _, ok := m[s]; ok {
			t.Error("TestStrRandomMD5 Hit:", i+1)
			break
		} else {
			m[s] = true
		}
	}
}

func TestStrEnumerate(t *testing.T) {
	origin := "a/b/c/d/e/f/g/h"
	a := ""
	b := ""
	StrEnumerate(origin, "/", false, func(surplus string, current string, stop *bool) {
		a += (current + "/")
	})
	StrEnumerate(origin, "/", true, func(surplus string, current string, stop *bool) {
		b = current + "/" + b
	})
	if a != origin+"/" || b != origin+"/" {
		t.Error(a, b)
	}
}

func TestStrSplitFirst(t *testing.T) {
	originA := "a/b/c/d/e/f/g/h"
	originB := "/a/b/c"
	if a, b := StrSplitFirst(originA, "/"), StrSplitFirst(originB, "/"); a != "a" || b != "" {
		t.Error(a, b)
	}
}

func TestStrSplitLast(t *testing.T) {
	originA := "a/b/c/d/e/f/g/h"
	originB := "a/b/c/"
	if a, b := StrSplitLast(originA, "/"), StrSplitLast(originB, "/"); a != "h" || b != "" {
		t.Error(a, b)
	}
}
