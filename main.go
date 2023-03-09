package main

import "fmt"

func main() {
	
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	
	for j := 0; j < 5; j++ {
		fmt.Println("Nilai j = ", j)
	}


	var unicode = "САШАРВО"
    	for i, v := range unicode {
        	fmt.Printf("character %#U starts at byte position %d\n", v, i)
    	}

	
	for j := 6; j <= 10; j++ {
		fmt.Println("Nilai j = ", j)
	}


}
