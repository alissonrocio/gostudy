package main

import "fmt"

func main(){

	slice_de_int:=[]int{1,2,3,4,5,6}

	fmt.Println("********* Estruturas de dados *********")

	// Loop tradicional
	fmt.Println("Repetição tradicional")
	for i:=0 ; i < len(slice_de_int) ; i++{
		fmt.Printf("Posição %v \t Valor %v \n",i,slice_de_int[i])
	}

	// Loop com range
	fmt.Println("Repetição com range")
	for p , v:= range slice_de_int{
		fmt.Printf("Posição %v \t Valor %v \n",p,v)
	}

	// While
	fmt.Println("Repetição com while")
	repeticao:= 0
	for repeticao < 10{
		fmt.Printf("Repetição %v \n",repeticao)
		repeticao++
	}

	// if / else / else if
	valor:= 9
	if valor <= 10 {
		fmt.Println("Valor menor que 10")
	} else if valor > 10 {
		fmt.Println("Valor maior que 10")
	}

	// switch variação 1
	cor1:= "azul"
	switch cor1 {
		case "azul":
			fmt.Println("Cor escolhida é azul")
		case "preto":
			fmt.Println("Cor escolhida é preto")
		case "amarelo":
			fmt.Println("Cor escolhida é amarelo")
		default:
			fmt.Println("Cor não encontrada")
	}

	cor2:= "roxo"
	switch {
		case cor2 == "azul":
			fmt.Println("Cor escolhida é azul")
		case cor2 == "preto":
			fmt.Println("Cor escolhida é preto")
		case cor2 == "amarelo":
			fmt.Println("Cor escolhida é amarelo")
		default:
			fmt.Println("Cor não encontrada")
	}
}