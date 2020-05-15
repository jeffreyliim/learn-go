// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

type TestStruct struct {
	TestInterface
	Msg string
}

type TestInterface interface {
	DoSomething()
}

func (t *TestStruct) DoSomething(){
	fmt.Print(t.Msg)
}

func main() {
	t := TestStruct{Msg:"asd"}
	t.DoSomething()
}
