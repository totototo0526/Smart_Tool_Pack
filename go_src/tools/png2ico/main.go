package main

import (
	"encoding/binary"
	"fmt"
	"image/png"
	"io"
	"os"
)

// ICO Header
type ICOHeader struct {
	Reserved uint16
	Type     uint16 // 1 for Icon
	Count    uint16
}

// Directory Entry
type ICODirEntry struct {
	Width    uint8
	Height   uint8
	Colors   uint8
	Reserved uint8
	Planes   uint16
	BitCount uint16
	Size     uint32
	Offset   uint32
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: png2ico <input.png> <output.ico>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read PNG
	file, err := os.Open(inputFile)
	if err != nil {
		fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fatal(err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width > 256 || height > 256 {
		fmt.Println("Warning: Dimensions larger than 256x256 might not work well for standard ICO.")
	}

	// Read raw PNG data for embedding
	file.Seek(0, 0)
	pngData, err := io.ReadAll(file)
	if err != nil {
		fatal(err)
	}

	// Create ICO
	out, err := os.Create(outputFile)
	if err != nil {
		fatal(err)
	}
	defer out.Close()

	// Write Header
	header := ICOHeader{
		Reserved: 0,
		Type:     1,
		Count:    1,
	}
	binary.Write(out, binary.LittleEndian, header)

	// Write Directory Entry
	w := uint8(width)
	h := uint8(height)
	if width >= 256 {
		w = 0
	}
	if height >= 256 {
		h = 0
	}

	entry := ICODirEntry{
		Width:    w,
		Height:   h,
		Colors:   0,
		Reserved: 0,
		Planes:   1,
		BitCount: 32,
		Size:     uint32(len(pngData)),
		Offset:   6 + 16, // Header (6) + 1 DirEntry (16)
	}
	binary.Write(out, binary.LittleEndian, entry)

	// Write PNG Data
	out.Write(pngData)

	fmt.Printf("Converted %s to %s\n", inputFile, outputFile)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
