package main

import "fmt"

func main(){
	
 	chave_valor1:= map[string]string{
		 "jose_pereira":"masculino",
		 "pedro_silda":"masculino",
		 "fernanda_souza":"feminino",
	 }

	 for c , v:= range chave_valor1{
		 fmt.Printf("Chave: %v \t Valor: %v \n", c , v)
	 }

	 chave_valor2:= map[string]int{
		"jose_pereira":10,
		"pedro_silda":20,
		"fernanda_souza":30,
	}

	for c , v:= range chave_valor2{
		fmt.Printf("Chave: %v \t Valor: %v \n", c , v)
	}

	chave_valor3:= map[string][]string{
		"jose_pereira":{"azul","preto"},
		"pedro_silda":{"verde","azul"},
		"fernanda_souza":{"branco","amarelo"},
	}

	for c , v:= range chave_valor3{
		fmt.Printf("Chave: %v \t Valor: %v \n", c , v)
	}

}