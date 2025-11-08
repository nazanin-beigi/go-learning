package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var f float64 = 12.8
	var i int = int(f)
	fmt.Println(i)

	//type inference
	var w int
	j := w
	fmt.Println(j)

	//for

	sum := 0
	for i := 0; i < 10; i++ {
		// sum+=i  means  sum=sum+i
		sum += i
	}
	fmt.Println(sum)

	//declaring short variable is only allowed in fun main & be outside of for loop.
	//this loop never ends
	//semicolons can be removed
	q := 0
	for q < 100 {
		q -= 1 // اینجا تغییر بعد از شرط نیست ;باید داخل فور باشه
	}

	//////////////////////
	z := 0
	for z < 100 {
		z = z + 1
	}
	fmt.Println(z)

}

var n, c, e bool
var t, y = true, false

const New = 2.65
