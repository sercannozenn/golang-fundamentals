package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Go HTTP Server Başlıyor...")
	fmt.Println("Tarayıcıya yaz: http://localhost:8081")

	// Route tanımlama
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/json", JsonHandler)

	//Server Başlatma
	// ListenAndServe(port,handler/router)
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/sercan", SercanHandler)

	err := http.ListenAndServe(":8080", nil)
	err2 := http.ListenAndServe(":8081", mux)
	if err !=nil {
		fmt.Println("Server hata verdi:", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Anasayfaya Hoş Geldiniz!")
}
func SercanHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Sercan Özen")
}

func HelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Merhaba dünya!")
}
func JsonHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	jsonData := `{"message": "Go HTTP Server Temlleri dersinden selamlar!"}`
	fmt.Fprintln(w,jsonData)
}
