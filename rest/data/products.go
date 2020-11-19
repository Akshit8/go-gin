package data

import (
	"time"
)

// Product defines structure for an API endpoint
type Product struct {
	ID			int
	Name		string
	Description	string
	Price		float32
	SKU			string
	CreatedOn	string
	UpdatedOn	string
	DeletedOn	string
}

// diff style of declaring a slice
// works for all data type like arrays etc.
var productList = []*Product{
	&Product{
		
	}
}