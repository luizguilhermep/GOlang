package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 432, saldo: 3221.645}

	fmt.Println(contaGuilherme)

}
