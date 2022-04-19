# Firmware Dump Splitter / Section & Partition Extractor
## License
This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

See the GNU General Public License for more details.
For full license text, see http://www.gnu.org/licenses.

## Release
| File        | Hash (SHA256) | Platform  |
|:-----:|:-----:|:-----:|
| fwsplitter-v1.0.2-windows-x64.exe      | 46CEFBEFCFB353D83436072C520B1C65110B02B1A8AFED4C4D9AE1360707BBD4 | Windows x64 |
| fwsplitter-v1.0.2-linux-x64      | 7572FAF560F6AF1E6E379127D7131BC49EDE0B58EEF1208E7FA56147C86E8E5B | Linux x64 |
| fwsplitter-v1.0.2-linux-x86      | 9C0EA2C6C9EBA06983D4716EFCBFA8A733768E726C85D56DE22E3E2945C00075 | Linux x86 |
| fwsplitter-v1.0.2-linux-arm-v7      | 46CEFBEFCFB353D83436072C520B1C65110B02B1A8AFED4C4D9AE1360707BBD4 | Linux armv7 |
| fwsplitter-v1.0.2-macos-x64      | 649FE6C0BA0652B70215743F850EF841E1AF3D11CB852A23C5F377F8387577BA | MacOS x64 |

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
