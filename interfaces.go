package main

import "fmt"

type I1 interface {
	M1()
}
type I2 interface {
	M2()
}
type T struct{}
func (T) M1() { fmt.Println("T.M1") }
func (T) M2() { fmt.Println("T.M2") }
func f1(i I1) { i.M1() }
func f2(i I2) { i.M2() }
func main() {
	t := T{}
	f1(t) // "T.M1"
	f2(t) // "T.M2"
}
