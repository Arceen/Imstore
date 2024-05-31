package encoder

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func EncodeToImage(input_file string, output_file string) {
	fmt.Println("Starting Encoder")
	file_in_bytes, err := os.ReadFile(input_file)
	if err != nil {
		panic(err)
	}
	img := imageFromBytes(file_in_bytes)

	out, _ := os.Create(output_file)
	defer out.Close()
	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Encoded Successfully.\nEncoded File Location -> ", output_file)
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

	// fmt.Println(i)

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
