package decoder

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func DecodeFromImage(input_file string, output_file string) {
	fmt.Println("Starting File Decode...")
	image_file, err := os.Open(input_file)

	if err != nil {
		panic(err)
	}
	image_png, err := png.Decode(image_file)
	if err != nil {
		panic(err)
	}

	img := bytesFromImage(image_png)
	err = os.WriteFile(output_file, img, 0666)
	if err != nil {
		panic(err)
	}
	// fmt.Println("OLD PNG DATA")
	// fmt.Println(len(img))
	// for _, item := range img {
	// 	fmt.Printf("%d ", item)
	// }

	// file_in_bytes, _ := os.ReadFile("f.txt")
	// for i := range file_in_bytes {
	// 	fmt.Printf("%d %d \n", file_in_bytes[i], img[i])
	// }
	fmt.Println("Successfully Decoded File -> " + output_file)
}

func bytesFromImage(img image.Image) []byte {
	bound := img.Bounds()
	data := make([]byte, (bound.Max.X-bound.Min.X)*(bound.Max.Y-bound.Min.Y)*4)
	// fmt.Println(len(data))
	i := 0
	// fmt.Println("RES:")
	// fmt.Printf("%d %d %d %d\n", bound.Min.X, bound.Max.X, bound.Min.Y, bound.Max.Y)
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			color := img.At(x, y).(color.NRGBA)
			data[i], data[i+1], data[i+2], data[i+3] = color.R, color.G, color.B, color.A
			i += 4
		}
	}
	return data
}
