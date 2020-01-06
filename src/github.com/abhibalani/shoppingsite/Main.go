package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const customerCSV = "/Users/abhishek.balani/goworkspace/src/github.com/abhibalani/shoppingsite/csv_file/Customers.csv"
const employeeCSV = "/Users/abhishek.balani/goworkspace/src/github.com/abhibalani/shoppingsite/csv_file/Employees.csv"
const orderDetailsCSV = "/Users/abhishek.balani/goworkspace/src/github.com/abhibalani/shoppingsite/csv_file/OrderDetails.csv"
const ordersCSV = "/Users/abhishek.balani/goworkspace/src/github.com/abhibalani/shoppingsite/csv_file/Orders.csv"
const productsCSV = "/Users/abhishek.balani/goworkspace/src/github.com/abhibalani/shoppingsite/csv_file/Products.csv"
const shippingMethodCSV = "/Users/abhishek.balani/goworkspace/src/github.com/abhibalani/shoppingsite/csv_file/ShippingMethods.csv"

func main() {

	// LoadCSVtoDB()
	// GetCustomers("city", "Irvine city")
	// getCustomersByEmployee("Adam", "Bar")
	HandleRequests()

}

// HandleRequests ...
func HandleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/customers", GetAllCustomer)
	http.HandleFunc("/orders/{orderId}", GetOrderDetails)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

// HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// GetOrderDetails ...
func GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderId"]
	fmt.Println(orderID)

	db := DbConn()
	query := "SELECT * FROM order_detail WHERE OrderId=?"
	selDB, err := db.Query(query, orderID)
	if err != nil {
		panic(err.Error())
	}

	orders := SerializedOrderData(selDB)
	json.NewEncoder(w).Encode(orders)
}

// GetAllCustomer ...
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	query := "SELECT * FROM customer"
	selDB, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	customers := SerializeCustomerData(selDB)
	fmt.Println(customers)
	json.NewEncoder(w).Encode(customers)
}

// LoadCSVtoDB ...
func LoadCSVtoDB() {

	db := DbConn()
	mysql.RegisterLocalFile(customerCSV)
	mysql.RegisterLocalFile(employeeCSV)
	mysql.RegisterLocalFile(orderDetailsCSV)
	mysql.RegisterLocalFile(ordersCSV)
	mysql.RegisterLocalFile(productsCSV)
	mysql.RegisterLocalFile(shippingMethodCSV)

	_, err := db.Exec("LOAD DATA LOCAL INFILE '" + customerCSV + "' INTO TABLE customer")
	if err != nil {
		fmt.Println("Unable to load data", err)
	}
	_, err1 := db.Exec("LOAD DATA LOCAL INFILE '" + employeeCSV + "' INTO TABLE employee")
	if err1 != nil {
		fmt.Println("Unable to load data", err1)
	}
	_, err2 := db.Exec("LOAD DATA LOCAL INFILE '" + orderDetailsCSV + "' INTO TABLE order_detail")
	if err2 != nil {
		fmt.Println("Unable to load data", err2)
	}
	_, err3 := db.Exec("LOAD DATA LOCAL INFILE '" + ordersCSV + "' INTO TABLE orders")
	if err3 != nil {
		fmt.Println("Unable to load data", err3)
	}
	_, err4 := db.Exec("LOAD DATA LOCAL INFILE '" + productsCSV + "' INTO TABLE product")
	if err4 != nil {
		fmt.Println("Unable to load data", err4)
	}
	_, err5 := db.Exec("LOAD DATA LOCAL INFILE '" + shippingMethodCSV + "' INTO TABLE shipping_method")
	if err5 != nil {
		fmt.Println("Unable to load data", err5)
	}

}

// DbConn ...
func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "admin123"
	dbName := "gotest"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to DB")
	}
	// defer db.Close()

	return db
}

// DbCheck ...
func DbCheck(db *sql.DB) {

	err := db.Ping()
	if err != nil {
		fmt.Println("Cannot ping db.")
	} else {
		fmt.Println("Pinged DB.")
	}

}

// GetCustomers ...
func GetCustomers(attr, value string) *sql.Rows {
	db := DbConn()
	query := "SELECT * FROM customer WHERE " + attr + "=?"
	selDB, err := db.Query(query, value)
	if err != nil {
		panic(err.Error())
	}
	return selDB
}

// GetCustomersByEmployee ...
func GetCustomersByEmployee(empFirstName, empLastName string) *sql.Rows {
	db := DbConn()
	selDB, err := db.Query("SELECT * FROM customer INNER JOIN employee WHERE employee.firstName=?", empFirstName, "AND employee.lastName=?", empLastName)
	fmt.Println(selDB)
	if err != nil {
		panic(err.Error())
	}
	return selDB
}

//SerializeCustomerData ...
func SerializeCustomerData(selDB *sql.Rows) []Customer {
	var result []Customer = make([]Customer, 100)
	for selDB.Next() {
		cust := Customer{}
		var id int
		var companyName, firstname, lastname, billAdd, city, state, zip, email, companySite, phoneNum, faxNum, shipAdd, shipCity, shipState, shipZip, shipPhone string
		err := selDB.Scan(&id, &companyName, &firstname, &lastname, &billAdd, &city, &state, &zip, &email, &companySite, &phoneNum, &faxNum, &shipAdd, &shipCity, &shipState, &shipZip, &shipPhone)
		if err != nil {
			panic(err.Error())
		}
		cust.CustomerID = id
		cust.CompanyName = companyName
		cust.FirstName = firstname
		cust.LastName = lastname
		cust.BillingAddress = billAdd
		cust.City = city
		cust.StateOrProvince = state
		cust.ZipCode = zip
		cust.Email = email
		cust.CompanyWebsite = companySite
		cust.PhoneNumber = phoneNum
		cust.FaxNumber = faxNum
		cust.ShipAddress = shipAdd
		cust.ShipCity = shipCity
		cust.ShipStateProvince = shipState
		cust.ShipZipCode = shipZip
		cust.ShipPhoneNumber = shipPhone

		result = append(result, cust)
	}
	return result
}

// SerializedOrderData ...
func SerializedOrderData(selDB *sql.Rows) []Order {
	var result []Order = make([]Order, 100)
	for selDB.Next() {
		order := Order{}
		var orderID, customerID, employeeID, shippingMethodID, freightCharge, taxes int
		var orderDate, shipDate, paymentReceived, comment string
		err := selDB.Scan(&orderID, &customerID, &employeeID, &shippingMethodID, &orderDate, &shipDate, &freightCharge, &taxes, &paymentReceived, &comment)
		if err != nil {
			panic(err.Error())
		}
		order.OrderID = orderID
		order.customerID = customerID
		order.employeeID = employeeID
		order.shippingMethodID = shippingMethodID
		order.orderDate = orderDate
		order.shipDate = shipDate
		order.freightCharge = freightCharge
		order.taxes = taxes
		order.paymentReceived = paymentReceived
		order.comment = comment

		result = append(result, order)
	}
	return result
}

// Customer model
type Customer struct {
	CustomerID        int
	CompanyName       string
	FirstName         string
	LastName          string
	BillingAddress    string
	City              string
	StateOrProvince   string
	ZipCode           string
	Email             string
	CompanyWebsite    string
	PhoneNumber       string
	FaxNumber         string
	ShipAddress       string
	ShipCity          string
	ShipStateProvince string
	ShipZipCode       string
	ShipPhoneNumber   string
}

//Employee Model
type Employee struct {
	employeeID int
	firstName  string
	lastName   string
	title      string
	workPhone  string
}

//Product Model
type Product struct {
	productID   string
	productName string
	unitPrice   float32
	inStock     string
}

// ShippingMethod model
type ShippingMethod struct {
	shippingMethodID int
	shippingMethod   string
}

// Order Model
type Order struct {
	OrderID          int
	customerID       int
	employeeID       int
	shippingMethodID int
	orderDate        string
	shipDate         string
	freightCharge    int
	taxes            int
	paymentReceived  string
	comment          string
}

// OrderDetail model
type OrderDetail struct {
	orderDetailID int
	orderID       int
	productID     int
	quantity      int
	unitPrice     int
	discount      int
}
