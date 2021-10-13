package year2015

import (
	"advent/types"
	"encoding/json"
)

type Day12 struct{}

func walk(d interface{}, pred func(interface{}) bool) int {
	if pred(d) {
		return 0
	}
	var t int
	switch vv := d.(type) {
	case float64:
		t += int(vv)
	case map[string]interface{}:
		for _, v := range vv {
			t += walk(v, pred)
		}
	case []interface{}:
		for _, v := range vv {
			t += walk(v, pred)
		}
	}
	return t
}

func (Day12) Part1(input string) interface{} {
	var j interface{}
	json.Unmarshal([]byte(input), &j)
	return walk(j, func(_ interface{}) bool { return false })
}

func (Day12) Part2(input string) interface{} {
	var j interface{}
	json.Unmarshal([]byte(input), &j)
	return walk(j, func(j interface{}) bool {
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
	types.Register(Probs, Day12{})
}
