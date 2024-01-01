package main

import (
	"fmt"
	"strings"

	s "github.com/patricho/advent-of-code/go/shared"
)

func main() {
	solve(1, "input.txt")
}

var (
	workflows                 map[string]Workflow
	parts, accepted, rejected []map[string]int
)

func solve(part int, filename string) {
	accepted = []map[string]int{}
	rejected = []map[string]int{}

	parts := parseInput(filename)

	for _, p := range parts {
		applyWorkflow("in", p)
	}

	sum := 0
	for _, acc := range accepted {
		for k := range acc {
			sum += acc[k]
		}
	}

	fmt.Println("accepted", len(accepted), "sum", sum)
}

func parseInput(filename string) []map[string]int {
	lines := s.ReadFile(filename)
	parts = []map[string]int{}
	workflows = map[string]Workflow{}
	parsingWorkflows := true
	for _, line := range lines {
		if line == "" {
			parsingWorkflows = false
			continue
		}
		if parsingWorkflows {
			wfstr := strings.Split(strings.Replace(line, "}", "", -1), "{")
			rulesstringarr := strings.Split(wfstr[1], ",")
			rules := []Rule{}
			for _, rlstr := range rulesstringarr {
				rulearr := strings.Split(rlstr, ":")
				if len(rulearr) > 1 {
					if strings.Contains(rulearr[0], ">") {
						ruleargs := strings.Split(rulearr[0], ">")

						rules = append(rules, Rule{
							Key:      ruleargs[0],
							Comparer: "gt",
							Value:    s.ToInt(ruleargs[1]),
							Target:   rulearr[1],
						})
					} else {
						ruleargs := strings.Split(rulearr[0], "<")

						rules = append(rules, Rule{
							Key:      ruleargs[0],
							Comparer: "lt",
							Value:    s.ToInt(ruleargs[1]),
							Target:   rulearr[1],
						})
					}
				} else {
					rules = append(rules, Rule{
						Key:      "",
						Comparer: "move",
						Target:   rulearr[0],
					})
				}
			}
			workflows[wfstr[0]] = Workflow{Rules: rules, Parts: []map[string]int{}}
		} else {
			ptstrs := strings.Split(strings.Replace(strings.Replace(line, "{", "", -1), "}", "", -1), ",")
			pt := map[string]int{}
			for _, ptstr := range ptstrs {
				ptarr := strings.Split(ptstr, "=")
				pt[ptarr[0]] = s.ToInt(ptarr[1])
			}
			parts = append(parts, pt)
		}
	}
	return parts
}

func applyWorkflow(wfkey string, part map[string]int) {
	if wfkey == "A" {
		accepted = append(accepted, part)
		return
	} else if wfkey == "R" {
		rejected = append(rejected, part)
		return
	}

	wf := workflows[wfkey]

	for _, rule := range wf.Rules {
		if ruleMatches(rule, part) {
			applyWorkflow(rule.Target, part)
			break
		}
	}
}

func ruleMatches(r Rule, part map[string]int) bool {
	switch r.Comparer {
	case "move":
		return true
	case "gt":
		return part[r.Key] > r.Value
	case "lt":
		return part[r.Key] < r.Value
	}
	panic("wtf")
}

type Workflow struct {
	Rules []Rule
	Parts []map[string]int
}

type Rule struct {
	Key      string
	Comparer Comparer
	Value    int
	Target   string
}

type Comparer string
