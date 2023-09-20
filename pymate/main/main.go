package main

import "pymate"

func createGroup(numPrimates int) Group {
	g := Group{numPrimates: numPrimates}
	g.CreatePrimates(numPrimates)
	return g
}

func main() {
	createGroup(10)
}