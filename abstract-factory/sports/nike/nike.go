package nike

import "abstract-factory/sports"

type Nike struct {
}

func (n *Nike) MakeShoe() sports.IShoe {
	return &nikeShoe{
		Shoe: sports.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() sports.IShirt {
	return &nikeShirt{
		Shirt: sports.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
