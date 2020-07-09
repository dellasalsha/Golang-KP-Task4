package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Profile struct {
	Clients struct {
		Address struct {
			City   string `json:"city"`
			State  string `json:"state"`
			Street string `json:"street"`
			Zip    string `json:"zip"`
		} `json:"address"`
		Age      int64  `json:"age"`
		Company  string `json:"company"`
		Email    string `json:"email"`
		Gender   string `json:"gender"`
		ID       string `json:"id"`
		IsActive string `json:"isActive"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
	} `json:"clients"`
}

type Coloring struct {
	Colors struct {
		Category string `json:"category"`
		Code     struct {
			Hex  string `json:"hex"`
			Rgba string `json:"rgba"`
		} `json:"code"`
		Color string `json:"color"`
		Type  string `json:"type"`
	} `json:"colors"`
	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
}

func getProfiles(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Profile

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	Kota := request.Clients.Address.City
	Negara := request.Clients.Address.State
	Alamat := request.Clients.Address.Street
	KodePos := request.Clients.Address.Zip
	Usia := request.Clients.Age
	Perusahaan := request.Clients.Company
	Email := request.Clients.Email
	JenisKelamin := request.Clients.Gender
	ID := request.Clients.ID
	Status := request.Clients.IsActive
	Nama := request.Clients.Name
	Telepon := request.Clients.Phone

	stmt, err := db.Prepare("INSERT INTO profiles (ID,Name,Age,Gender,City,Region,Address,PostalCode,Company,Email,isActive,Phone) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(ID, Nama, Usia, JenisKelamin, Kota, Negara, Alamat, KodePos, Perusahaan, Email, Status, Telepon)

	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

func getColors(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var req Coloring

	if err = json.Unmarshal(body, &req); err != nil {
		fmt.Println("Failed decoding json message")
	}

	Kategori := req.Colors.Category
	Hexa := req.Colors.Code.Hex
	Rgb := req.Colors.Code.Rgba
	Warna := req.Colors.Color
	Tipe := req.Colors.Type
	Tinggi := req.Thumbnail.Height
	Link := req.Thumbnail.URL
	Lebar := req.Thumbnail.Width

	stmt, err := db.Prepare("INSERT INTO colors (Color,Category,Type,Rgba,Hex,URL,Width,Height) VALUES(?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(Warna, Kategori, Tipe, Rgb, Hexa, Link, Lebar, Tinggi)

	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_task4")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	//fmt.Println("Server on :8080")

	// Route handles & endpoints
	r.HandleFunc("/getProfiles", getProfiles).Methods("POST")
	r.HandleFunc("/getColors", getColors).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))

}
