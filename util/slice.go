package util

import (
	"strings"

	"github.com/spf13/cast"
)

// ConvertStringSliceToInt64Slice --
func ConvertStringSliceToInt64Slice(in []string) []int64 {
	res := make([]int64, 0)
	for i := range in {
		res = append(res, cast.ToInt64(in[i]))
	}
	return res
}

// ConvertInt64SliceToStringSlice --
func ConvertInt64SliceToStringSlice(in []int64) []string {
	res := make([]string, 0)
	for i := range in {
		res = append(res, cast.ToString(in[i]))
	}
	return res
}

// IsInt64SliceContains --
func IsInt64SliceContains(sl []int64, s int64) bool {
	for _, i := range sl {
		if i == s {
			return true
		}
	}
	return false
}

// IsStringSliceContains --
func IsStringSliceContains(sl []string, s string) bool {
	for _, i := range sl {
		if i == s {
			return true
		}
	}

	return false
}

// IntersectionStringSlice --
func IntersectionStringSlice(s1 []string, s2 []string) []string {
	r := []string{}
	mx := map[string]bool{}
	for _, x := range s2 {
		mx[x] = true
	}
	for _, x := range s1 {
		if _, ok := mx[x]; ok {
			r = append(r, x)
		}
	}
	return r
}

// IntersectionInt64Slice --
func IntersectionInt64Slice(s1 []int64, s2 []int64) []int64 {
	r := []int64{}
	mx := map[int64]bool{}
	for _, x := range s2 {
		mx[x] = true
	}
	for _, x := range s1 {
		if _, ok := mx[x]; ok {
			r = append(r, x)
		}
	}
	return r
}

// TrimedSpaceStringSlice -- slice string by seperate and trim space
func TrimedSpaceStringSlice(s, sep string) []string {
	var sl []string

	for _, p := range strings.Split(s, sep) {
		if str := strings.TrimSpace(p); len(str) > 0 {
			sl = append(sl, strings.TrimSpace(p))
		}
	}

	return sl
}

// IsEqualStringSlice --
func IsEqualStringSlice(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	ms1 := make(map[string]bool)
	for _, i1 := range s1 {
		ms1[i1] = true
	}

	for _, i2 := range s2 {
		if exist, ok := ms1[i2]; !exist || !ok {
			return false
		}
	}

	return true
}
