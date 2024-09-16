package users

import (
	"e-commerce-synapsis/database"
	utils_token "e-commerce-synapsis/utils"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func UserLoginDB(data UserLoginRequest) (string, error) {
	db := database.PgOpenConnection()
	defer db.Close()

	var response User
	var token string

	query :=`SELECT id, name, email, password FROM users WHERE email =$1`

	row := db.QueryRow(query, &data.Email)

	err := row.Scan(&response.ID, &response.Name, &response.Email, &response.Password)

	if err != nil {
		log.Println("[atom][users][UserRegisterDB] error get user data : ", err)
		return "", err
	}

	// check password

	err = bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(data.Password))

	if err != nil {
		log.Println("[atom][users][UserRegisterDB] error compare password : ", err)
		return "", err
	}

	token, err = utils_token.GenerateToken(response.ID)
	if err!= nil {
		log.Println("[atom][users][UserRegisterDB] error creating generated token : ", err)
        return "", err
    }

	return token, nil
}

func UserRegisterDB(data UserRegisterRequest) (string, error){
	db := database.PgOpenConnection()
	defer db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	
	var email string

	checkQuery := `SELECT email FROM users WHERE email =$1`

	row := db.QueryRow(checkQuery, &data.Email)
	err = row.Scan(&email)

	if err == nil {
        return "", errors.New("email is already registered")
    }

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	_, err = db.Exec(query, data.Name, data.Email, hashedPassword)

	if err != nil {
		log.Println("[atom][users][UserRegisterDB] error = creating new user")
		return "", errors.New("registration failed")
	}

	return data.Name, nil
}