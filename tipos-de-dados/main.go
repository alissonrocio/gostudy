package main

import "fmt"

// Variáveis - package scope

// Declaração
var sou_um_string string
var sou_um_int int
var sou_um_float float64

// Declaraçao e inicialização 
var sou_um_bool bool = true
var sou_um_slice_string []string = []string{"sou","um","slice"}
var sou_um_rune rune = 123 // rune é um alias para int32
var sou_um_signed_int int32 = -1234 
var sou_um_unsigned_int uint32 = 1234 // apenas número positivos

// Constantes
const 
( 
	sou_um_contante_float float64 = 3.14
	sou_um_constante_string = "https://url_do_meu_servico.com"
)

func main(){
	
	fmt.Println("********* Tipos básicos de dados *********")

	//Dados vazios
	fmt.Println("Valor Default") 
	Formatar(sou_um_string)
	Formatar(sou_um_int)
	Formatar(sou_um_float)
	Formatar(sou_um_bool)
	Formatar(sou_um_slice_string)
	Formatar(sou_um_rune)
	Formatar(sou_um_constante_string)

	//Inicialização das variáveis já declaradas
	sou_um_string = "Um texto qualqer"
	sou_um_int = 100
	sou_um_float = 1.234

	//Print na tela
	fmt.Println("Valor Preenchido - declaracao var ")
	Formatar(sou_um_string)
	Formatar(sou_um_int)
	Formatar(sou_um_float)
	Formatar(sou_um_bool)
	Formatar(sou_um_slice_string)
	Formatar(sou_um_rune)

	// Variáveis - function scope

	//Declaração e Inicialização usando o :=
	sou_outro_string:= "Outro texto qualquer"
	sou_outro_int := 100
	sou_outro_float := 2.345
	sou_outro_bool := true
	sou_outro_slice_string := []string{"Valor1","Valor2","Valor3"}
	sou_outro_rune := 11111

	fmt.Println("Valor Preenchido - declaração e inicialização :=")
	Formatar(sou_outro_string)
	Formatar(sou_outro_int)
	Formatar(sou_outro_float)
	Formatar(sou_outro_bool)
	Formatar(sou_outro_slice_string)
	Formatar(sou_outro_rune)
}

func Formatar(tipo interface{}){
	fmt.Printf("\tTipo: %T \t Valor: %v \n"  , tipo , tipo)
}