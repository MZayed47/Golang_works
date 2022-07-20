package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const a = 100
	const d = 3e20 / n
	const b = 3e20 / a

	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(b)
	fmt.Println(int64(b))

	fmt.Println(math.Cos(n))
}
