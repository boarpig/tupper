package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
)

var fill, empty byte
var imgtext string = "4858450636189713423582095962494202044581400587983244549483093085061934704708809928450644769865524364849997247024915119110411605739177407856919754326571855442057210445735883681829823754139634338225199452191651284348332905131193199953502413758765239264874613394906870130562295813219481113685339535565290850023875092856892694555974281546386510730049106723058933586052544096664351265349363643957125565695936815184334857605266940161251266951421550539554519153785457525756590740540157929001765967965480064427829131488548259914721248506352686630476300"

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
				taulu[i][j] = fill
			} else {
				taulu[i][j] = empty
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
