package shopping_cart

import (
	"e-commerce-synapsis/database"
	"log"
)

func AddCartDB(data AddCartRequest) (string, error) {
	db := database.PgOpenConnection()
	defer db.Close()

	query := `INSERT INTO carts (user_id, product_id,quantity) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, &data.UserID, &data.ProductID, &data.Quantity)

	if err!= nil {
		log.Println("[atom][shopping_cart][resource_db][AddCartDB] Error inserting cart : ", err)
        return "", err
    }
	
	return "Cart added successfully", nil
}

func GetCartDB(id int) ([]GetCartResponse, error) {
	db := database.PgOpenConnection()
	defer db.Close()

	var carts []GetCartResponse

	query := `SELECT carts.id, user_id, users.name, product_id, products.name, quantity FROM carts
	          LEFT JOIN users ON users.id = carts.user_id
			  LEFT JOIN products ON products.id = carts.product_id
	          WHERE user_id = $1 and is_active = 1`
	rows, err := db.Query(query, &id)

	if err!= nil {
        log.Println("[atom][shopping_cart][resource_db][GetCartDB] Error fetching carts : ", err)
        return nil, err
    }

	for rows.Next(){
		var cart GetCartResponse
        err := rows.Scan(&cart.ID, &cart.UserID, &cart.Username, &cart.ProductID, &cart.ProductName, &cart.Quantity)
        if err!= nil {
            log.Println("[atom][shopping_cart][resource_db][GetCartDB] Error scanning carts : ", err)
            return nil, err
        }
        carts = append(carts, cart)
	}

	return carts, nil
}

func DeleteCartByIdDB(data DeleteCartRequest) (string, error) {
	db := database.PgOpenConnection()
    defer db.Close()

    query := `UPDATE carts SET is_active = 0 WHERE id = $1 AND user_id = $2`
    _, err := db.Exec(query, &data.CartID, &data.UserID)

    if err!= nil {
        log.Println("[atom][shopping_cart][resource_db][DeleteCartByIdDB] Error deleting cart : ", err)
        return "", err
    }

    return "Cart deleted successfully", nil
}