package main

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var full = Set{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
	6: {},
	7: {},
	8: {},
	9: {},
}

type Set map[uint8]struct{}

func (s Set) First() uint8 {
	for key := range s {
		return key
	}

	return 0
}

func (s Set) Append(value uint8) {
	s[value] = struct{}{}
}

func (s Set) Erase(value uint8) {
	delete(s, value)
}

func (s Set) Contains(value uint8) bool {
	_, ok := s[value]
	return ok
}

func (s Set) Clear() {
	for key := range s {
		delete(s, key)
	}
}

func (s Set) Equal(other Set) bool {
	return reflect.DeepEqual(s, other)
}

func (s Set) Not() Set {
	res := make(Set)

	for val := range full {
		if !s.Contains(val) {
			res.Append(val)
		}
	}

	return res
}

func (s Set) Hash() string {
	keys := make([]string, 0, len(s))
	for key := range s {
		keys = append(keys, strconv.Itoa(int(key)))
	}

	sort.Strings(keys)

	return strings.Join(keys, ",")
}
