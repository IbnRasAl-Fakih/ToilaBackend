package main

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

var bdFlaq bool
var restaurants []Restaurant
var bdTam bool
var tamada []Tamada

func main() {
	http.HandleFunc("/getRestaurants", getRestaurants)
	http.HandleFunc("/setRestaurant", setRestaurant)
	http.HandleFunc("/deleteRestaurant", deleteRestaurant)
	http.HandleFunc("/getTamada", getTamada)
	http.HandleFunc("/setTamada", setTamada)
	http.HandleFunc("/deleteTamada", deleteTamada)
	http.ListenAndServe(":8080", nil)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	if bdFlaq {
		json.NewEncoder(w).Encode(restaurants)
		return
	}
	restaurants = restaurants[:0]
	db := GetDatabaseInstance()
	//defer db.connection.Close()
	rows, err := db.connection.Query(`select * from restaurants`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var name string
		var city string
		var img string
		var price int
		var numOfGuests int
		var address string
		var description string
		var googleMap string
		var images pq.StringArray
		var phone string
		if err := rows.Scan(&name, &city, &img, &price, &numOfGuests, &address, &description, &googleMap, &phone, &images, &id); err != nil {
			panic(err)
		}
		res := Restaurant{id, name, city, img, strconv.Itoa(price), strconv.Itoa(numOfGuests), address, description, googleMap, []string(images), phone}
		restaurants = append(restaurants, res)
	}
	//defer rows.Close()
	json.NewEncoder(w).Encode(restaurants)
	bdFlaq = true
}

func setRestaurant(w http.ResponseWriter, r *http.Request) {
	bdFlaq = false
	setHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var restaurant Restaurant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&restaurant); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	price, _ := strconv.Atoi(restaurant.Price)
	numOfGuests, _ := strconv.Atoi(restaurant.NumOfGuests)
	db := GetDatabaseInstance()
	//defer db.connection.Close()
	_, err := db.connection.Exec("insert into restaurants (name, city, img, price, numOfGuests, address, description, googleMap, phone, images) values ($1, $2, $3, $4, $5, $6, $7, $8, $9 , $10)", restaurant.Name, restaurant.City, restaurant.Images[0], price, numOfGuests, restaurant.Address, restaurant.Description, restaurant.GoogleMap, restaurant.Phone, pq.Array(restaurant.Images))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteRestaurant(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var id int
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&id); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	db := GetDatabaseInstance()
	//defer db.connection.Close()
	_, err := db.connection.Exec("delete from restaurants where id_res = $1", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	bdFlaq = false
}

func getTamada(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	if bdTam {
		json.NewEncoder(w).Encode(tamada)
		return
	}
	tamada = tamada[:0]
	db := GetDatabaseInstance()
	//defer db.connection.Close()
	rows, err := db.connection.Query(`select * from tamada`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var name string
		var price int
		var age int
		var description string
		var images pq.StringArray
		var language pq.StringArray
		var instagram string
		if err := rows.Scan(&id, &name, &price, &age, &description, &images, &language, &instagram); err != nil {
			panic(err)
		}
		tam := Tamada{id, name, strconv.Itoa(price), strconv.Itoa(age), description, []string(images), []string(language), instagram}
		tamada = append(tamada, tam)
	}
	//defer rows.Close()
	json.NewEncoder(w).Encode(tamada)
	bdTam = true
}

func setTamada(w http.ResponseWriter, r *http.Request) {
	bdTam = false
	setHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var tam Tamada
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tam); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	price, _ := strconv.Atoi(tam.Price)
	age, _ := strconv.Atoi(tam.Age)
	db := GetDatabaseInstance()
	//defer db.connection.Close()
	_, err := db.connection.Exec("insert into Tamada (name, price, age, description, images, language, instagram) values ($1, $2, $3, $4, $5, $6, $7)", tam.Name, price, age, tam.Description, pq.Array(tam.Images), pq.Array(tam.Language), tam.Instagram)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteTamada(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var id int
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&id); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	db := GetDatabaseInstance()
	//defer db.connection.Close()
	_, err := db.connection.Exec("delete from tamada where id = $1", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	bdTam = false
}
