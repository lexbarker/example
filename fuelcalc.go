package main

import "fmt"

//import "fmt"

func main() {
	miles := 300
	var litres float64
	//var litre float64
	const l2gal = 4.53
	fmt.Scan(&litres)
	mpl := float64(miles) / float64(litres)
	fmt.Printf("you achieved %f mile per litre\n", mpl)
	mpg := mpl * l2gal
	fmt.Printf("mpg  = %f\n", mpg)
}
