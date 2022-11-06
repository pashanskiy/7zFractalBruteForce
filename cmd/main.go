package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/pashanskiy/7zFractalBruteforce/internal/brute"
	"github.com/pashanskiy/7zFractalBruteforce/internal/genpasses"
	"github.com/pashanskiy/7zFractalBruteforce/internal/support"
)

func main() {

	bruteDir := flag.String("dir", "./", "Path to the folder for bruteforce files")
	threads := flag.Int("threads", 11, "Threads to brute")
	extract := flag.Bool("extract", false, "Extract archive? [true|false]")
	flag.Parse()

	files := support.GetFiles(*bruteDir)
	passArr := genpasses.Gen(support.ScanPasses())

	fmt.Printf("\033[%d;%dHCount of passes: %d Extract: %v", *threads+1, 0, len(passArr), *extract)

	brute.BrutePasses(*bruteDir, *threads, *extract, passArr, files)

	fmt.Printf("\033[%d;%dHPress any key to exit...", *threads+4+len(files), 0)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Print("\033c")
}
