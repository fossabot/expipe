// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package tools

import (
	"fmt"
	"testing"
)

func TestStringInSlice(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		niddle   string
		haystack []string
		result   bool
	}{
		{"aaa", []string{"aaa", "bbb"}, true},
		{"aaa", []string{"aaa", "aaa"}, true},
		{"aaa", []string{"bbb"}, false},
		{"aaa", []string{}, false},
		{"aaa", []string{"aaaa"}, false},
		{"aaa", []string{"AAA"}, false},
	}
	for i, tc := range tcs {
		name := fmt.Sprintf("case_%d", i)
		t.Run(name, func(t *testing.T) {
			if ok := StringInSlice(tc.niddle, tc.haystack); ok != tc.result {
				t.Errorf("want (%t), got (%t)", tc.result, ok)
			}
		})
	}
}

func TestStringInMapKeys(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		niddle   string
		haystack map[string]string
		result   bool
	}{
		{"aaa", map[string]string{"aaa": "a"}, true},
		{"aaa", map[string]string{"aaa": "a", "bbbb": "a"}, true},
		{"aaa", map[string]string{"bbb": "a"}, false},
		{"aaa", map[string]string{"aaaa": "a"}, false},
		{"aaa", map[string]string{"AAA": "a"}, false},
	}
	for i, tc := range tcs {
		name := fmt.Sprintf("case_%d", i)
		t.Run(name, func(t *testing.T) {
			if ok := StringInMapKeys(tc.niddle, tc.haystack); ok != tc.result {
				t.Errorf("want (%t), got (%t)", tc.result, ok)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "empty object", input: `{}`, want: true},
		{name: "string key", input: `{"sss": 1}`, want: true},
		{name: "string key value", input: `{"sss": "tt"}`, want: true},
		{name: "number", input: `666`, want: true},
		{name: "string", input: `"666"`, want: true},
		{name: "no ending", input: `{"sss": 666`, want: false},
		{name: "no object", input: `"sss": 666`, want: false},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsJSON([]byte(tc.input)); got != tc.want {
				t.Errorf("IsJSON() = (%t); want (%t)", got, tc.want)
			}
		})
	}
}
