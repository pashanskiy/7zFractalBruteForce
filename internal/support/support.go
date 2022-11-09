package support

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const TMPDir = ".7zFractalBruteforce"

var EXTMap = map[string]struct{}{
	"7Z":       {},
	"XZ":       {},
	"BZIP2":    {},
	"GZIP":     {},
	"TAR":      {},
	"ZIP":      {},
	"WIM":      {},
	"APFS":     {},
	"AR":       {},
	"ARJ":      {},
	"CAB":      {},
	"CHM":      {},
	"CPIO":     {},
	"CRAMFS":   {},
	"DMG":      {},
	"EXT":      {},
	"FAT":      {},
	"GPT":      {},
	"HFS":      {},
	"IHEX":     {},
	"ISO":      {},
	"LZH":      {},
	"LZMA":     {},
	"MBR":      {},
	"MSI":      {},
	"NSIS":     {},
	"NTFS":     {},
	"QCOW2":    {},
	"RAR":      {},
	"RPM":      {},
	"SQUASHFS": {},
	"UDF":      {},
	"UEFI":     {},
	"VDI":      {},
	"VHD":      {},
	"VHDX":     {},
	"VMDK":     {},
	"XAR":      {},
	"Z":        {},
}

type FilePass struct {
	Pass   string
	Bruted bool
}

func PrintBrutedInfo(files []fs.FileInfo, brutedPassesMap map[string]FilePass, threads int) {
	fmt.Printf("\033[%d;%dH------------------------", threads+2, 0)
	for index, file := range files {
		if bpm, ok := brutedPassesMap[file.Name()]; ok {
			if bpm.Bruted {
				fmt.Printf("\033[%d;%dHID: %d, \033[32mSUCCESS: \033[35m%s\033[37m, File: %s,", threads+3+index, 0, index+1, bpm.Pass, file.Name())
			} else {
				fmt.Printf("\033[%d;%dH\033[2KID: %d, \033[31mFAILED\033[37m, File: %s", threads+3+index, 0, index+1, file.Name())
			}
		} else {
			fmt.Printf("\033[%d;%dHID: %d, \033[33mWAITING\033[37m, File: %s", threads+3+index, 0, index+1, file.Name())
		}
	}
}

func ScanPasses() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Passs")
	passArr := []string{}
	for {
		textByte, _, _ := reader.ReadLine()
		text := string(textByte)
		if len(text) != 0 {
			passArr = append(passArr, text)
		} else {
			break
		}
	}
	fmt.Print("\033c")
	return passArr
}

func GetFiles(readDir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(readDir)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(files); i++ {
		ext := filepath.Ext(files[i].Name())
		if len(ext) == 0 {
			files = append(files[:i], files[i+1:]...)
			i--
			continue
		}

		if _, ok := EXTMap[strings.ToUpper(ext[1:])]; !ok {
			files = append(files[:i], files[i+1:]...)
			i--
		}
	}

	fmt.Println("Files to 7z brute:")
	for index, file := range files {
		fmt.Println(index+1, file.Name(), file.Size())
	}
	return files
}