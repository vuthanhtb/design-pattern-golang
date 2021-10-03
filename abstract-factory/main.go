package main

import (
	"abstract-factory/sports"
	"abstract-factory/sports/adidas"
	"abstract-factory/sports/nike"
	"fmt"
)

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	fmt.Println("----------------------------------------")
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func getSportsFactory(brand string) (sports.ISportsFactory, error) {
	if brand == "adidas" {
		return &adidas.Adidas{}, nil
	}

	if brand == "nike" {
		return &nike.Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

func printShoeDetails(s sports.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s sports.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
