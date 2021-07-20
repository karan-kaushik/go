package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("User Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)
}