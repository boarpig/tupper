package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
)

var fill, empty byte
var imgtext string = "960939379918958884971672962127852754715004339660129306651505519271702802395266424689642842174350718121267153782770623355993237280874144307891325963941337723487857735749823926629715517173716995165232890538221612403238855866184013235585136048828693337902491454229288667081096184496091705183454067827731551705405381627380967602565625016981482083418783163849115590225610003652351370343874461848378737238198224849863465033159410054974700593138339226497249461751545728366702369745461014655997933798537483143786841806593422227898388722980000748404719"

func init() {
	fill_s := flag.String("pixel", "#", "Character to use for filled pixel.")
	empty_s := flag.String("empty", " ", "Character to use for empty pixel.")
	flag.StringVar(&imgtext, "var", imgtext, "Number to read image from")
	fill = (*fill_s)[0]
	empty = (*empty_s)[0]
	flag.Parse()
	if len(*fill_s) > 1 {
		fmt.Println("Error: Fill pixel: Can be only one character long")
		os.Exit(1)
	}
	if len(*empty_s) > 1 {
		fmt.Println("Error: Empty pixel: Can be only one character long")
		os.Exit(1)
	}
}

func main() {
	var taulu [106][17]byte
	bits := 106 * 17
	taulu[0][0] = 1
	luku := big.NewInt(0)
	_, success := luku.SetString(imgtext, 10)
	if !success {
		fmt.Println("Creating big int from text failed.")
		os.Exit(1)
	}
	luku.Div(luku, big.NewInt(17))
	for i := 0; i < 106; i++ {
		for j := 0; j < 17; j++ {
			bits--
			if luku.Bit(bits) == 1 {
				taulu[105-i][16-j] = fill
			} else {
				taulu[105-i][16-j] = empty
			}
		}
	}
	for i := 0; i < 17; i++ {
		for j := 105; j >= 0; j-- {
			fmt.Printf("%c", taulu[j][i])
		}
		fmt.Println()
	}
}
