package year2015

import (
	"advent/types"
	"advent/utils"
	"regexp"
	"strings"
)

type Day19 struct{}

type Rep struct {
	src, dest string
}

func parseMappings(input string) (string, []Rep) {
	v := strings.Split(input, "\n\n")
	var ms []Rep
	for _, line := range strings.Split(v[0], "\n") {
		pts := strings.Split(line, " => ")
		ms = append(ms, Rep{src: pts[0], dest: pts[1]})
	}
	return v[1], ms
}

func singleReplacements(src, k, v string) []string {
	reg := regexp.MustCompile(k)
	var reps []string
	for _, is := range reg.FindAllStringIndex(src, -1) {
		reps = append(reps, src[:is[0]]+v+src[is[1]:])
	}
	return reps
}

func (Day19) Part1(input string) interface{} {
	s, reps := parseMappings(input)
	m := make(map[string]bool)
	for _, rep := range reps {
		for _, r := range singleReplacements(s, rep.src, rep.dest) {
			m[r] = true
		}
	}
	return len(m)
}

func (Day19) Part2(input string) interface{} {
	s, reps := parseMappings(input)
	mol := utils.Reverse(s)
	mrep := make(map[string]string)
	for _, rep := range reps {
		mrep[utils.Reverse(rep.dest)] = utils.Reverse(rep.src)
	}
	var re string
	for r := range mrep {
		re += "|"
		re += r
	}
	reg := regexp.MustCompile(re[1:])
	var i int
	for i = 0; mol != "e"; i++ {
		repl := false
		mol = reg.ReplaceAllStringFunc(mol, func(k string) string {
			if repl {
				return k
			}
			repl = true
			return mrep[k]
		})
	}
	return i
}

func init() {
	types.Register(Probs, Day19{})
}
