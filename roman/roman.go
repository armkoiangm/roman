package roman

var romanToArabicNumbers = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

type iRomanNumberConverter interface {
	convertToNumber() int
	convert(iRomanNumberConverter) int
}

type romanNumber struct {
	value byte
}

func (rn *romanNumber) convert(i iRomanNumberConverter) int {
	return i.convertToNumber()
}

type number struct {
	romanNumber
}

func (ran *number) convertToNumber() int {
	return romanToArabicNumbers[ran.romanNumber.value]
}
