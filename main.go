package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	convertFileToImage("file", ".pdf")
	convertImageToPDFFile("file.png", ".pdf")
}
func convertFileToImage(file_name string, extension string) {

	file_in_bytes, err := os.ReadFile(file_name + extension)
	if err != nil {
		panic("Error reading file")
	}
	img := imageFromBytes(file_in_bytes)

	out, _ := os.Create(file_name + ".png")
	defer out.Close()
	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Encoded")
}

func imageFromBytes(data []byte) *image.NRGBA {
	delimiterKey := []byte{0x35, 0x36, 0x37, 0x38}
	delimiterIndex := 0
	// fmt.Println("Original Data")
	// for _, d := range data {
	// 	fmt.Println(d)
	// }
	width := 1024
	height := (len(data)/4 + width + 1 + len(delimiterKey)) / (width)
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	i := 0
	for ; i+3 < len(data); i += 4 {
		x := (i / 4) % (width)
		y := (i / 4) / (width)
		img.Set(x, y, color.NRGBA{R: data[i], G: data[i+1], B: data[i+2], A: data[i+3]})
	}

	fmt.Println(i)

	if i < len(data) {
		x := (i / 4) % width
		y := (i / 4) / width
		var final_colors [4]byte
		final_colors[0] = data[i]
		if i < len(data)-1 {
			final_colors[1] = delimiterKey[delimiterIndex]
			delimiterIndex++
		} else {
			final_colors[1] = data[i+1]
		}

		if i < len(data)-2 {
			final_colors[2] = delimiterKey[delimiterIndex]
			delimiterIndex++
		} else {
			final_colors[2] = data[i+2]
		}
		final_colors[3] = delimiterKey[delimiterIndex]
		delimiterIndex++
		img.Set(x, y, color.NRGBA{R: final_colors[0], G: final_colors[1], B: final_colors[2], A: final_colors[3]})
		i += 4
	}
	l := len(delimiterKey)
	for (i / 4) < height*width {
		x := (i / 4) % width
		y := (i / 4) / width
		img.Set(x, y, color.NRGBA{R: delimiterKey[delimiterIndex%l], G: delimiterKey[(delimiterIndex+1)%l], B: delimiterKey[(delimiterIndex+2)%l], A: delimiterKey[(delimiterIndex+3)%l]})
		i += 4
	}
	// newData := bytesFromImage(img)
	// fmt.Println("NEW DATA")
	// for _, d := range newData {
	// 	fmt.Printf("%d ", d)
	// }
	return img
}

func convertImageToPDFFile(file_name string, extension string) {
	image_file, err := os.Open(file_name)

	if err != nil {
		panic(err)
	}
	image_png, err := png.Decode(image_file)
	if err != nil {
		panic(err)
	}

	img := bytesFromImage(image_png)
	err = os.WriteFile("decoded_"+file_name+extension, img, 0666)
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
	fmt.Println("Successfully Decoded file")
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
