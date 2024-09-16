package orders

func CheckoutOrderUseCase(id int)(int, error){
	res, err := CheckoutOrderDB(id)

    return res, err
}