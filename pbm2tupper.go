package main

import (
	_ "github.com/spakin/netpbm"
	"image"
	"os"
	"math/big"
)

func calculate_tupper(imgfile string) {
	file, err := os.Open(imgfile)
	if err != nil {
		println("Image file doesn't exist.")
		os.Exit(3)
	}
	tupper_img, str, err := image.Decode(file)
	tupper := big.NewInt(0)
	if str == "pbm" {
		bound := tupper_img.Bounds()
		position := 1802
		if bound.Max.X == 106 && bound.Max.Y == 17 {
			for x := 0; x < 106; x = x + 1 {
				for y := 16; y >= 0; y = y - 1 {
					b, _, _, _ := tupper_img.At(x, y).RGBA()
					position--
					if b == 0 {
						tupper.SetBit(tupper, position, 1)
					} else {
						tupper.SetBit(tupper, position, 0)
					}
				}
			}
		} else {
			println("Image dimensions must be 106x17.")
			os.Exit(1)
		}
	} else {
		println("Image must be in pbm format.")
		os.Exit(2)
	}
	tupper.Mul(tupper, big.NewInt(17))
	println(tupper.String())
}

func main() {
	calculate_tupper("testi.pbm")
}
