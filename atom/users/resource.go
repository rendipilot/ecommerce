package users

func UserLoginUseCase(data UserLoginRequest)(string, error){
	res, err := UserLoginDB(data)

	return res, err
}

func UserRegisterUseCase(data UserRegisterRequest)(string, error){
	res, err := UserRegisterDB(data)

    return res, err
}