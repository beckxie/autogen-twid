package generate

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// checksum rule: http://www2.lssh.tp.edu.tw/~hlf/class-1/lang-c/id/index.htm
var checksumMagic []int = []int{1, 9, 8, 7, 6, 5, 4, 3, 2, 1}

// prefix number
var prefixMappingNumber map[int]string = map[int]string{
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
	16: "G",
	17: "H",
	34: "I",
	18: "J",
	19: "K",
	20: "L",
	21: "M",
	22: "N",
	35: "O",
	23: "P",
	24: "Q",
	25: "R",
	26: "S",
	27: "T",
	28: "U",
	29: "V",
	32: "W",
	30: "X",
	31: "Y",
	33: "Z",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateID return one random Taiwan's ID.
func GenerateID() string {
	pre := generatePrefix()
	mid := generateMiddleNumber()
	suf, err := generateSuffix(pre, mid)
	if err != nil {
		log.Printf("generateSuffix is fail:%v\n", err)
		return ""
	}
	return prefixMappingNumber[pre] + strconv.Itoa(mid) + strconv.Itoa(suf)
}

// generatePrefix return one random alphabet.
func generatePrefix() int {
	// A~Z:10~35
	return rand.Intn(26) + 10
}

// generateMiddleNumber return one random number.
func generateMiddleNumber() int {
	return genGenderNumber()*10000000 + rand.Intn(9999999)
}

// genGenderNumber The first digit depends on gender;
// 1 for male, 2 for female.
// 8 for male(foreign), 9 for female(foreign).
func genGenderNumber() int {
	const typeCount = 4
	var genderNum = [typeCount]int{1, 2, 8, 9}
	return genderNum[rand.Intn(typeCount)]
}

func intToSliceInt(i int) (sliceInt []int) {
	for _, v := range strconv.Itoa(i) {
		j, _ := strconv.Atoi(string(v))
		sliceInt = append(sliceInt, j)
	}
	return sliceInt
}

// generateSuffix return one valid number.
func generateSuffix(pre int, mid int) (suf int, err error) {
	valid := intToSliceInt(pre)
	sMid := intToSliceInt(mid)
	allNumber := append(valid, sMid...)
	log.Printf("len(allNumber):%d\n", len(allNumber))
	log.Printf("allNumber:%#v\n", allNumber)

	if len(checksumMagic) != len(allNumber) {
		return suf, fmt.Errorf("generateSuffix is fail, because length is not match: checksumMagic:%d,allNumber:%d", checksumMagic, allNumber)
	}

	var total int
	// calculate checksum
	for i := 0; i < len(checksumMagic); i++ {
		total = total + checksumMagic[i]*allNumber[i]
		log.Printf("checksumMagic[%d]*allNumber[%d]=%d+%d=%d\n", i, i, checksumMagic[i], allNumber[i], checksumMagic[i]*allNumber[i])
	}
	log.Printf("total:%d\n", total)
	v := total % 10
	if v != 0 {
		suf = 10 - v
	} else {
		suf = v
	}

	return suf, nil
}
