package models

type Product struct {
	Id          int
	Description string
	Quantify    int
	Price       float64
	CreatedAt   string
}

func GetAll() ([]Product, error) {
	con := Connect()
	sql := "select * from product"
	rs, erro := con.Query(sql)

	if erro != nil {
		return nil, erro
	}

	var products []Product

	for rs.Next() {
		var product Product
		erro := rs.Scan(
			&product.Id,
			&product.Description,
			&product.Quantify,
			&product.Price,
			&product.CreatedAt)

		if erro != nil {
			return nil, erro
		}
		products = append(products, product)
	}

	defer con.Close()
	defer con.Close()

	return products, nil
}
