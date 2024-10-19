package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/medragneel/lex/repl"
)

func main() {
	// 	input := `let five = 5;
	// let ten = 10;
	// let add = fn(x, y) {
	// x + y;
	// };
	// let result = add(five, ten);
	// `
	current, err := user.Current()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("Welcome %v  to The monkey REPL\n", current.Username)
	fmt.Println("feel free to type monkey commands")
	repl.Start(os.Stdin, os.Stdout)
}
