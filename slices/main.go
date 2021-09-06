package main

import "fmt"

func main(){
	fmt.Println("********* Slices *********")

	slice1:= []int{1,2,3,4,5}

	slice2:= []int{6,7,8,9,10}

	slice3:= make([]int,5,10)

	slice3 = []int{1,2,3,4,5}

	slice4 := [][]int{ 
		{1,3,4} , 
		{100,200,300}, 
	}

	fmt.Printf("Slice todas posições %v \n",slice1)

	fmt.Printf("Slice todas posições %v \n",slice1[:])

	fmt.Printf("Slice fim index 2 %v \n",slice1[:2])

	fmt.Printf("Slice início index 3 até o fim %v \n",slice1[3:])

	fmt.Println("Adicionando item ao slice" , append(slice1,6,7,8,9,10))

	fmt.Println("Adicionando slice ao slice" , append(slice1,slice2...))

	fmt.Println("Removendo itens do slice " , append(slice1[:2],slice1[3:]...))

	fmt.Println("Slice usando o make" , slice3)

	fmt.Println("Slice multi dimensional" , slice4)

}