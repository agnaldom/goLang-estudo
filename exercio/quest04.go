/*
*Faça um Programa que peça as 4 notas bimestrais e mostre a média. 
*/
package main

import "fmt"

func main(){
	var a, b, c, d int
	fmt.Println("Entre com 4 notas bimestrais: ")
	fmt.Scan(&a, &b, &c, &d)
	media := (a+b+c+d)/4
	fmt.Println("A media :  ", media)
}


