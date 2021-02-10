package mod

import (
	"fmt"
	"os"
)

func SayHello() {
	fmt.Println("hello")
	fmt.Fprintf(os.Stderr, "opening the garden door")
}

