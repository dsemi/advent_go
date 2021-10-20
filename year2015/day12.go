package year2015

import (
	"advent/problems"
	"encoding/json"
)

type Day12 struct{}

func (d Day12) walk(data interface{}, pred func(interface{}) bool) int {
	if pred(data) {
		return 0
	}
	var t int
	switch vv := data.(type) {
	case float64:
		t += int(vv)
	case map[string]interface{}:
		for _, v := range vv {
			t += d.walk(v, pred)
		}
	case []interface{}:
		for _, v := range vv {
			t += d.walk(v, pred)
		}
	}
	return t
}

func (d Day12) Part1(input string) interface{} {
	var j interface{}
	json.Unmarshal([]byte(input), &j)
	return d.walk(j, func(_ interface{}) bool { return false })
}

func (d Day12) Part2(input string) interface{} {
	var j interface{}
	json.Unmarshal([]byte(input), &j)
	return d.walk(j, func(j interface{}) bool {
		switch vv := j.(type) {
		case map[string]interface{}:
			for _, v := range vv {
				if v == "red" {
					return true
				}
			}
		}
		return false
	})
}

func init() {
	problems.Register(Day12{})
}
