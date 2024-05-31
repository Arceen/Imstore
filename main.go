package main

import (
	"ahmannur.dev/imstore/internal/decoder"
	"ahmannur.dev/imstore/internal/encoder"
)

type sample_file struct {
	input   string
	encoded string
	decoded string
}

func main() {
	samples := []sample_file{
		{input: "sample/sample_image.jpg", encoded: "encoded/encoded_image.png", decoded: "decoded/decoded_image.jpg"},
		{input: "sample/sample_pdf.pdf", encoded: "encoded/encoded_pdf.png", decoded: "decoded/decoded_pdf.pdf"},
		{input: "sample/sample_video.mp4", encoded: "encoded/encoded_video.png", decoded: "decoded/decoded_video.mp4"},
	}
	for _, sample := range samples {
		encoder.EncodeToImage(sample.input, sample.encoded)
		decoder.DecodeFromImage(sample.encoded, sample.decoded)
	}
}
