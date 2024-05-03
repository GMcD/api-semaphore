package main

import (
	"strconv"

	"github.com/GMcD/api-semaphore/api"
)

func addProducts(a api.App, count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO products(name, price) VALUES($1, $2)", "Product "+strconv.Itoa(i), (i+1.0)*10)
	}
}
