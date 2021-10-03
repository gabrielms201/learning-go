package main

import (
	"fmt"
	"reflect"
)

func main(){ // Infelizmente não da para colocar chaves abaixo da declaração, e sim ao lado ;(
	// Declarando variáveis =)
	var a float32 = 3
	var b float32 = 2
	// Existe um maneira de declarar uma var sem utilizar "var" e a tipagem: -> Não gostei. TEM QUE TIPAR TUDO MEU DEUS
	nome := "Ricardo"

	// Go também tem uma característica positiva/negativa: da erro de compilação ao declarar uma variável e não utilizar
	// (ruim para o comeco do desenvolvimento de uma ideia, tinham que dar a liberdade ao programador se ele quer zoar o código ou não)

	// Output da biblioteca padrão go (fmt)
	fmt.Printf("Olá, mundo\n")
	fmt.Println("a+b: ", sum(a, b))
	fmt.Println("Tipo da variável \"a\": ", reflect.TypeOf(b))
	fmt.Println("Meu nome é: ", nome)
}
// dhr go n precisa de header e da pra declarar a main antes :)
// tipagem dos parâmetros dps do nome, assim como o retorno da função:
func sum(a float32, b float32) float32{
	return a+b
}

