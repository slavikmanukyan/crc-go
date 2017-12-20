package main

import (
	"fmt"
)

const (
	// WIDTH CRC32 Width
	WIDTH = uint(32)
	// POLYNOMIAL CRC32 polynomial
	POLYNOMIAL = uint64(0x04C11DB7)
	// INIT CRC32 init value
	INIT = uint64(0xFFFFFFFF)
	// FINALXOR CRC32 final xor value
	FINALXOR = uint64(0xFFFFFFFF)
)

// reflect reverses order of last count bits
func reflect(in uint64, count uint) uint64 {
	ret := in
	for idx := uint(0); idx < count; idx++ {
		srcbit := uint64(1) << idx
		dstbit := uint64(1) << (count - idx - 1)
		if (in & srcbit) != 0 {
			ret |= dstbit
		} else {
			ret = ret & (^dstbit)
		}
	}
	return ret
}

func CRC(data []byte) uint64 {

	curValue := INIT
	topbit := uint64(1) << (WIDTH - 1)
	mask := (topbit << 1) - 1

	for i := 0; i < len(data); i++ {
		var curByte = uint64(data[i]) & 0x00FF
		curByte = reflect(curByte, 8)

		curValue ^= (curByte << (WIDTH - 8))
		for j := 0; j < 8; j++ {
			if (curValue & topbit) != 0 {
				curValue = (curValue << 1) ^ POLYNOMIAL
			} else {
				curValue = (curValue << 1)
			}
		}

	}
	curValue = reflect(curValue, WIDTH)

	curValue = curValue ^ FINALXOR

	return curValue & mask
}

func main() {
	fmt.Println(CRC([]byte("long long text")))
}
