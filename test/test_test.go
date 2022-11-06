package test

import (
	"fmt"
	"testing"

	"github.com/pashanskiy/7zFractalBruteforce/internal/genpasses"
)

func TestPermutation(t *testing.T) {
	ses := []string{"1", "2", "3", "4"}

	keks := Gen(ses)
	_ = keks
	for _, kek := range keks {
		fmt.Println(kek)
	}
}

func (ap *GenAllPermutations) generateAllPermutation(arr []string, spass string) {
	for i, pass := range arr {
		ap.allPasses = append(ap.allPasses, spass+pass)
		reArr := []string{}
		reArr = append(reArr, arr[:i]...)
		reArr = append(reArr, arr[i+1:]...)
		if len(reArr) == 0 {
			continue
		}
		ap.generateAllPermutation(reArr, spass+pass)
	}
}

type GenAllPermutations struct {
	allPasses []string
}

func Gen(passArr []string) []string {
	pcountOfPerm := 0
	for i := 1; i < len(passArr)+1; i++ {
		countOfPerm := 1
		for j := 1; j < i+1; j++ {
			countOfPerm *= j
		}
		pcountOfPerm += countOfPerm
	}

	gap := GenAllPermutations{allPasses: make([]string, 0, pcountOfPerm)}
	gap.generateAllPermutation(passArr, "")
	fmt.Println("all ", pcountOfPerm)

	return gap.allPasses
}

func TestGen(t *testing.T) {
	kek := []string{"1", "2", "3", "4", "5", "6", "7"}
	genpasses.Gen(kek)
}
