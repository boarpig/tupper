package main

import (
	_ "github.com/spakin/netpbm"
	"image"
	"os"
)

func main() {
	file, err := os.Open("testi.pbm")
	if err != nil {
		panic(err)
	}
	imgfile, str, err := image.Decode(file)
	bound := imgfile.Bounds()
	if str == "pbm" {
		if bound.Max.X == 106 && bound.Max.Y == 17 {
			for x := 0; x < 106; x = x + 1 {
				for y := 0; y < 17; y = y + 1 {
					b, _, _, _ := imgfile.At(x, y).RGBA()
					if b == 0 {
						print(0)
					} else {
						print(1)
					}
				}
				println()
			}
		} else {
			println("Image dimensions must be 106x17.")
			os.Exit(1)
		}
	} else {
		println("Image must be in pbm format.")
		os.Exit(2)
	}
}
