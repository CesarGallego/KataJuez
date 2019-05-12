package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goJuez [expected] [reported]",
	Short: "goJuez compare kata files",
	Run:   goJuez,
}

func fileReader(file string, ch chan []byte) {
	f, err := os.Open(file)
	if err != nil {
		f := fmt.Errorf("%s", err)
		fmt.Println(f)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		ch <- scanner.Bytes()
	}
	close(ch)
}

func goJuez(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		cmd.Help()
		os.Exit(1)
	}
	expChan := make(chan []byte, 10)
	go fileReader(args[0], expChan)
	repChan := make(chan []byte, 10)
	go fileReader(args[1], repChan)

	var (
		exp        []byte
		rep        []byte
		expContent = true
		repContent = true
	)

	for {
		exp, expContent = <-expChan
		rep, repContent = <-repChan
		if expContent != repContent || !bytes.Equal(exp, rep) {
			os.Exit(1)
		}
		if !expContent || !repContent {
			break
		}
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
