package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mcaci/fibonacci/fibonacci"
)

func main() {
	serieSize := flag.Int("n", 0, "size of the fibonacci serie")
	fileName := flag.String("o", "", "output file name")
	flag.Parse()
	serie := fibonacci.Serie(*serieSize)
	writer := os.Stdout
	if *serieSize == 0 {
		fmt.Println("FIBONACCI - produce a sequence of the fibonacci serie")
		fmt.Println("")
		fmt.Println("Usage: fibonacci -n <seq length> [-o <output file>]")
		fmt.Println("")
		return
	}
	if *fileName != "" {
		var err error
		writer, err = os.OpenFile(*fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}
	fmt.Fprintf(writer, "%+v\n", serie)
}
