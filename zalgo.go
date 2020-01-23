package gostringconverters

import (
	"math/rand"
	"strings"
)

var (
	zalgoCharList = []rune{
		'\u0300',
		'\u0301',
		'\u0302',
		'\u0303',
		'\u0304',
		'\u0305',
		'\u0306',
		'\u0307',
		'\u0308',
		'\u0309',
		'\u030A',
		'\u030B',
		'\u030C',
		'\u030D',
		'\u030E',
		'\u030F',
		'\u0310',
		'\u0311',
		'\u0312',
		'\u0313',
		'\u0314',
		'\u0315',
		'\u031A',
		'\u031B',
		'\u033D',
		'\u033E',
		'\u033F',
		'\u0340',
		'\u0341',
		'\u0342',
		'\u0343',
		'\u0344',
		'\u0346',
		'\u034A',
		'\u034B',
		'\u034C',
		'\u0350',
		'\u0351',
		'\u0352',
		'\u0357',
		'\u0358',
		'\u035B',
		'\u035D',
		'\u035E',
		'\u0360',
		'\u0361',
		'\u0316',
		'\u0317',
		'\u0318',
		'\u0319',
		'\u031C',
		'\u031D',
		'\u031E',
		'\u031F',
		'\u0320',
		'\u0321',
		'\u0322',
		'\u0323',
		'\u0324',
		'\u0325',
		'\u0326',
		'\u0327',
		'\u0328',
		'\u0329',
		'\u032A',
		'\u032B',
		'\u032C',
		'\u032D',
		'\u032E',
		'\u032F',
		'\u0330',
		'\u0331',
		'\u0332',
		'\u0333',
		'\u0339',
		'\u033A',
		'\u033B',
		'\u033C',
		'\u0345',
		'\u0347',
		'\u0348',
		'\u0349',
		'\u034D',
		'\u034E',
		'\u0353',
		'\u0354',
		'\u0355',
		'\u0356',
		'\u0359',
		'\u035A',
		'\u035C',
		'\u035F',
		'\u0362',
		'\u0334',
		'\u0335',
		'\u0336',
		'\u0337',
		'\u0338',
		'\u0363',
		'\u0364',
		'\u0365',
		'\u0366',
		'\u0367',
		'\u0368',
		'\u0369',
		'\u036A',
		'\u036B',
		'\u036C',
		'\u036D',
		'\u036E',
		'\u036F',
	}

	numZalgoChars = len(zalgoCharList)

	zalgoCharMap = makeZalgoCharMap(zalgoCharList)
)

func makeZalgoCharMap(l []rune) map[rune]bool {
	m := make(map[rune]bool, len(l))
	for _, char := range l {
		m[char] = true
	}
	return m
}

// ZalgoEncoder - Will convert a string to Zalgo format
// with numChars number of characters per rune
func ZalgoEncoder(s string, numChars int) string {
	var builder strings.Builder
	for _, r := range s {
		builder.WriteRune(r)
		for i := 0; i < numChars; i++ {
			builder.WriteRune(zalgoCharList[rand.Intn(numZalgoChars)])
		}
	}
	return builder.String()
}

// ZalgoDecoder - Will convert a Zalgo string to regular
// not very smart. It will get rid of all diacritical marks
func ZalgoDecoder(s string) string {
	var builder strings.Builder
	for _, r := range s {
		_, found := zalgoCharMap[r]
		if !found {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
