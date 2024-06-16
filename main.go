/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Intisma/budget-buddy/cmd"

	_ "github.com/Intisma/budget-buddy/cmd/moneyType"
	_ "github.com/Intisma/budget-buddy/cmd/periodic"
	_ "github.com/Intisma/budget-buddy/cmd/transfer"
)

func main() {
	cmd.Execute()
}
