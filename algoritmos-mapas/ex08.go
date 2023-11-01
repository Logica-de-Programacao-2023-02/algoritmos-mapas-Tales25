//Escreva uma função que receba um mapa onde as chaves são nomes de pessoas e os valores são as quantias de dinheiro
//que cada pessoa gastou em uma conta compartilhada. A função deve calcular o valor que cada
//pessoa deve receber ou pagar para igualar as despesas.

package main

import "fmt"

func main() {
	expensesMap := map[string]float64{
		"Tawe":  250,
		"Duds":  300,
		"Rafa":  150,
		"Benio": 200,
		"Breno": 100,
	}

	equalizedExpensesMap := equalizeExpenses(expensesMap)

	fmt.Println(expensesMap)
	fmt.Println(equalizedExpensesMap)
}

func equalizeExpenses(expensesMap map[string]float64) map[string]float64 {
	var totalExpenses, equalizedExpenses, resultToEqualize float64

	equalizedExpensesMap := make(map[string]float64)

	for _, valor := range expensesMap {
		totalExpenses += valor
	}

	quantidadeChaves := len(expensesMap)
	equalizedExpenses = totalExpenses / float64(quantidadeChaves)

	for chave, valor := range expensesMap {
		switch {
		//pagou mais do que deveria, deve ser retirado X do que já tinha pago
		case valor > equalizedExpenses:
			resultToEqualize = (valor - equalizedExpenses) * -1
			equalizedExpensesMap[chave] = resultToEqualize

		//pagou menos do que deveria, precisa pagar X a mais
		case valor < equalizedExpenses:
			resultToEqualize = equalizedExpenses - valor
			equalizedExpensesMap[chave] = resultToEqualize

		//pagou a quantia certa, não precisa pagar nada a mais nem menos
		default:
			equalizedExpensesMap[chave] = 0
		}
	}

	return equalizedExpensesMap
}
