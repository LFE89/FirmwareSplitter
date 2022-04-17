# Firmware Dump Splitter
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
