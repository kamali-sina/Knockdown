package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	p := MakePlayer(name)
	fmt.Println(p.getInfo())
}
