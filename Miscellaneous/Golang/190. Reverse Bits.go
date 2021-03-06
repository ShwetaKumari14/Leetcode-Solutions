package main

import "fmt"

func main() {
	num := 00000000000000000000000000001011
	var n uint32 = uint32(num)
	result := reverseBits(n)
	fmt.Println(result)
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num >>= 1
	}
	return res
}
