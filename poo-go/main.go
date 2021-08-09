package main

import (
	"fmt"
	"poo-go/contas"
)

func main() {
	contaDouglas := contas.ContaCorrente{Titular: "Douglas", Saldo: 225.5}
	contaEdgar := contas.ContaCorrente{Titular: "Edgar", Saldo: 300}

	satatus := contaDouglas.Transferir(200, &contaEdgar)
	fmt.Println(satatus)

	fmt.Println(contaDouglas)
	fmt.Println(contaEdgar)
}
