package main

import (
	"bufio"
	gormdb "cinema/infra/gormdb"
	"fmt"
	"os"
	"strings"
)

func StringPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}

func main() {
	fmt.Println("Hello world")

	name :=  StringPrompt("Qual é seu nome");

	db,err := gormdb.SetupGormDB()

	if err != nil {
		panic(err)
	}

	fmt.Println(name)
}