package genpasses

type GenAllPermutations struct {
	allPasses []string
}

func Gen(passArr []string) []string {
	gap := GenAllPermutations{allPasses: make([]string, 0, genAllPermutationsCount(len(passArr)))}
	gap.generateAllPermutation(passArr, "")
	return gap.allPasses
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

func genAllPermutationsCount(countOfPasses int) int {
	count := 0
	for i := 1; i <= countOfPasses; i++ {
		count += factorial(countOfPasses) / factorial(countOfPasses-i)
	}
	return count
}

func factorial(n int) int {
	f := 1
	for j := 1; j <= n; j++ {
		f *= j
	}
	return f
}
