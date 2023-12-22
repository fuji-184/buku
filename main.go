package main

import (
	"database/sql"
	"encoding/json"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
)

//go:embed all:views/build
var PembacaUI embed.FS

type Buku struct {
	Nama   string  `json:"nama"`
	Gambar string  `json:"gambar"`
	Rating float64 `json:"rating"`
	Tahun  string  `json:"tahun"`
	Harga  float64 `json:"harga"`
}

type Kategori struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type Soal struct {
	Id      int    `json:"id"`
	Judul   string `json:"judul"`
	Kategori int   `json:"kategori"`
	Isi     []struct {
		Question string   `json:"question"`
		Answer   int      `json:"answer"`
		Options  []string `json:"options"`
	} `json:"Isi"`
}

var db *sql.DB
var ServeUI http.Handler

func InitDB() error {
	var err error
	db, err = sql.Open("mysql", "root:1234@/uas2")
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil
}

func InitUI() {
	UI, err := fs.Sub(PembacaUI, "views/build")
	if err != nil {
		log.Fatal(err)
	}

	ServeUI = http.FileServer(http.FS(UI))
}

func main() {
	if err := InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	InitUI()

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Handle("/", ServeUI)
	r.Handle("/_app/*", ServeUI)

	r.Get("/buku", buku)
	r.Get("/buku/{judul}-{penulis}", cari)
	r.Get("/bukuterbaru", newbuku)
	r.Get("/topbuku", topbuku)
	r.Get("/kategori", kategori)
	r.Get("/search/{ktg}-{min}-{max}", filter)
	r.Get("/soal", soal)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/build/index.html")
	})

	port := 3000
	fmt.Printf("Server is running on :%d...\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal(err)
	}
}

func buku(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("SELECT b.Nama, b.Gambar, b.Tahun, b.Harga, AVG(r.Rating) AS rating FROM Buku b JOIN Rating r ON b.Id = r.Id_Buku GROUP BY b.Id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	hasil, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer hasil.Close()

	var data []Buku
	for hasil.Next() {
		var i Buku
		err := hasil.Scan(&i.Nama, &i.Gambar, &i.Tahun, &i.Harga, &i.Rating)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, i)
	}

	if len(data) == 0 {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func cari(w http.ResponseWriter, r *http.Request) {
	judul := chi.URLParam(r, "judul")
	penulis := chi.URLParam(r, "penulis")
	
	var stmt *sql.Stmt
	var err error

	if penulis == "1" {
	        stmt, err = db.Prepare("SELECT b.Nama, b.Gambar, b.Harga, b.Tahun, AVG(r.Rating) AS rating FROM Buku b JOIN Rating r ON b.Id = r.Id_Buku WHERE b.Nama LIKE ? GROUP BY b.Id")
	}
	
	if judul == "1" {
	        stmt, err = db.Prepare("SELECT b.Nama, b.Gambar, b.Harga, b.Tahun, AVG(r.Rating) AS rating FROM Buku b JOIN Rating r ON b.Id = r.Id_Buku WHERE b.Penulis in (SELECT Id FROM Penulis WHERE Nama like ?) GROUP BY b.Id")
	}
	
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	
	var hasil *sql.Rows

	if penulis == "1" {
	        hasil, err = stmt.Query("%"+judul+"%")
	}
	
	if judul == "1" {
	        hasil, err = stmt.Query("%"+penulis+"%")
	}
	
	if err != nil {
		log.Fatal(err)
	}
	defer hasil.Close()

	var data []Buku
	for hasil.Next() {
		var i Buku
		err := hasil.Scan(&i.Nama, &i.Gambar, &i.Harga, &i.Tahun, &i.Rating)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, i)
	}

	if len(data) == 0 {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func newbuku(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("SELECT b.Nama, b.Gambar, b.Tahun, b.Harga, AVG(r.Rating) AS rating FROM Buku b JOIN Rating r ON b.Id = r.Id_Buku GROUP BY b.Id order by tahun desc")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	hasil, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer hasil.Close()

	var data []Buku
	for hasil.Next() {
		var i Buku
		err := hasil.Scan(&i.Nama, &i.Gambar, &i.Tahun, &i.Harga, &i.Rating)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, i)
	}

	if len(data) == 0 {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func topbuku(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("SELECT b.Nama, b.Tahun, b.Gambar, b.Harga, AVG(r.Rating) AS rating FROM Buku b JOIN Rating r ON b.Id = r.Id_Buku GROUP BY b.Id ORDER BY rating DESC;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	hasil, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer hasil.Close()

	var data []Buku
	for hasil.Next() {
		var i Buku
		err := hasil.Scan(&i.Nama, &i.Tahun, &i.Gambar, &i.Harga, &i.Rating)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, i)
	}

	if len(data) == 0 {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", 	"application/json")
	json.NewEncoder(w).Encode(data)
}

func kategori(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("SELECT Id, Nama from Kategori;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	hasil, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer hasil.Close()

	var data []Kategori
	for hasil.Next() {
		var i Kategori
		err := hasil.Scan(&i.Id, &i.Nama)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, i)
	}

	if len(data) == 0 {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func filter(w http.ResponseWriter, r *http.Request) {
        ktg := chi.URLParam(r, "ktg")
 	min := chi.URLParam(r, "min")
 	max := chi.URLParam(r, "max")

	stmt, err := db.Prepare("SELECT b.Nama, b.Gambar, b.Tahun, b.Harga, AVG(r.Rating) AS rating FROM Buku b JOIN Rating r ON b.Id = r.Id_Buku where b.Kategori = ? AND b.Harga BETWEEN ? AND ? GROUP BY b.Id;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	hasil, err := stmt.Query(&ktg, &min, &max)
	if err != nil {
		log.Fatal(err)
	}
	defer hasil.Close()

	var data []Buku
	for hasil.Next() {
		var i Buku
		err := hasil.Scan(&i.Nama, &i.Gambar, &i.Tahun, &i.Harga, &i.Rating)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, i)
	}

	if len(data) == 0 {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func soal(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("SELECT Id, Judul, Kategori, JSON_UNQUOTE(JSON_EXTRACT(Isi, '$')) AS Isi FROM Soal;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	hasil, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer hasil.Close()

	var soals []Soal
	for hasil.Next() {
		var soal Soal
		err := hasil.Scan(&soal.Id, &soal.Judul, &soal.Kategori, &soal.Isi)
		if err != nil {
			panic(err.Error())
		}

		// Decode string JSON menjadi struktur data Golang
		err = json.Unmarshal([]byte(soal.Isi), &soal.Isi)
		if err != nil {
			panic(err.Error())
		}

		soals = append(soals, soal)
	}

	if len(soals) == 0 {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(soals)
}
