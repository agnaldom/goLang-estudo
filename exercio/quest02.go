/*
* Faça um programa que peça um número e então mostre a mensagem 
* "O número informado foi [numero]"
*/
package main

import "fmt"

func main(){
	var i int
	fmt.Println("Entre com um número: ")
	fmt.Scan(&i)
	fmt.Println("O número informado foi", i)
}
