package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	firmware "feicho.com/reverseengineering/suite/firmware"
)

func main() {
	startAddress := flag.String("startAddress", "0x00000000", "Start address (hex value, e.g. 0x00000000)")
	length := flag.String("length", "", "Length of section (hex value, e.g. 0x00010000)")
	chunkSize := flag.Int64("chunkSize", 4096, "Chunk size to be used")

	inputFilePath := flag.String("inputFilePath", "", "Full path to input file")
	outputDirectoryPath := flag.String("outputDirectoryPath", "", "Full path to output directory")
	outputFileName := flag.String("outputFileName", "", "Name of output file")

	flag.Parse()

	*outputFileName = strings.ReplaceAll(*outputFileName, "\\", "")

	if *startAddress == "" || !strings.HasPrefix(*startAddress, "0x") {
		fmt.Println("Invalid startAddress.")
		os.Exit(1)
	}

	if *length == "" || !strings.HasPrefix(*length, "0x") {
		fmt.Println("Invalid length.")
		os.Exit(1)
	}

	if *chunkSize == 0 {
		fmt.Println("Invalid chunkSize.")
		os.Exit(1)
	}

	if *inputFilePath == "" {
		fmt.Println("Invalid inputFile param.")
		os.Exit(1)
	}

	if *outputDirectoryPath == "" {
		fmt.Println("Invalid outputDirectory")
		os.Exit(1)
	}

	inputFileInfo, err := os.Stat(*inputFilePath)

	if err != nil {
		fmt.Println("Input file can't accessed.")
		os.Exit(1)
	}

	if *chunkSize > inputFileInfo.Size() {
		fmt.Println("chunkSize cannot be larger than file size.")
		os.Exit(1)
	}

	if *outputFileName == "" {
		fmt.Println("outputFileName cannot be empty.")
		os.Exit(1)
	}

	_, folderError := os.Stat(*outputDirectoryPath)
	if folderError != nil {
		fmt.Println("Output folder does not exist.")
		os.Exit(1)
	}

	*startAddress = strings.ReplaceAll(*startAddress, "0x", "")
	*length = strings.ReplaceAll(*length, "0x", "")

	firmware.SplitFile(*startAddress, *length, *chunkSize, *inputFilePath, strings.Join([]string{*outputDirectoryPath, *outputFileName}, "\\\\"))
}
