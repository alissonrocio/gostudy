package main

import (
	"fmt"
)

type pessoa struct{
	nome string
	idade int
	telefones []string
	endereco
}

type endereco struct{
	estado string
	cidade string
}

func main(){

	pessoas:= []pessoa{
		{
			nome: "Jose Pereira",
			idade:30, 
			telefones: []string{"5555-1111","4444-1111"},
			endereco: endereco{
				estado: "Minas Gerais",
				cidade: "Ipatinga",
			},
		},
		{
			nome: "Pedro da Silva",
			idade:40, 
			telefones: []string{"5555-1111","4444-1111"},
			endereco: endereco{
				estado: "São Paulo",
				cidade: "São Paulo",
			},
		},
		{
			nome: "Marcela Sampaio",
			idade:24, 
			telefones: []string{"5555-1111","4444-1111"},
			endereco: endereco{
				estado: "Rio de Janeiro",
				cidade: "Rio de Janeiro",
			},
		},
	}

	for _ , v:= range pessoas{
		fmt.Printf("Meu nome é %v , tenho %v anos e sou do estado de %v e moro na cidade de %v , meu telefone é %v\n",
		v.nome,v.idade,v.endereco.estado,v.endereco.cidade,v.telefones[0])
	}

}