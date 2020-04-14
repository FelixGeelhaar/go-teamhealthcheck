package main

import (
	"fmt"
	"teamhealthcheck/metrics"
)

func main() {
	m := metrics.New("Health of Codebase")
	m.SetTendency(-4)

	fmt.Printf("%#v\n", m)

	m.SetGreen(5)
	m.SetYellow(3)
	m.SetRed(2)
	fmt.Printf("%#v\n", m)
}
