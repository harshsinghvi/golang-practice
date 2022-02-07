package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan string, 10)
	inputs := make(chan int, 10)
	go complexCalc(inputs, results)
	// go complexCalc(inputs)
	// go complexCalc(inputs)

	

	for i:=0; i<20; i++ {
		inputs <- i
	}
	go func() {
		for result:= range results{
			fmt.Println(result)
		}
	}()
	close(inputs)
	fmt.Println("done")
	
	fmt.Scanln()
}

func complexCalc(inp <- chan int, out chan<- string) {
	for input:= range inp{
		for i:=0; i<2; i++{
			fmt.Printf("%d-%d ",input, i)
			time.Sleep((time.Millisecond * 50))
		}
		out <- fmt.Sprintf("Calculated %d\n",input)
	}
		
	// out <- inp
}
