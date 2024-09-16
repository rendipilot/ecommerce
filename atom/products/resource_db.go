package products

import (
	"e-commerce-synapsis/database"
	"log"
)

func GetProductByCategoryDB(category string) ([]Products, error) {
	db := database.PgOpenConnection()
	defer db.Close()

	var products []Products

	query := `SELECT id, name, description, price, category, stock FROM products WHERE category = $1`

	rows, err := db.Query(query, &category)

	if err != nil {
		log.Println("[atom][products][resource_db][GetProductByCategoryDB] error failed to query products : ", err)
		return nil, err
	}

	for rows.Next() {
		var product Products
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Stock)

		if err != nil {
			log.Println("[atom][products][resource_db][GetProductByCategoryDB] error failed to scanning products : ", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, err
}