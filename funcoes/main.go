package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
}

func main(){
	fmt.Println("olá")

	defer funcao_sem_retorno()
	funcao_retorno_simples()
	funcao_retorno_composto()
	funcao_parametro_a(100)
	funcao_parametro_b([]int {10,20,30,40})
	funcao_varios_parametros(10,20,30,40)

	func (){
		fmt.Println("Sou uma função anônima")
	}()

	p:= pessoa{"Alisson","Rodrigues"}
	p.mostrar_info()

	variavel_funcao:= func () int {
		fmt.Println("Sou uma função dentro da variável")
		return 9999
	}
	fmt.Println(variavel_funcao())

	variavel_funcao_funcao:= funcao_dentro_funcao()
	fmt.Println(variavel_funcao_funcao())
	fmt.Println(funcao_dentro_funcao()())

	callback_a(funcao_alguma_coisa_a)

	callback_b(funcao_alguma_coisa_b)
}

func funcao_sem_retorno(){
	fmt.Println("Sou uma função sem retorno")
}

func funcao_retorno_simples() int{
	fmt.Println("Sou uma função com retorno " , 100)
	return 100
}

func funcao_retorno_composto() (bool,string) {
	fmt.Println("Sou uma função com retorno composto " , true , "deu erro")
	return true , "deu erro"
}

func funcao_parametro_a(p1 int) int {
	fmt.Println("Sou uma função com o parametro " , p1)
	return p1
}

func funcao_parametro_b(slice []int) (int,int) {
	fmt.Println("Sou uma função com parametro " , slice)
	soma1:=0
	soma2:=0
	for i , v:= range slice{
		soma1+= slice[i]
		soma2+= v
	}
	return soma1 , soma2
}

func funcao_varios_parametros(inteiros ...int) int {
	fmt.Println("Sou uma função que aceito vários parametros" , inteiros)
	soma:=0
	for _ , v:= range inteiros{
		soma+= v
	}
	return soma
}

func (p pessoa) mostrar_info() string {
	fmt.Println(p.nome + " " + p.sobrenome)
	return p.nome + " " + p.sobrenome
}

func funcao_dentro_funcao() func() int {
	return func() int{
		fmt.Println("Estou dentro de outra função")
		return 12345
	}
}

func funcao_alguma_coisa_a(){
	fmt.Println("callback a")
}

func funcao_alguma_coisa_b() int{
	fmt.Println("callback b")
	return 1
}

func callback_a(f func()){
	f()
}

func callback_b(f func() int){
	f()
}

 

