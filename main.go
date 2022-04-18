// Copyright (C) 2022 Lars D. Feicho
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the
// GNU General Public License for more details.
// You should have received a copy of the GNU General Public License
// along with this program.If not, see<http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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

	*outputFileName = strings.ReplaceAll(*outputFileName, "\\", "")
	_, folderError := os.Stat(*outputDirectoryPath)
	if folderError != nil {
		fmt.Println("Output folder does not exist.")
		os.Exit(1)
	}

	*startAddress = strings.ReplaceAll(*startAddress, "0x", "")
	*length = strings.ReplaceAll(*length, "0x", "")

	startAddressHex, _ := strconv.ParseInt(*startAddress, 16, 64)
	lengthHex, _ := strconv.ParseInt(*length, 16, 64)

	if inputFileInfo.Size() < (startAddressHex + lengthHex) {
		fmt.Println("File size cannot be smaller than startAddress + length.")
		os.Exit(1)
	}

	firmware.SplitFile(startAddressHex, lengthHex, *chunkSize, *inputFilePath, strings.Join([]string{*outputDirectoryPath, *outputFileName}, "\\\\"))
}
