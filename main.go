package main

import (
	"fmt"
	"go-basics/calculator"
	"os"
	"unsafe"
)

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)



func main(){
	fmt.Println(os.Getenv("GO_DEV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(20,40))
	fmt.Println(calculator.Multiply(1,2))


	//値と型を表示
	i := 2
	ui := uint16(2)
	f := 1.23333
	s := "hello"
	b := true
	pi, title := 3.14, "GO"
	fmt.Println(i)
	fmt.Printf("i: %v %T\n",i, i)
	fmt.Printf("f: %[1]v %[1]T\n",f)
	fmt.Printf("s: %[1]v %[1]T\n",s)
	fmt.Printf("b: %[1]v %[1]T\n",b)
	//indexで指定
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i ,ui)
	fmt.Printf("pi: %v title: %v\n",pi,title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println(z)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n",Mac,Windows,Linux)

	var ui1 uint16
	//&をつけることで先頭の1byteを取得
	fmt.Printf("memory address of uil: %p\n", &ui1)
	var ui2 uint16
	fmt.Printf("memory address of uil: %p\n", &ui2)
	//ポインタ変数
	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	//ui1をポインタ変数p1に代入
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1) //一致する
	//ポインタ変数の長さと先頭アドレスを出力
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("value of uil(dereference): %v\n", *p1)
	//ui1の値を書き換える
	*p1 = 1
	fmt.Printf("value of uil: %v\n", ui1)
	//ダブルポインタ
	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n",ui1)

	//shadowing
	ok, result := true, "A"
	if ok {
		//このresultはスコープ内限定
		result := "B"
		println(result)
	} else {
		result :="C"
		println(result)
	}
	println(result) //A
}