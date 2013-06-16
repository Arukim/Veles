package main

import (
	"fmt"
	"os"
	"image"
	"image/jpeg"
	"image/draw"
)

func countFreq(rgba []uint8) []float64{
	res := make([]float64, 256)
	for i := range rgba{
//		if rgba[i] != 0{
//			fmt.Printf("Yohoho %v\n", i)
//		}
		res[rgba[i]]++
	}
	return res
}

func showFreq(freq []float64){
	for i := range freq{
		if freq[i] != 0{
			fmt.Printf("Freq of %v is %v\n",i,freq[i])
		}
	}
}

func main() {
	fImg1, err := os.Open("img/red.jpeg")
	if err != nil {
		fmt.Printf("Error opening file\n");
		return
	}
	defer fImg1.Close()
	img1, err := jpeg.Decode(fImg1)
	if err != nil {
		fmt.Printf("Error %v\n",err)
		return
	}
	fmt.Printf("Color is %v\n", img1.At(0,0))
	b := img1.Bounds()
	m := image.NewRGBA(image.Rect(0,0,b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), img1, b.Min, draw.Src)
	sum := 0
	for _ = range m.Pix{
		sum++
	}
	fmt.Printf("Pixels count is %v\n", sum)
	
	greyFreq := countFreq(m.Pix)
	showFreq(greyFreq)
}
