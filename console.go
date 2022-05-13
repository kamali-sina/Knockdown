package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	return name
}
