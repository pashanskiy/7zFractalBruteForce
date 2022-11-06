package unpack

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/pashanskiy/7zFractalBruteforce/internal/support"
)

type BruteGoriuteneConf struct {
	WG        *sync.WaitGroup
	CTX       context.Context
	Extract   bool
	PassChan  chan support.FilePass
	DoneChan  chan struct{}
	Thread    int
	Passes    []string
	Dir, File string
}

func BruteGoriutine(c BruteGoriuteneConf) {
	defer c.WG.Done()
	c.WG.Add(1)
	fileName := c.File[:len(c.File)-len(filepath.Ext(c.File))]
	outDIR := filepath.Join(c.Dir, support.TMPDir, strconv.Itoa(c.Thread), fileName)
	os.MkdirAll(outDIR, os.ModePerm)

	for index, pass := range c.Passes {
		select {
		case <-c.CTX.Done():
			fmt.Printf("\033[%d;%dH\033[2KKILLED File: %s Thread: %d", c.Thread, 0, c.File, c.Thread)
			os.RemoveAll(outDIR)
			return
		default:
			fmt.Printf("\033[%d;%dHBRUTE: %d/%d File: %s Thread: %d", c.Thread, 0, index+1, len(c.Passes), c.File, c.Thread)
			cmd := exec.Command("7z")
			if c.Extract {
				cmd.Args = []string{"", "x", filepath.Join(c.Dir, c.File), "-p" + pass, "-o" + outDIR, "-aoa"}
			} else {
				cmd.Args = []string{"", "t", filepath.Join(c.Dir, c.File), "-p" + pass, "-o" + outDIR, "-aoa"}
			}

			if cmd.Run() == nil {
				fmt.Printf("\033[%d;%dH\033[2KSUCCESS! File: %s Thread: %d", c.Thread, 0, c.File, c.Thread)
				c.PassChan <- support.FilePass{Pass: pass, Bruted: true}
				os.Rename(outDIR, filepath.Join(c.Dir, fileName))
				return
			}
		}
	}
	fmt.Printf("\033[%d;%dH\033[2KFAILED! File: %s Thread: %d", c.Thread, 0, c.File, c.Thread)
	c.DoneChan <- struct{}{}
	os.RemoveAll(outDIR)
	return
}
