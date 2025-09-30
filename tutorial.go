package main

import (
	"tutorial/basic"
	interfacego "tutorial/interface"
	"tutorial/methods"
)

func main() {
	basic.RunBasic()
	methods.RunMethods()
	interfacego.InterfaceGo()
	interfacego.IntercaceImplicitDemo()
	interfacego.DependencyInversionDemo()
	interfacego.InterfaceValue()
	interfacego.EmptyInterface()
}
