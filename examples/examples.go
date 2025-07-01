// Package main provides examples for the safe conversion package.
package main

import (
	"log"

	safeconversion "github.com/bsv-blockchain/go-safe-conversion"
)

// main function demonstrates the usage of safe conversion functions.
func main() {

	// Example of converting int32 to uint32 safely
	val, err := safeconversion.Int32ToUint32(42)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Converted value: %d", val)
}
