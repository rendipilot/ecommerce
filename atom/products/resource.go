package products

func GetProductByCategoryUseCase(category string)([]Products, error){
	res, err := GetProductByCategoryDB(category)

    return res, err
}