package main

import "fmt"
// import "rsc.io/quote"

func add(x int, y int) int { return x + y }

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// Doesnot have functional overloading
// func typeof(a int) string { return "INT"}
// func typeof(a float32) string { return "FLOAT32"}
// func typeof(a float64) string { return "FLOAT64"}


func main() {

	// var x int;
	// var s = "TEST"
	// a,b,c := 1,2,3
	// x = 10

	// fmt.Printf("Hello World!! %d\n", 12)
	// fmt.Printf("%d\n",add(1,2))
	// fmt.Println(x)
	// fmt.Println(s)
	// fmt.Println(add(1,2))
	// fmt.Println(a,b,c)

	
	//fmt.Println(Hello("HARSH"))
	//fmt.Println(quote.Go())
	
	const A = "SA"
	fmt.Printf("%T %d %s\n", A, A, A)
	if fmt.Sprintf("%T", A) == "int" {
		fmt.Println("type comarision")
	}

}
