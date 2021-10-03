package sports

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}
