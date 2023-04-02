package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"stocksapp/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to postgres...")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Printf("Unable to decode request body. %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID:      insertID,
		Message: "stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert string into int. %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Printf("Unable to get stock. %v", err)
		http.Error(w, "stock not found", http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(stock)
	}

}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()

	if err != nil {
		log.Printf("Unable to get all stocks %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert string into int. %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Printf("Unable to decode request body. %v", err)
	}

	updatedRows := updateStock(int64(id), stock)

	res := response{
		ID:      stock.StockID,
		Message: fmt.Sprintf("stock updated. Total rows affected %v", updatedRows),
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert string into int. %v", err)
	}

	deletedRows := deleteStock(int64(id))

	res := response{
		ID:      int64(id),
		Message: fmt.Sprintf("stock deleted. Total rows affected %v", deletedRows),
	}

	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`

	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Printf("Unable to execute query. %v", err)
	}

	log.Printf("inserted single recorde %d", id)

	return id
}

func getStock(id int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement := `SELECT * FROM stocks where stockid = $1`

	err := db.QueryRow(sqlStatement, id).Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		log.Println("No rows returned.")
		return stock, err
	case nil:
		return stock, nil
	default:
		log.Printf("Unable to execute query. %v", err)
	}

	return stock, err
}

func getAllStocks() ([]models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stocks []models.Stock

	sqlStatement := `SELECT * FROM stocks`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Printf("unable to execute query. %v", err)
		return nil, err
	}

	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Printf("unable to scan row %v", err)
		}

		stocks = append(stocks, stock)
	}

	return stocks, err
}

func updateStock(id int64, stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Printf("unable to execute query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("error. %v", err)
		return 0
	}

	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Printf("unable to execute query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("error. %v", err)
		return 0
	}

	return rowsAffected
}
