package lrucache

import "errors"

var testLrucacheTestSuite = []struct {
	name          string
	capacity      int
	expectedCache *Cache
	expectedError error
}{
	{"BasicFuntionality", 1, createCache(1), nil},
	{"LargeCache", 10000000, createCache(10000000), nil},
	{"NoCapacity", 0, createCache(0), errors.New("capacity must be greater than zero")},
}

var testClearTestSuite = []struct {
	name          string
	baseCache     *Cache
	expectedCache *Cache
}{
	{"BasicFuntionality", fillCache(3, []tuple{{"a", 1}, {"b", 2}, {"c", 3}}), createCache(3)},
	{"SingleElement", fillCache(1, []tuple{{"a", 1}}), createCache(1)},
	{"EmptyCache", fillCache(1, []tuple{}), createCache(1)},
}

var testSetTestSuite = []struct {
	name          string
	baseCache     *Cache
	expectedTuple tuple
}{
	{"SingleElement", fillCache(1, []tuple{{"a", 1}}), tuple{"a", 1}},
	{"BasicFunctionality", fillCache(3, []tuple{{"a", 1}, {"b", 2}, {"c", 3}}), tuple{"c", 3}},
	{"OverwriteElements", fillCache(3, []tuple{{"a", 1}, {"b", 2}, {"a", 3}}), tuple{"a", 3}},
	{"DuplicateElements", fillCache(1, []tuple{{"a", 1}, {"a", 1}, {"a", 1}}), tuple{"a", 1}},
}

var testGetTestSuite = []struct {
	name          string
	baseCache     *Cache
	key           string
	expectedValue interface{}
}{
	{"BasicFuntionality", fillCache(3, []tuple{{"a", 1}, {"b", 2}, {"c", 3}}), "b", 2},
	{"DuplicateElements", fillCache(1, []tuple{{"a", 1}, {"a", 1}, {"a", 1}}), "a", 1},
	{"RewrittenElements", fillCache(1, []tuple{{"a", 1}, {"a", 2}, {"a", 3}}), "a", 3},
	{"ElementNotPresent", fillCache(3, []tuple{{"a", 1}, {"b", 2}, {"c", 3}}), "z", nil},
	{"SingleElement", fillCache(1, []tuple{{"a", 1}}), "a", 1},
	{"EmptyCache", fillCache(3, []tuple{}), "b", nil},
}
