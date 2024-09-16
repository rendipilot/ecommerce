package shopping_cart

func AddCartUseCase(data AddCartRequest)(string, error){
	res, err := AddCartDB(data)

    return res, err
}

func GetCartUseCase(id int)([]GetCartResponse, error){
	res, err := GetCartDB(id)

    return res, err
}

func DeleteCartByIdUseCase(data DeleteCartRequest)(string, error){
	res, err := DeleteCartByIdDB(data)

    return res, err
}