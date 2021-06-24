package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i uint8 = 127
	fmt.Println(i, i+1, i*i)
}

func ex2() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y) // XOR
	fmt.Printf("%08b\n", x&^y)

	for i := uint8(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x>>1)
	fmt.Printf("%08b\n", x<<1)
}

func ex3() {
	f := 1e100
	i := int(f)
	fmt.Println(f, i)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x\n", x)
}

func ex4() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}
