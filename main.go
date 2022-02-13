package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TKMAX777/kiritan_handler"
)

func main() {
	kiritan, err := kiritan_handler.New()
	if err != nil {
		panic(err)
	}

	fmt.Println("Getting KIRITAN handler success!")

	var scan = bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		var text = scan.Text()
		kiritan.SetText(text)
		kiritan.Play()
	}
}
