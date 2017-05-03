/*
* Faça um programa que peça dois número e imprima a soma
*/
package main

import "fmt"

func main(){
	var i,n int
	fmt.Println("Entre com dois números: ")
	fmt.Scan(&i,&n)
	soma := i+n
	fmt.Println("A soma dos números é: ", soma)
}
