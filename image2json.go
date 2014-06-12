package main

import (
	//"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var (
	filename = flag.String("f", "", "Image file to decode")
	output   = flag.String("o", "", "file to output")
	nocolor  = flag.Bool("nc", false, "Output image without color info")
	res      []byte
)

type Pixel struct {
	X int
	Y int
	C color.Color
}

type Img struct {
	Width  int
	Height int
	Bytes  []Pixel
}

type Px [2]int
type ImgNoColor struct {
	Width  int
	Height int
	Bytes  []Px
}

func main() {

	white := color.RGBA{255, 255, 255, 255}

	flag.Parse()

	if *filename == "" {
		fmt.Println("Specify and image")
		return
	}

	reader, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Image %s not found\n", filename)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error reading image %s\n", err)
	}

	reader.Seek(0, 0)

	config, _, err := image.DecodeConfig(reader)
	if err != nil {
		fmt.Printf("config %s", err)
	}

	// When processing image without color, we are outputing an array, of
	// pixel positions, with the blank color out.
	if *nocolor {

		bounds := m.Bounds()
		iout := ImgNoColor{Width: config.Width, Height: config.Height}

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				ocolor := m.At(x, y)
				c := color.RGBAModel.Convert(ocolor)
				//c := Color{r, g, b, a}
				//fmt.Printf("c%v\n", c)
				if c == white {
					continue
				}
				p := Px{x, y}
				iout.Bytes = append(iout.Bytes, p)
			}
		}
		fmt.Printf("Result array length: %d\n", len(iout.Bytes))
		res, _ = json.Marshal(iout)

	} else {

		bounds := m.Bounds()

		iout := Img{Width: config.Width,
			Height: config.Height}

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				ocolor := m.At(x, y)
				c := color.RGBAModel.Convert(ocolor)
				//c := Color{r, g, b, a}
				//fmt.Printf("c%v\n", c)
				if c == white {
					continue
				}

				p := Pixel{X: x, Y: y, C: c}
				iout.Bytes = append(iout.Bytes, p)
			}
		}
		fmt.Printf("Result array length: %d\n", len(iout.Bytes))
		res, _ = json.Marshal(iout)

	}

	//fmt.Printf("%v", iout)

	//os.Stdout.Write(res)
	if *output == "" {
		fmt.Println("Needs a output file")
	}

	ff, err := os.Create(*output)
	defer ff.Close()
	ff.Write(res)

	fmt.Printf("Processed image %s, width: %d height: %d\n", *filename, config.Width, config.Height)

}
