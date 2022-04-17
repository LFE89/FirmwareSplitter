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

package splitter

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func SplitFile(startAddressHex string, lengthHex string, chunkSize int64, inputFile string, outputFile string) error {

	fmt.Println("[1/4] Process started...Please wait.")

	startAddress, _ := strconv.ParseInt(startAddressHex, 16, 64)
	length, _ := strconv.ParseInt(lengthHex, 16, 64)

	inputStream, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	outputStream, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	fmt.Println("[2/4] Creating splitted file: " + outputFile)

	var bufferSize int64 = chunkSize
	var isNextRunCompleted bool = false
	var totalBytesWritten int64 = 0
	var skip int64 = 0

	fmt.Println("[3/4] Writing to file...")

	defer func() {
		outputStream.Close()
		inputStream.Close()
	}()

	for {
		buffer := make([]byte, bufferSize)
		readBytes, err := inputStream.ReadAt(buffer, (startAddress + skip))
		if err != nil {
			return err
		}

		if err != nil && err != io.EOF {
			return err
		}

		if _, err := outputStream.Write(buffer[:]); err != nil {
			return err
		}

		if isNextRunCompleted {
			break
		}

		var writtenBytesInRun int64 = int64(readBytes)
		totalBytesWritten += writtenBytesInRun

		if (totalBytesWritten + writtenBytesInRun) >= length {
			newSkip := (startAddress + skip) - length
			skip = skip - newSkip
			bufferSize = length - totalBytesWritten
			isNextRunCompleted = true
		} else {
			skip = skip + writtenBytesInRun
		}
	}

	fmt.Println("[4/4] Process completed.")
	return nil
}
