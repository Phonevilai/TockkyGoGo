package main

import (
	"fmt"
	"github.com/Phonevilai/TockkyGoGo/tockprint"
)

func init() {
	tockprint.TockInit()
}

func main() {
	name := tockprint.P()
	fmt.Println(name)
}
