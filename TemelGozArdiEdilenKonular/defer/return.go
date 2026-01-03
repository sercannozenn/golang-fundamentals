package main

import "fmt"


func normalReturn() int {
	x := 10

	// defer fonksiyonu return çalıştıktan sonra ama
	// fonksiyon gerçekten dönmeden hemen önce çalışır

	defer func(){
		// Bu noktada x artık dönüş değerini etkilemez
		//. çünkü return x ile değer çoktan kopyalanmıştır
		x = 20
	}()

	// 1- x in değerini kopyalar
	// 2- defer çalışır
	// 3- kopyalanmış değer döner
	return x
}

func namedReturn() (x int){
	x = 10

	defer func(){
		x = 20
	}()

	// 1- return değişkeni x'tir
	// 2- defer çalışır -> x = 20
	// 3- x döner

	return
}


func main(){

	fmt.Println("normalReturn : ",  normalReturn())
	fmt.Println("namedReturn : ",  namedReturn())
}