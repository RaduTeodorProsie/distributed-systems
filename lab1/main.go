package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ex1() error {
	stdin := bufio.NewScanner(os.Stdin)

	stdin.Scan()
	num, err := strconv.Atoi(stdin.Text())
	if err != nil{
		return err
	}
	
	prime := num >= 2 
	for i := 2; i * i <= num; i++{
		if num % i == 0{
			prime = false 
			break
		} 
	}

	if prime{
		fmt.Println("The number is prime!")
	} else{
		fmt.Println("The number is not prime!")
	}

	return nil
}

func ex2() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	
	// vector := []int32{}
}

func main(){
	fmt.Printf("what is the exercise number?\n")
	reader := bufio.NewScanner(os.Stdin)
	
	reader.Scan()
	input := reader.Text()
	strings.TrimSpace(input)

	switch {
	case input == "1":
		ex1()
	}
}
