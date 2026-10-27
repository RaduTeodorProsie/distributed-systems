package main 

import(
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func main() {

	stdin := bufio.NewScanner(os.Stdin)

	stdin.Scan()
	num, _ := strconv.Atoi(stdin.Text())
	
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
}
