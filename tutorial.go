package main

import (
	"tutorial/basic"
	"tutorial/concurrency"
	"tutorial/generics"
	interfacego "tutorial/interface"
	"tutorial/methods"
)

func main() {
	basic.RunBasic()
	methods.RunMethods()
	interfacego.RunInterface()
	generics.RunGenerics()
	concurrency.RunConcurrency()
}
