package orders

import (
	"e-commerce-synapsis/database"
	"log"
)

func CheckoutOrderDB(userId int) (int, error) {
	db := database.PgOpenConnection()
	defer db.Close()

	var carts []CartsItem
	
	query := `SELECT product_id, quantity FROM carts WHERE user_id = $1 AND is_active = 1`
	rows, err := db.Query(query, userId)

	if err != nil {
		log.Println("[atom][orders][resource_db][CheckoutOrder] error getting cart items : ", err)
		return 0,err
	}

	defer rows.Close()

	isEmpty := true
	var totalPrice float64
	for rows.Next() {
		isEmpty = false
		var cart CartsItem
        err := rows.Scan(&cart.ProductID, &cart.Quantity)
        if err != nil {
			log.Println("[atom][orders][resource_db][CheckoutOrder] error scanning carts : ", err)
            return 0,err
        }

		var price float64
		query := `SELECT price FROM products WHERE id = $1`
		err = db.QueryRow(query, cart.ProductID).Scan(&price)

		if err!= nil {
			log.Println("[atom][orders][resource_db][CheckoutOrder] error getting product price : ", err)
            return 0,err
        }

		totalPrice += price * float64(cart.Quantity)
		cart.Price = price
		carts = append(carts, cart)
	}

	if (isEmpty){
		return 0, nil
	}

	var orderId int
	query = `INSERT INTO orders (user_id, total_price) VALUES ($1, $2) RETURNING id`
	err = db.QueryRow(query, userId, totalPrice).Scan(&orderId)

	if err!= nil {
		log.Println("[atom][orders][resource_db][CheckoutOrder] error inserting order : ", err)
        return 0,err
    }

	for _, cart := range carts {
		query = `INSERT INTO order_items (order_id, product_id,quantity, price) values ($1, $2, $3, $4)`
		_, err = db.Exec(query, orderId, cart.ProductID, cart.Quantity, cart.Price)

		if err != nil {
			log.Println("[atom][orders][resource_db][CheckoutOrder] error inserting order items : ", err)
			return 0,err
		}
	}


	// clear carts data by id user
	query = `UPDATE carts SET is_active = 0 WHERE user_id = $1`
	_, err = db.Exec(query, userId)

	if err!= nil {
		log.Println("[atom][orders][resource_db][CheckoutOrder] error clearing cart data : ", err)
		return 0,err
	}

	return orderId, nil
}

func PaymentOrderDB(data PaymentRequest)(string, error){
	db := database.PgOpenConnection()
    defer db.Close()

    query := `UPDATE orders SET payment = $1, payment_status = 0 WHERE id = $2`
    _, err := db.Exec(query, &data.Method, &data.OrderID)

    if err!= nil {
        log.Println("[atom][orders][resource_db][PaymentOrder] error updating payment method : ", err)
        return "", err
    }

    return "success", nil
}