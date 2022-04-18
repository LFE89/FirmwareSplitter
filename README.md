# Firmware Dump Splitter
## License
This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

See the GNU General Public License for more details.
For full license text, see http://www.gnu.org/licenses.

## Release
| File        | Hash (SHA256) | Platform  |
|:-----:|:-----:|:-----:|
| release/fwsplitter-windows-x64.7z/fwsplitter.exe      | FD8A5F732E32951324D0120E8D21E50A6DCE57222F239A92B6886A36C5649C45 | Windows x64 |

## General
FirmwareSplitter is a fast and reliable software to extract different sections / partitions from firmware images.
![image](https://user-images.githubusercontent.com/60022287/163734878-9a0011c3-a1fa-4262-9b84-e810ff164e63.png)

It also supports custom chunkSizes, to handle large files such as 16GB full firmware dumps.
The software has been developed, to extract different paritions from Android devices directly on the device itself (.. or anywhere else). 
Has been used to mainly analyze a MediaTek Android OS device with a MT8167 chipset (with scatter files) - anyhow, the software can be used for any kind of files.

## Usage
### Help
```
fwsplitter -h

Usage of fwsplitter:
  -chunkSize int
        Chunk size to be used (default 4096)
  -inputFilePath string
        Full path to input file
  -length string
        Length of section (hex value, e.g. 0x00010000)
  -outputDirectoryPath string
        Full path to output directory
  -outputFileName string
        Name of output file
  -startAddress string
        Start address (hex value, e.g. 0x00000000) (default "0x00000000")
```

### Extract Partition From Dump - Example
```
fwsplitter 
  -chunksize 4096 
  -length "0x56000000" 
  -startAddress "0x08000000" 
  -inputFilePath "Z:\\Firmware\\dump.img" 
  -outputFileName "boot.img" 
  -outputDirectoryPath "Z:\\Firmware"
```
