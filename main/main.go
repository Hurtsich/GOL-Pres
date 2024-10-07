package main

import (
	"GOL-Pres/organism"
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
)

func main() {
	billy := organism.NewOrganism(100)

	err := video(billy)
	if err != nil {
		fmt.Printf("%w", err)
	}
}

func video(o organism.Organism) error {
	var images []*image.Paletted
	var delays []int

	fmt.Println("Action")

	for i := 0; i < 150; i++ {
		img := o.Move()
		fmt.Printf("Cliché n°%v", i)
		images = append(images, img)
		delays = append(delays, 7)
		o.Breathe()
		fmt.Println()
	}

	file, err := os.Create("../data/billy.gif")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	return gif.EncodeAll(writer, &gif.GIF{Image: images, Delay: delays})
}

func photo(img *image.Paletted) error {
	file, err := os.Create("../data/billy.png")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	return png.Encode(writer, img)
}
