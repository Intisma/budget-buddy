/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"budget-buddy/cmd"

	_ "budget-buddy/cmd/moneyType"
	_ "budget-buddy/cmd/periodic"
	_ "budget-buddy/cmd/transfer"
)

func main() {
	cmd.Execute()
}
