// hash_sha1.go
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	hash := sha1.New()
	io.WriteString(hash, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hash.Sum(b))
	fmt.Printf("Result: %d\n", hash.Sum(b))
	//
	hash.Reset()
	data := []byte("We shall overcome!")
	n, err := hash.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hash.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
