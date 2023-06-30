package main

import (
	"fmt"
	"math"
)

// q.4

func updateIntValue(ptr *int) {
	num := *ptr
	UltimoDigito := num % 10
	num /= 10
	SegundoUltimoDigito := num % 10
	sum := UltimoDigito + SegundoUltimoDigito
	*ptr = sum
}

func main() {
	num := 1234
	updateIntValue(&num)
	fmt.Println("Updated value:", num)
}

// q.5

func updateFloatValue(ptr *float64) {
	pi := math.Pi
	avg := (*ptr + pi) / 2
	*ptr = avg
}

func main() {
	num := 3.14
	updateFloatValue(&num)
	fmt.Println("Updated value:", num)
}

// q.6

type Livro struct {
	Titulo string
	Autor  string
}

func atualizarTituloLivro(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro A",
		Autor:  "Autor A",
	}
	atualizarTituloLivro(&livro)
	fmt.Println("Título do livro:", livro.Titulo) // Saída: Título do livro: Livro A

	livroAnonimo := Livro{
		Titulo: "Livro B",
		Autor:  "Anônimo",
	}
	atualizarTituloLivro(&livroAnonimo)
	fmt.Println("Título do livro anônimo:", livroAnonimo.Titulo) // Saída: Título do livro anônimo: Desconhecido
}

// q.7

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValorConta(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	conta := Conta{
		Saldo:   100.0,
		Titular: "Lucas",
	}
	adicionarValorConta(&conta, 50.0)
	fmt.Println("Saldo da conta:", conta.Saldo) // Saída: Saldo da conta: 150.0
}

// q.8

func modifyValue(ptr *int) {
	*ptr = 42
}

func main() {
	ptr := new(int)
	*ptr = 10

	fmt.Println("Valor antes da modificação:", *ptr)

	modifyValue(ptr)

	fmt.Println("Valor após a modificação:", *ptr)
	free(ptr)
}

func free(ptr *int) {
	ptr = nil
}

// q.9

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livro) {
	desconto := livro.Preco * 0.1
	livro.Preco -= desconto
}

func main() {
	livro := Livro{
		Titulo: "Diário de Banana",
		Autor:  "Jeff Kinney",
		Preco:  50.0,
	}

	aplicarDesconto(&livro)

	fmt.Println("Livro:", livro.Titulo)
	fmt.Println("Autor:", livro.Autor)
	fmt.Println("Preço com desconto:", livro.Preco) // Saída: Preço com desconto: 45.0
}

// q.10

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func fillPrimes(slice *[]int, size int) {
	count := 0
	num := 2

	for count < size {
		if isPrime(num) {
			*slice = append(*slice, num)
			count++
		}
		num++
	}
}

func main() {
	var primes []int
	size := 10

	fillPrimes(&primes, size)

	fmt.Println("Primeiros", size, "números primos:", primes)
}
