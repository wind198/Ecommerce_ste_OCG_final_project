package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//sort out 3 category ID with highest number of product and put them into a categoryList
	categoryList := make(map[int]string)
	q := `select c.categoryID,c.categoryName from product as p
	left join category as c on p.categoryID=c.categoryID
	group by c.categoryID,c.categoryName
	order by count(productID) desc
	limit 3
	`
	db := newDbHanlder.db
	rows, err := db.Query(q)
	if err != nil {
		log.Println("Error during querying product table", err)
	}

	for rows.Next() {
		var categoryID int
		var categoryName string
		if err := rows.Scan(&categoryID, &categoryName); err != nil {
			log.Println("Error during Scanning categoryID:", err)
		}
		categoryList[categoryID] = categoryName
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		log.Println("Error during query and iteration,", err)
	}

	//for every ID in categoryList, query all product of this id and generate a ProductPerCategory struct for it
	productPerCategoryList := make([]ProductPerCategory, 0)
	for categoryID, categoryName := range categoryList {
		aProductPerCategory := ProductPerCategory{
			CategoryID:   categoryID,
			CategoryName: categoryName,
		}
		q = `select productID, productName,productPrice, productImageSource from product where categoryID=? limit 20`
		rows, err := db.Query(q, categoryID)
		if err != nil {
			log.Printf("Error during query for categoryID %v, %v\n", categoryID, err)
		}
		for rows.Next() {
			var aProduct Product
			if err := rows.Scan(&aProduct.ProductID, &aProduct.ProductName, &aProduct.ProductPrice, &aProduct.ProductImageSource); err != nil {
				log.Printf("Error during Scanning for product of category %v, %v \n", categoryID, err)
			} else {
				aProductPerCategory.ProductArray = append(aProductPerCategory.ProductArray, aProduct)
			}
		}
		rows.Close()
		productPerCategoryList = append(productPerCategoryList, aProductPerCategory)
		if err := rows.Err(); err != nil {
			log.Printf("Error during query and iteration of category ID %v, %v\n", categoryID, err)
		}

	}

	productPerCategoryListJSON, err := json.Marshal(productPerCategoryList)
	if err != nil {
		log.Printf("Error maping slice to json %v %v  ", categoryList, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Sorry but we have an internal sever problem"))
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(productPerCategoryListJSON)

}
