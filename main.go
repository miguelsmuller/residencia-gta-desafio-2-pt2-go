/*
Larisse Rodrigues
Marcel Pontes
Miguel Muller
Taís Amanda
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type tickets struct {
	name  string
	qnt   int
	total float32
}

func main() {
	const price = float32(10.50)
	listaTickets := make(map[string]tickets)

	fmt.Println("EMPRESA AREA XPTO")
	for {
		var name string
		var qnt int

		fmt.Print("\nQual é o seu nome? ")
		fmt.Scanln(&name)
		fmt.Println("Nome:", name)

		fmt.Print("\nQuantos tickets deseja adquirir? ")
		fmt.Scanln(&qnt)
		fmt.Println("Quantidade:", qnt)

		listaTickets[randomString()] = tickets{
			name:  name,
			qnt:   qnt,
			total: float32(qnt) * price,
		}
		fmt.Println("Ticket adicionado com sucesso!!!!")

		resposta := "S"
		fmt.Print("\nQuer reservar mais tickets? (S/N) ")
		fmt.Scanln(&resposta)

		if strings.ToUpper(resposta) == "N" {
			break
		}
	}

	var total = float32(0)
	for index := range listaTickets {
		//total += price * float32(listaTickets[index].qnt)
		total += listaTickets[index].total
	}

	fmt.Println("\nLista de Compras:", listaTickets)
	fmt.Printf("\nTotal: %6.2f", total)
}

func randomString() string {
	charSet := "abcdedfghijklmnopqrst1234567890"
	var output strings.Builder
	length := 6
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
