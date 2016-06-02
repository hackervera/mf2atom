package main

import (
	"fmt"
	"os"

	"github.com/tjgillies/mf2atom"
)

func main() {
	atom := mf2atom.Parse(os.Args[1])
	fmt.Println(atom)
}
