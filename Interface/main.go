package main

import "fmt"

// Interface bir davranış(method) imzasını tanımlar.
// Yani "bir davranışı gerçekleştiren her yapı, bu interfacei uygular"
type Greeter interface{
	Greet() // Herhangi ir şey, Greet() fonksiyonunu içeriyorsa Greeter sayılır.
}
type TurkishGreeter struct{}
type EnglishGreeter struct{}

func (t TurkishGreeter) Greet(){
	fmt.Println("Merhaba")
}
func (e EnglishGreeter) Greet(){
	fmt.Println("Hello")
}

func SayHello(g Greeter){
	fmt.Println("Selamlama başladı")
	g.Greet()
}

func main()  {
	tr := TurkishGreeter{}
	en := EnglishGreeter{}

	SayHello(tr)
	SayHello(en)
}