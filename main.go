package main

import (
	"fmt"
	"time"

	"github.com/beckxie/autogen-twid/generate"
)

func main() {
	fmt.Printf("%v ID is generated:%s\n", time.Now().Format(time.RFC3339), generate.GenerateID())
}
