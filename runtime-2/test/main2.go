package main

import (
	"fmt"
	_ "unsafe"
)

// byteIndex is strings.IndexByte. It returns the index of the
// first instance of c in s, or -1 if c is not present in s.
// strings.IndexByte is implemented in  runtime/asm_$GOARCH.s
//go:linkname byteIndex strings.IndexByte
func byteIndex(s string, c byte) int

func main() {
	fmt.Printf("found in %v\n", byteIndex("abcdefg", 'e'))
}
