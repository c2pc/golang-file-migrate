package merger

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/c2pc/config-migrate/replacer"
)

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name     string
		oldMap   map[string]interface{}
		newMap   map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "Nil maps",
			oldMap:   nil,
			newMap:   nil,
			expected: map[string]interface{}{},
		},
		{
			name:     "Empty maps",
			oldMap:   map[string]interface{}{},
			newMap:   map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			name:     "New is nil, old is empty",
			oldMap:   map[string]interface{}{},
			newMap:   nil,
			expected: map[string]interface{}{},
		},
		{
			name:     "New is empty, old is nil",
			oldMap:   nil,
			newMap:   map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			name:   "Old is nil",
			oldMap: nil,
			newMap: map[string]interface{}{
				"foo": "bar",
			},
			expected: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name:   "Old is empty",
			oldMap: map[string]interface{}{},
			newMap: map[string]interface{}{
				"foo": "bar",
			},
			expected: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name: "New is empty",
			oldMap: map[string]interface{}{
				"foo": "bar",
			},
			newMap:   map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			name: "New is nil",
			oldMap: map[string]interface{}{
				"foo": "bar",
			},
			newMap:   nil,
			expected: map[string]interface{}{},
		},
		{
			name: "New string",
			oldMap: map[string]interface{}{
				"foo": "bar",
			},
			newMap: map[string]interface{}{
				"foo2": "bar",
			},
			expected: map[string]interface{}{
				"foo2": "bar",
			},
		},
		{
			name: "New string and old string",
			oldMap: map[string]interface{}{
				"foo": "bar",
			},
			newMap: map[string]interface{}{
				"foo":  "bar",
				"foo2": "bar",
			},
			expected: map[string]interface{}{
				"foo":  "bar",
				"foo2": "bar",
			},
		},
		{
			name: "New number",
			oldMap: map[string]interface{}{
				"foo": 1,
			},
			newMap: map[string]interface{}{
				"foo2": 2,
			},
			expected: map[string]interface{}{
				"foo2": 2,
			},
		},
		{
			name: "New number and old number",
			oldMap: map[string]interface{}{
				"foo": 1,
			},
			newMap: map[string]interface{}{
				"foo":  1,
				"foo2": 2,
			},
			expected: map[string]interface{}{
				"foo":  1,
				"foo2": 2,
			},
		},
		{
			name: "New bool",
			oldMap: map[string]interface{}{
				"foo": true,
			},
			newMap: map[string]interface{}{
				"foo2": false,
			},
			expected: map[string]interface{}{
				"foo2": false,
			},
		},
		{
			name: "New bool and old bool",
			oldMap: map[string]interface{}{
				"foo": true,
			},
			newMap: map[string]interface{}{
				"foo":  true,
				"foo2": false,
			},
			expected: map[string]interface{}{
				"foo":  true,
				"foo2": false,
			},
		},
		{
			name: "Old value is empty(bool)",
			oldMap: map[string]interface{}{
				"foo": nil,
			},
			newMap: map[string]interface{}{
				"foo": true,
			},
			expected: map[string]interface{}{
				"foo": nil,
			},
		},
		{
			name: "Old value is empty(string)",
			oldMap: map[string]interface{}{
				"foo": nil,
			},
			newMap: map[string]interface{}{
				"foo": "bar",
			},
			expected: map[string]interface{}{
				"foo": nil,
			},
		},
		{
			name: "Old value is empty(number)",
			oldMap: map[string]interface{}{
				"foo": nil,
			},
			newMap: map[string]interface{}{
				"foo": 1,
			},
			expected: map[string]interface{}{
				"foo": nil,
			},
		},
		{
			name: "New map",
			oldMap: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "bar",
				},
			},
			newMap: map[string]interface{}{
				"foo2": map[string]interface{}{
					"bar": "bar",
				},
			},
			expected: map[string]interface{}{
				"foo2": map[string]interface{}{
					"bar": "bar",
				},
			},
		},
		{
			name: "New map and old map",
			oldMap: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "bar",
				},
			},
			newMap: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "bar",
				},
				"foo2": map[string]interface{}{
					"bar": "bar",
				},
			},
			expected: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "bar",
				},
				"foo2": map[string]interface{}{
					"bar": "bar",
				},
			},
		},
		{
			name: "Old value is empty(map)",
			oldMap: map[string]interface{}{
				"foo": nil,
			},
			newMap: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "bar",
				},
			},
			expected: map[string]interface{}{
				"foo": nil,
			},
		},
		{
			name: "New array",
			oldMap: map[string]interface{}{
				"foo": []interface{}{
					"bar",
				},
			},
			newMap: map[string]interface{}{
				"foo2": []interface{}{
					"bar",
				},
			},
			expected: map[string]interface{}{
				"foo2": []interface{}{
					"bar",
				},
			},
		},
		{
			name: "New array and old array",
			oldMap: map[string]interface{}{
				"foo": []interface{}{
					"bar",
				},
			},
			newMap: map[string]interface{}{
				"foo": []interface{}{
					"bar",
				},
				"foo2": []interface{}{
					"bar",
				},
			},
			expected: map[string]interface{}{
				"foo": []interface{}{
					"bar",
				},
				"foo2": []interface{}{
					"bar",
				},
			},
		},
		{
			name: "Old value is empty(array)",
			oldMap: map[string]interface{}{
				"foo": nil,
			},
			newMap: map[string]interface{}{
				"foo": []interface{}{
					"bar",
				},
			},
			expected: map[string]interface{}{
				"foo": nil,
			},
		},
		{
			name: "New array of numbers",
			oldMap: map[string]interface{}{
				"foo": []int{
					1, 2, 3,
				},
			},
			newMap: map[string]interface{}{
				"foo": []int{
					1, 2, 3,
				},
				"foo2": []int{
					1, 2, 3, 4,
				},
			},
			expected: map[string]interface{}{
				"foo": []int{
					1, 2, 3,
				},
				"foo2": []int{
					1, 2, 3, 4,
				},
			},
		},
		{
			name: "New array of maps",
			oldMap: map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{"foo2": "bar"},
				},
			},
			newMap: map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{"foo2": "bar"},
				},
				"foo2": []interface{}{
					map[string]interface{}{"foo3": "bar"},
				},
			},
			expected: map[string]interface{}{
				"foo": []map[string]interface{}{
					{"foo2": "bar"},
				},
				"foo2": []map[string]interface{}{
					{"foo3": "bar"},
				},
			},
		},
		{
			name: "All types",
			oldMap: map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{
						"foo_1": "bar",
						"foo_2": 1,
						"foo_3": true,
					},
				},
				"foo2": map[string]interface{}{
					"foo2_1": "bar",
					"foo2_2": 1,
					"foo2_3": true,
				},
				"foo3": 1,
				"foo4": "bar",
				"foo5": true,
				"foo6": []interface{}{
					1, 2, 3,
				},
				"foo7": []interface{}{
					"bar", "bar2", "bar3",
				},
				"foo8": []interface{}{
					true, false, true,
				},
				"foo14": []interface{}{
					1, 2, 3,
				},
			},
			newMap: map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{
						"foo_1": "bar1",
						"foo_2": 2,
						"foo_3": false,
						"foo_4": "bar",
						"foo_5": 1,
						"foo_6": true,
					},
				},
				"foo2": map[string]interface{}{
					"foo2_1": "bar",
					"foo2_2": 2,
					"foo2_3": true,
					"foo2_4": "bar",
					"foo2_5": 1,
					"foo2_6": true,
				},
				"foo3": 1,
				"foo4": "bar2",
				"foo5": false,
				"foo6": []interface{}{
					map[string]interface{}{
						"foo6_1": "bar",
						"foo6_2": 1,
						"foo6_3": false,
					},
				},
				"foo7": map[string]interface{}{
					"foo7_1": "bar",
					"foo7_2": 1,
					"foo7_3": true,
				},
				"foo8":  1,
				"foo9":  "bar",
				"foo10": true,
				"foo11": []interface{}{
					1, 2, 3, 4, 5,
				},
				"foo12": []interface{}{
					"bar", "bar2", "bar3",
				},
				"foo13": []interface{}{
					true, false, true,
				},
				"foo14": []interface{}{
					1, 2, 3, 4, 5,
				},
			},
			expected: map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{
						"foo_1": "bar",
						"foo_2": 1,
						"foo_3": true,
						"foo_4": "bar",
						"foo_5": 1,
						"foo_6": true,
					},
				},
				"foo2": map[string]interface{}{
					"foo2_1": "bar",
					"foo2_2": 1,
					"foo2_3": true,
					"foo2_4": "bar",
					"foo2_5": 1,
					"foo2_6": true,
				},
				"foo3": 1,
				"foo4": "bar",
				"foo5": true,
				"foo6": []interface{}{
					map[string]interface{}{
						"foo6_1": "bar",
						"foo6_2": 1,
						"foo6_3": false,
					},
				},
				"foo7": map[string]interface{}{
					"foo7_1": "bar",
					"foo7_2": 1,
					"foo7_3": true,
				},
				"foo8":  1,
				"foo9":  "bar",
				"foo10": true,
				"foo11": []interface{}{
					1, 2, 3, 4, 5,
				},
				"foo12": []interface{}{
					"bar", "bar2", "bar3",
				},
				"foo13": []interface{}{
					true, false, true,
				},
				"foo14": []interface{}{
					1, 2, 3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Merge(tt.newMap, tt.oldMap)

			res, err := json.Marshal(result)
			if err != nil {
				t.Error(err)
			}
			exp, err := json.Marshal(tt.expected)
			if err != nil {
				t.Error(err)
			}

			if string(res) != string(exp) {
				t.Errorf("Expected\n %s, got\n %s", string(exp), string(res))
			}
		})
	}
}

func TestMergeReplace(t *testing.T) {
	replacer.Register("__test_replacer__", func() string {
		return "-TEST_STRING-"
	})

	tests := []struct {
		name     string
		oldMap   map[string]interface{}
		newMap   map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:   "New map",
			oldMap: map[string]interface{}{},
			newMap: map[string]interface{}{
				"foo":  "some_string__test_replacer__some_string",
				"foo2": "some_string__test_replacer__",
				"foo3": "__test_replacer__",
				"foo4": "__test_replacer__some_string",
			},
			expected: map[string]interface{}{
				"foo":  "some_string-TEST_STRING-some_string",
				"foo2": "some_string-TEST_STRING-",
				"foo3": "-TEST_STRING-",
				"foo4": "-TEST_STRING-some_string",
			},
		},
		{
			name:   "New array",
			oldMap: map[string]interface{}{},
			newMap: map[string]interface{}{
				"foo": []interface{}{"some_string__test_replacer__some_string"},
			},
			expected: map[string]interface{}{
				"foo": []interface{}{"some_string-TEST_STRING-some_string"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Merge(tt.newMap, tt.oldMap)

			res, err := json.Marshal(result)
			if err != nil {
				t.Error(err)
			}
			exp, err := json.Marshal(tt.expected)
			if err != nil {
				t.Error(err)
			}

			if string(res) != string(exp) {
				t.Errorf("Expected\n %s, got\n %s", string(exp), string(res))
			}
		})
	}
}

func BenchmarkMergeMapsFlat(b *testing.B) {
	newMap := map[string]interface{}{
		"string": "value",
		"int":    1,
		"bool":   true,
	}

	oldMap := map[string]interface{}{
		"string": "old_value",
		"int":    99,
		"bool":   false,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mergeMaps(newMap, oldMap)
	}
}

func BenchmarkMergeMapsNested(b *testing.B) {
	newMap := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Alice",
			"age":  30,
		},
	}

	oldMap := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Bob",
			"city": "New York",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mergeMaps(newMap, oldMap)
	}
}

func BenchmarkMergeMapsDeepNestedArray(b *testing.B) {
	newMap := map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{
				"id":   0,
				"name": "default",
			},
		},
	}

	oldMap := map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{
				"id":   1,
				"name": "item1",
			},
			map[string]interface{}{
				"id":   2,
				"name": "item2",
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mergeMaps(newMap, oldMap)
	}
}

func generateDeepNestedMap() map[string]interface{} {
	return map[string]interface{}{
		"level1": map[string]interface{}{
			"level2": map[string]interface{}{
				"level3": map[string]interface{}{
					"string": "value",
					"number": 123,
					"bool":   true,
					"array": []interface{}{
						map[string]interface{}{
							"id":   1,
							"name": "item1",
							"meta": map[string]interface{}{
								"enabled": true,
								"tags":    []interface{}{"a", "b", "c"},
							},
						},
						map[string]interface{}{
							"id":   2,
							"name": "item2",
							"meta": map[string]interface{}{
								"enabled": false,
								"tags":    []interface{}{},
							},
						},
					},
				},
			},
		},
	}
}

func BenchmarkMergeMaps_DeepNestedLarge(b *testing.B) {
	newMap := generateDeepNestedMap()
	oldMap := generateDeepNestedMap()

	// Изменим значения в oldMap для "реального" слияния
	oldMap["level1"].(map[string]interface{})["level2"].(map[string]interface{})["level3"].(map[string]interface{})["string"] = "oldValue"
	oldMap["level1"].(map[string]interface{})["level2"].(map[string]interface{})["level3"].(map[string]interface{})["new_key"] = "added"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mergeMaps(newMap, oldMap)
	}
}

func generateMassiveMap(from, to int) map[string]interface{} {
	m := make(map[string]interface{}, from-to)
	for i := from; i < to; i++ {
		key := "key_" + fmt.Sprint(i)
		m[key] = map[string]interface{}{
			"nested": map[string]interface{}{
				"value": i,
				"bool":  i%2 == 1,
				"array": []interface{}{
					map[string]interface{}{
						"id":   i,
						"name": fmt.Sprintf("item_%d", i),
					},
					map[string]interface{}{
						"id":   i,
						"name": fmt.Sprintf("item_%d", i),
					},
					map[string]interface{}{
						"id":   i,
						"name": fmt.Sprintf("item_%d", i),
					},
				},
			},
		}
	}
	return m
}

func BenchmarkMergeMaps_Massive(b *testing.B) {
	newMap := generateMassiveMap(0, 1000)
	oldMap := generateMassiveMap(200, 1200)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mergeMaps(newMap, oldMap)
	}
}
