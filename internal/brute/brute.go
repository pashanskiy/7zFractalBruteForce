package brute

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/pashanskiy/7zFractalBruteforce/internal/support"
	"github.com/pashanskiy/7zFractalBruteforce/internal/unpack"
)

func BrutePasses(bruteDir string, threads int, extract bool, passArr []string, files []fs.FileInfo) {

	os.MkdirAll(filepath.Join(bruteDir, support.TMPDir), os.ModePerm)

	wg := sync.WaitGroup{}
	batch := len(passArr) / threads
	brutedPassesMap := make(map[string]support.FilePass, len(files))
	for _, file := range files {
		support.PrintBrutedInfo(files, brutedPassesMap, threads)
		ctx, cancel := context.WithCancel(context.Background())
		passChan := make(chan (support.FilePass), 0)
		doneChan := make(chan (struct{}), threads)

		for i := 1; i < threads; i++ {
			go unpack.BruteGoriutine(unpack.BruteGoriuteneConf{
				WG:       &wg,
				CTX:      ctx,
				Extract:  extract,
				PassChan: passChan,
				DoneChan: doneChan,
				Thread:   i,
				Passes:   passArr[(i-1)*batch : i*batch],
				Dir:      bruteDir,
				File:     file.Name(),
			})
		}
		go unpack.BruteGoriutine(unpack.BruteGoriuteneConf{
			WG:       &wg,
			CTX:      ctx,
			Extract:  extract,
			PassChan: passChan,
			DoneChan: doneChan,
			Thread:   threads,
			Passes:   passArr[(threads-1)*batch:],
			Dir:      bruteDir,
			File:     file.Name(),
		})
		doneCounter := 0
	F:
		for {
			select {
			case brutedPass := <-passChan:
				cancel()
				close(passChan)
				brutedPassesMap[file.Name()] = brutedPass
				break F

			case <-doneChan:
				doneCounter++
				if doneCounter >= threads {
					close(passChan)
					brutedPassesMap[file.Name()] = support.FilePass{
						Pass:   "",
						Bruted: false,
					}
					break F
				}
			}
		}
		cancel()
		wg.Wait()
		os.RemoveAll(filepath.Join(bruteDir, support.TMPDir))
	}
	support.PrintBrutedInfo(files, brutedPassesMap, threads)
}
