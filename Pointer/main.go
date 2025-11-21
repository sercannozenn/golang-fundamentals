package main

import "fmt"

// Pointer işaretçi
// Bir değerin bellekteki adresini tutan özel değişken türüdür.
func main() {
	number :=10
	fmt.Println("number değeri: ", number)
	fmt.Println("number adresi (&number): ", &number)

	// Bir pointer başka bir değişkenin adresini tutar.

	var ptr *int
	ptr = &number

	fmt.Println("\nPointer değişkeni ptr adresi:", ptr)
	fmt.Println("Pointerın gösterdiği değer (*ptr):", *ptr)

	*ptr = 99 //adreste tutulan değeri günceller
	fmt.Println("\nPointer ile güncellendikten sonra:")
	fmt.Println("number:", number)
	fmt.Println("*ptr:", *ptr)
	fmt.Println("ptr:", ptr)
	fmt.Println("&ptr adres:", &ptr)
	fmt.Println("&number adres:", &number)

	fmt.Println("Bu ikisi aynı mı?", ptr == &number)

	fmt.Println("***************************")

	value := 50
	updateValue(value)
	fmt.Println("Fonksiyona değer kopyası gönderildi:", value)

	updatePointer(&value)
	fmt.Println("Pointer gönderidi:", value)

}
func updateValue(v int){
	v= v+10
	fmt.Println("[updateValue] v değeri:", v)
}
func updatePointer(v *int){
	*v = *v + 10 // * adresin içindeki değere erişim
	fmt.Println("[updatePointer] v değeri:", *v)
}