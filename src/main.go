/*
Copyright © 2024 HEITOR FREITAS FERREIRA <heitor.ff@hotmail.com>
*/
package main

import (
	"fmt"
	"os"
	"tcc/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
