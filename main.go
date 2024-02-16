package main

import (
	"fmt"

	dfaanalyzer "github.com/0721santi/DFAAnalyzer"
)

func main() {
	dfa := dfaanalyzer.GenDFA()
	dfa.AddSymbol("0")
	dfa.AddSymbol("1")
	q0, _ := dfa.AddState("0", true)
	q1, _ := dfa.AddState("1", false)
	q2, _ := dfa.AddState("2", false)
	dfa.SetStartState("0")
	dfa.AddTransitionFunc(q0, q0, []string{"0"})
	dfa.AddTransitionFunc(q0, q1, []string{"1"})
	dfa.AddTransitionFunc(q1, q0, []string{"1"})
	dfa.AddTransitionFunc(q1, q2, []string{"0"})
	dfa.AddTransitionFunc(q2, q1, []string{"0"})
	dfa.AddTransitionFunc(q2, q2, []string{"1"})
	var str string
	fmt.Print("> ")
	fmt.Scanln(&str)
	if str == "E" {
		fmt.Printf("The DFA accepts the string %s\n", str)
	} else {
		accepted := dfa.AnalizeChain(str)
		if accepted {
			fmt.Printf("The DFA accepts the string %s\n", str)
		} else {
			fmt.Printf("The DFA rejects the string %s\n", str)
		}
	}
}
