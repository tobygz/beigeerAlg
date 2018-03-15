package main

import (
	"fmt"
)

func main() {
	parseAry([]int{1, 2, 3, 4, 5, 6, 7})
	//parseAry([]int{1, 2, 3, 4, 5, 6, 7, 8})
}

func revSlc(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func parseAry(srcAry []int) [][][]int {
	idx := 0
	retSlice := make([][]int, 0)
	aryLen := len(srcAry)
	initLen := aryLen - 1
	if len(srcAry)%2 != 0 {
		srcAry = append(srcAry, 0)
		initLen = aryLen
	}
	idx++

	step := 1
	if aryLen > 4 {
		step = (aryLen-4)/2 + 1
	}

	max := srcAry[len(srcAry)-1]
	//shift
	tmpAry := make([]int, len(srcAry))
	copy(tmpAry, srcAry)
	for n := 0; n < initLen; n++ {

		isMax := false
		if tmpAry[0] == max {
			isMax = true
		}

		rmIdx := len(tmpAry) - 1
		if isMax {
			rmIdx = 0
		}
		//remove empty
		tmpAry = append(tmpAry[:rmIdx], tmpAry[rmIdx+1:]...)

		tmpAryLen := len(tmpAry)
		stepLen := step
		if isMax {
			stepLen = step + 2
		}
		for i := 0; i < stepLen; i++ {
			temp := tmpAry[tmpAryLen-1]
			for j := tmpAryLen - 2; j >= 0; j-- {
				tmpAry[j+1] = tmpAry[j]
			}
			tmpAry[0] = temp
		}

		//bukong
		if isMax {
			tmpAry = append(tmpAry, max)
		} else {
			tmpAry = append([]int{max}, tmpAry[0:]...)
		}
		//fmt.Println(tmpAry)
		dstAry := make([]int, len(srcAry))
		copy(dstAry, tmpAry)
		retSlice = append(retSlice, dstAry)
	}

	//get pairs
	pairSlc := make([][][]int, 0)
	for _, slc := range retSlice {
		roundSlc := make([][]int, 0)
		nowLen := len(slc) / 2
		left := slc[:nowLen]
		right := revSlc(slc[nowLen:])
		for eidx, _ := range left {
			roundSlc = append(roundSlc, []int{left[eidx], right[eidx]})
		}
		pairSlc = append(pairSlc, roundSlc)
	}

	return pairSlc
}
