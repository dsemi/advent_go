package main

import (
	"strings"
	"utils"
)

type food struct {
	ingredients, allergens []string
}

func parse(input string) []food {
	foods := make([]food, 0)
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Split(line, " (contains ")
		foods = append(foods, food{
			ingredients: strings.Fields(pts[0]),
			allergens:   strings.Split(pts[1][:len(pts[1])-1], ", "),
		})
	}
	return foods
}

func allergens(foods []food) map[string]map[string]bool {
	m := make(map[string]map[string]bool)
	for _, f := range foods {
		for _, allergen := range f.allergens {
			ings := make(map[string]bool)
			for _, ing := range f.ingredients {
				ings[ing] = true
			}
			if _, ok := m[allergen]; !ok {
				m[allergen] = ings
			}
			for k := range m[allergen] {
				if _, ok := ings[k]; !ok {
					delete(m[allergen], k)
				}
			}
		}
	}
	return m
}

func Part1(input string) interface{} {
	foods := parse(input)
	alls := allergens(foods)
	safe := make(map[string]bool)
	for _, f := range foods {
		for _, i := range f.ingredients {
			safe[i] = true
		}
	}
	for _, v := range alls {
		for x := range v {
			delete(safe, x)
		}
	}
	var cnt int
	for _, f := range foods {
		for _, i := range f.ingredients {
			if safe[i] {
				cnt++
			}
		}
	}
	return cnt
}

func Part2(input string) interface{} {
	alls := allergens(parse(input))
	done := make(map[string]string)
	for len(alls) > 0 {
		for k, v := range alls {
			if len(v) == 1 {
				for x := range v {
					done[k] = x
				}
			}
		}
		s := make(map[string]bool)
		for _, v := range done {
			s[v] = true
		}
		alls2 := make(map[string]map[string]bool)
		for k, v := range alls {
			if _, ok := done[k]; ok {
				continue
			}
			s2 := make(map[string]bool)
			for j := range v {
				if !s[j] {
					s2[j] = true
				}
			}
			alls2[k] = s2
		}
		alls = alls2
	}
	keys := make([]string, 0)
	for k := range done {
		keys = append(keys, k)
	}
	utils.Sort(keys)
	vals := make([]string, len(keys))
	for i, k := range keys {
		vals[i] = done[k]
	}
	return strings.Join(vals, ",")
}
