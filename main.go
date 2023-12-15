package main

import (
	"net/http"
	"io/fs"
	"embed"
	"log"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	
	_ "github.com/go-sql-driver/mysql"
)

//go:embed all:views/build

var PembacaUI embed.FS

type Laporan struct {
        Id int `json:"id"`
        Isi string `json:"isi"`
        Tanggal string `json:"tanggal"`
}

type Pegawai struct {
        Id int `json:"id"`
        Nama string `json:"nama"`
        Umur int `json:"umur"`
}

var db *sql.DB

var ServeUI http.Handler

func InitDB(){
	db, _ = sql.Open("mysql", "root:1234@/uas")

        db.SetConnMaxLifetime(time.Minute * 3)
        db.SetMaxOpenConns(10)
        db.SetMaxIdleConns(10)
}

func InitUI(){
        UI, err := fs.Sub(PembacaUI, "views/build")
        if err != nil {
                log.Fatal(err)
        }
        
        ServeUI = http.FileServer(http.FS(UI))
}

func main() {
        InitDB()
        InitUI()
                
	r := chi.NewRouter()
	
	r.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))
	
 	r.Handle("/", ServeUI)
 	r.Handle("/_app/*", ServeUI)
 	
 	r.Get("/laporan/{tgl}", laporan)
 	r.Get("/pegawai/{umur}-{nama}", pegawai)
 	
  	 r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
	        r.URL.Path = "/"
	        ServeUI.ServeHTTP(w, r)
        })
        
        
 
	http.ListenAndServe(":3000", r)
}

func laporan(w http.ResponseWriter, r *http.Request){
        
       tgl := chi.URLParam(r, "tgl")
 	        
 	stmt, err := db.Prepare("select Id, Isi, Tanggal from Laporan where date(Tanggal) = ? ")
 	        if err != nil {
 	                log.Fatal(err)
 	        }
 	        defer stmt.Close()
 	        
 	        hasil, err := stmt.Query(tgl)
 	        if err != nil {
 	                log.Fatal(err)
 	        }
 	        defer hasil.Close()
 	        
 	        var data []Laporan
 	        for hasil.Next(){
 	                var i Laporan
 	                err := hasil.Scan(&i.Id, &i.Isi, &i.Tanggal)
 	                if err != nil {
 	                        log.Fatal(err)
 	                }
 	                
 	                data = append(data, i)
 	        }
 	        
 	        if len(data) == 0{
 	              http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
 	              
 	        }

  	        

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)


 	        
 	}
 	
func pegawai(w http.ResponseWriter, r *http.Request){
        
        umur := chi.URLParam(r, "umur")
        nama := chi.URLParam(r, "nama")
        
        stmt, err := db.Prepare("SELECT * FROM Pegawai WHERE Nama LIKE ? OR Umur LIKE ?;")
 	        if err != nil {
 	                log.Fatal(err)
 	        }
 	        defer stmt.Close()
 	        
 	        hasil, err := stmt.Query("%" + nama + "%", "%" + umur + "%")
 	        if err != nil {
 	                log.Fatal(err)
 	        }
 	        defer hasil.Close()
 	        
 	        
 	        var data []Pegawai
 	        for hasil.Next(){
 	                var i Pegawai
 	                err := hasil.Scan(&i.Id, &i.Nama, &i.Umur)
 	                if err != nil {
 	                        log.Fatal(err)
 	                }
 	                
 	                data = append(data, i)
 	        }
 	        
 	        if len(data) == 0{
 	              http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
 	              
 	        }
 	        

  	        

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
        
}