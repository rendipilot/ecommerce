package orders

func CheckoutOrderUseCase(id int)(int, error){
	res, err := CheckoutOrderDB(id)

    return res, err
}

func PaymentOrderUseCase(data PaymentRequest)(string, error){
	res, err := PaymentOrderDB(data)

    return res, err
}