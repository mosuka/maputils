// Copyright (c) 2019 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maputils

import (
	"reflect"
	"testing"

	"github.com/imdario/mergo"
)

func TestGet(t *testing.T) {
	data := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"d": 4,
		},
		"c": []string{
			"A", "B",
		},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}

	// path = "/"
	m, err := NewNestedMap(data)
	if err != nil {
		t.Errorf("%v", err)
	}

	val1, err := m.Get("/")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp1 := data
	act1 := val1
	if !reflect.DeepEqual(exp1, act1) {
		t.Errorf("expected content to see %v, saw %v", exp1, act1)
	}

	// path = "/a"
	val2, err := m.Get("/a")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp2 := 1
	act2 := val2
	if exp2 != act2 {
		t.Errorf("expected content to see %v, saw %v", exp2, act2)
	}

	// path = "/b"
	val3, err := m.Get("/b")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp3 := map[string]interface{}{"d": 4}
	act3 := val3
	if !reflect.DeepEqual(exp3, act3) {
		t.Errorf("expected content to see %v, saw %v", exp3, act3)
	}

	// path = "/c"
	val4, err := m.Get("/c")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp4 := []string{"A", "B"}
	act4 := val4
	if !reflect.DeepEqual(exp4, act4) {
		t.Errorf("expected content to see %v, saw %v", exp4, act4)
	}

	// path = "/b/d"
	val5, err := m.Get("/b/d")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp5 := 4
	act5 := val5
	if exp5 != act5 {
		t.Errorf("expected content to see %v, saw %v", exp5, act5)
	}

	// path = "/e"
	val6, err := m.Get("/e")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp6 := []interface{}{
		map[string]interface{}{"f": "F"},
		map[string]interface{}{"g": 10.5},
	}
	act6 := val6
	if !reflect.DeepEqual(exp6, act6) {
		t.Errorf("expected content to see %v, saw %v", exp6, act6)
	}

	// path = "/e[0]"
	val7, err := m.Get("/e[0]")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp7 := map[string]interface{}{"f": "F"}
	act7 := val7
	if !reflect.DeepEqual(exp7, act7) {
		t.Errorf("expected content to see %v, saw %v", exp7, act7)
	}

	// path = "/e[1]"
	val8, err := m.Get("/e[1]")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp8 := map[string]interface{}{"g": 10.5}
	act8 := val8
	if !reflect.DeepEqual(exp8, act8) {
		t.Errorf("expected content to see %v, saw %v", exp8, act8)
	}

	// path = "/e[0]/f"
	val9, err := m.Get("/e[0]/f")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp9 := "F"
	act9 := val9
	if exp9 != act9 {
		t.Errorf("expected content to see %v, saw %v", exp9, act9)
	}

	// path = "/e[1]/g"
	val10, err := m.Get("/e[1]/g")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp10 := 10.5
	act10 := val10
	if exp10 != act10 {
		t.Errorf("expected content to see %v, saw %v", exp10, act10)
	}

	// path = "/h"
	val11, err := m.Get("/h")
	if err == nil {
		t.Errorf("expected errors to occur: %v", err)
	}
	if val11 != nil {
		t.Errorf("expected nil to occur: %v", val11)
	}

	// path = "/e[2]"
	val12, err := m.Get("/e[2]")
	if err == nil {
		t.Errorf("expected errors to occur: %v", err)
	}
	if val12 != nil {
		t.Errorf("expected nil to occur: %v", val12)
	}

	// src = "aaaaaaaa"
	_, err = NewNestedMap("aaaaaaaa")
	if err == nil {
		t.Errorf("expected errors to occur: %v", err)
	}
}

func TestMakeMap(t *testing.T) {
	m, err := NewNestedMap(make(map[string]interface{}, 0))
	if err != nil {
		t.Errorf("%v", err)
	}

	// "/a": "A"
	val1 := m.makeMap("/a", "A")
	exp1 := map[string]interface{}{"a": "A"}
	act1 := val1
	if !reflect.DeepEqual(exp1, act1) {
		t.Errorf("expected content to see %v, saw %v", exp1, act1)
	}

	// "/a/b": "AB"
	val2 := m.makeMap("/a/b", "AB")
	exp2 := map[string]interface{}{
		"a": map[string]interface{}{
			"b": "AB",
		},
	}
	act2 := val2
	if !reflect.DeepEqual(exp2, act2) {
		t.Errorf("expected content to see %v, saw %v", exp2, act2)
	}
}

func TestSet(t *testing.T) {
	data := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"d": 4,
		},
		"c": []string{
			"A", "B",
		},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}

	m, err := NewNestedMap(data)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = m.Set("/a", "A")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp1 := map[string]interface{}{
		"a": "A",
		"b": map[string]interface{}{
			"d": 4,
		},
		"c": []string{
			"A", "B",
		},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}
	act1 := m.data
	if !reflect.DeepEqual(exp1, act1) {
		t.Errorf("expected content to see %v, saw %v", exp1, act1)
	}

	err = m.Set("/b", "B")
	if err != nil {
		t.Errorf("%v", err)
	}
	exp2 := map[string]interface{}{
		"a": "A",
		"b": "B",
		"c": []string{
			"A", "B",
		},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}
	act2 := m.data
	if !reflect.DeepEqual(exp2, act2) {
		t.Errorf("expected content to see %v, saw %v", exp2, act2)
	}

	err = m.Set("/c", map[string]interface{}{"d": "D"})
	if err != nil {
		t.Errorf("%v", err)
	}
	exp4 := map[string]interface{}{
		"a": "A",
		"b": "B",
		"c": map[string]interface{}{"d": "D"},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}
	act4 := m.data
	if !reflect.DeepEqual(exp4, act4) {
		t.Errorf("expected content to see %v, saw %v", exp4, act4)
	}

	err = m.Set("/h/i/j/k", map[string]interface{}{"l": "L"})
	if err != nil {
		t.Errorf("%v", err)
	}
	exp5 := map[string]interface{}{
		"a": "A",
		"b": "B",
		"c": map[string]interface{}{"d": "D"},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
		"h": map[string]interface{}{"i": map[string]interface{}{"j": map[string]interface{}{"k": map[string]interface{}{"l": "L"}}}},
	}
	act5 := m.data
	if !reflect.DeepEqual(exp5, act5) {
		t.Errorf("expected content to see %v, saw %v", exp5, act5)
	}
}

func TestDel(t *testing.T) {
	data := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"d": 4,
		},
		"c": []string{
			"A", "B",
		},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}

	m, err := NewNestedMap(data)
	if err != nil {
		t.Errorf("%v", err)
	}

	m.Delete("/a")
	exp1 := data
	act1 := map[string]interface{}{
		"b": map[string]interface{}{
			"d": 4,
		},
		"c": []string{
			"A", "B",
		},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}
	if !reflect.DeepEqual(exp1, act1) {
		t.Errorf("expected content to see %v, saw %v", exp1, act1)
	}

}

func TestMergo(t *testing.T) {
	dst := map[string]interface{}{
		"a": 1,
		"b": "B",
		"c": []string{"A", "B"},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}

	err := mergo.Merge(&dst, map[string]interface{}{"a": "A"}, mergo.WithOverride)
	if err != nil {
		t.Errorf("%v", err)
	}
	exp1 := map[string]interface{}{
		"a": "A",
		"b": "B",
		"c": []string{"A", "B"},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
	}
	act1 := dst
	if !reflect.DeepEqual(exp1, act1) {
		t.Errorf("expected content to see %v, saw %v", exp1, act1)
	}

	err = mergo.Merge(&dst, map[string]interface{}{"h": "H"}, mergo.WithOverride)
	if err != nil {
		t.Errorf("%v", err)
	}
	exp2 := map[string]interface{}{
		"a": "A",
		"b": "B",
		"c": []string{"A", "B"},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
		"h": "H",
	}
	act2 := dst
	if !reflect.DeepEqual(exp2, act2) {
		t.Errorf("expected content to see %v, saw %v", exp2, act2)
	}

	err = mergo.Merge(&dst, map[string]interface{}{"i": map[string]interface{}{"j": "J"}}, mergo.WithOverride)
	if err != nil {
		t.Errorf("%v", err)
	}
	exp3 := map[string]interface{}{
		"a": "A",
		"b": "B",
		"c": []string{"A", "B"},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
		"h": "H",
		"i": map[string]interface{}{"j": "J"},
	}
	act3 := dst
	if !reflect.DeepEqual(exp3, act3) {
		t.Errorf("expected content to see %v, saw %v", exp3, act3)
	}

	err = mergo.Merge(&dst, map[string]interface{}{"i": map[string]interface{}{"k": "K"}}, mergo.WithOverride)
	if err != nil {
		t.Errorf("%v", err)
	}
	exp4 := map[string]interface{}{
		"a": "A",
		"b": "B",
		"c": []string{"A", "B"},
		"e": []interface{}{
			map[string]interface{}{"f": "F"},
			map[string]interface{}{"g": 10.5},
		},
		"h": "H",
		"i": map[string]interface{}{"j": "J", "k": "K"},
	}
	act4 := dst
	if !reflect.DeepEqual(exp4, act4) {
		t.Errorf("expected content to see %v, saw %v", exp4, act4)
	}

	//err = mergo.Merge(&dst, map[string]interface{}{"h": map[string]interface{}{"l": "L"}})
	//if err != nil {
	//	t.Errorf("%v", err)
	//}
	//exp5 := map[string]interface{}{
	//	"a": "A",
	//	"b": "B",
	//	"c": []string{"A", "B"},
	//	"e": []interface{}{
	//		map[string]interface{}{"f": "F"},
	//		map[string]interface{}{"g": 10.5},
	//	},
	//	"h": map[string]interface{}{"l": "L"},
	//	"i": map[string]interface{}{"j": "J", "k": "K"},
	//}
	//act5 := dst
	//if !reflect.DeepEqual(exp5, act5) {
	//	t.Errorf("expected content to see %v, saw %v", exp5, act5)
	//}

	//src := map[string]interface{}{
	//	"c": map[string]interface{}{"d": "D", "e": "E"},
	//}
	//
	//err := mergo.Merge(&dst, src, mergo.WithOverride)
	//if err != nil {
	//	t.Errorf("%v", err)
	//}
	//exp1 := map[string]interface{}{
	//	"a": "A",
	//	"b": "B",
	//	"c": map[string]interface{}{"d": "D", "e": "E"},
	//	"e": []interface{}{
	//		map[string]interface{}{"f": "F"},
	//		map[string]interface{}{"g": 10.5},
	//	},
	//}
	//act1 := dst
	//if !reflect.DeepEqual(exp1, act1) {
	//	t.Errorf("expected content to see %v, saw %v", exp1, act1)
	//}
}

func TestIterator(t *testing.T) {
	itr := newIterator([]string{"A", "B", "C"})

	value, err := itr.value()
	if err != nil {
		t.Errorf("%v", err)
	}
	if value != "A" {
		t.Errorf("expected content to see %v, saw %v", "A", value)
	}

	hasNext := itr.hasNext()
	if !hasNext {
		t.Errorf("expected content to see %v, saw %v", true, hasNext)
	}

	itr.next()

	value, err = itr.value()
	if err != nil {
		t.Errorf("%v", err)
	}
	if value != "B" {
		t.Errorf("expected content to see %v, saw %v", "B", value)
	}

	hasNext = itr.hasNext()
	if !hasNext {
		t.Errorf("expected content to see %v, saw %v", true, hasNext)
	}

	itr.next()

	value, err = itr.value()
	if err != nil {
		t.Errorf("%v", err)
	}
	if value != "C" {
		t.Errorf("expected content to see %v, saw %v", "B", value)
	}

	hasNext = itr.hasNext()
	if hasNext {
		t.Errorf("expected content to see %v, saw %v", false, hasNext)
	}

	itr.next()

	value, err = itr.value()
	if err == nil {
		t.Errorf("expected errors to occur: %v", err)
	}
	if value != "" {
		t.Errorf("expected empty string to occur: %v", value)
	}
}
