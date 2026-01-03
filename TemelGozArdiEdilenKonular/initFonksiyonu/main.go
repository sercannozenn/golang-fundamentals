package main

import "fmt"

// init() Go’da program başlamadan hemen önce, otomatik çalışan bir fonksiyondur.
func init(){
	// init fonksiyonu için constructors gibi düşünebilirsiniz.

	// init() şunlar için kullanılır:
	// Ayar (config) yüklemek
	// Global değişkenleri hazırlamak
	// DB bağlantısını kurmak
	// Log sistemini başlatmak	
	// Ortam kontrolü yapmak

	// Bir dosyada birden fazla init() fonksiyonu olabilir. 
	// Çalışma sırası dosya içindeki sıralarına göre belirlenir.
	// init() fonksiyonu parametre almaz ve geriye değer döndürmez.
	// import edilen paketlerdeki init() fonksiyonları öncelikli çalışır.
	fmt.Println("init çalıştı")
}
func init(){
	fmt.Println("init2 çalıştı")
}
func init(){
	fmt.Println("init3 çalıştı")
}

func main(){
	fmt.Println("ana fonksiyon çalıştı")
}
