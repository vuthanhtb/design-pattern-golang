package adidas

import "abstract-factory/sports"

type Adidas struct {}

func (a *Adidas) MakeShoe() sports.IShoe {
	return &adidasShoe{
		Shoe: sports.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() sports.IShirt {
	return &adidasShirt{
		Shirt: sports.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
