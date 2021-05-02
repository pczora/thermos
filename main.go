package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/knq/escpos"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: tprint font fontwidth fontheight text")
		return
	}

	fontWidth, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("fontwidth must be an integer")
	}

	fontHeight, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("fontheight must be an integer")
	}

	printAndCut(os.Args[1], uint8(fontWidth), uint8(fontHeight), os.Args[4])
}

func printAndCut(font string, width uint8, height uint8, text string) {
	printerConnection, err := net.Dial("tcp", "192.168.178.37:9100")
	if err != nil {
		fmt.Printf("Error creating connection: %s\n", err)
		return
	}
	defer printerConnection.Close()

	reader := bufio.NewReader(printerConnection)
	writer := bufio.NewWriter(printerConnection)
	w := bufio.NewReadWriter(reader, writer)
	p := escpos.New(w)

	//printerConnection.Write([]byte("\x1B@"))
	p.Init()
	p.SetSmooth(1)
	p.SetFontSize(width, height)
	p.SetFont(font)
	p.Write(text)
	p.FormfeedN(5)

	p.Cut()
	p.End()

	w.Flush()
}
