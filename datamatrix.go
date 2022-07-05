package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/pschlump/filelib"
)

var outFile = flag.String("o", "out.png", "output .png file")
var txt = flag.String("s", "", "text to put into barcode")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "godatamatrix: Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(1)
	}

	datamatrix, err := datamatrix.Encode(*txt) // Create the datamatrix barcode
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// Scale the barcode to 200x200 pixels
	datamatrix, err = barcode.Scale(datamatrix, 200, 200)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// create the output file
	fn := *outFile
	file, err := filelib.Fopen(fn, "w")
	if err != nil {
		fmt.Printf("Unable to open %s for output, Error: %s\n", fn, err)
		os.Exit(1)
	}
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, datamatrix)
}
