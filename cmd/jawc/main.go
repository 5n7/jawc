package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ikawaha/kagome/tokenizer"
	"github.com/skmatz/jawc"
)

func run() error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	app := jawc.NewApp(tokenizer.Search)
	fmt.Println(app.CountWords(string(b)))

	return nil
}

func main() {
	exitCode := 0

	if err := run(); err != nil {
		exitCode = 1
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(exitCode)
}
