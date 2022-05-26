package roman

import "strings"

type Converter struct {
}

func NewConverter() *Converter {
	return &Converter{}
}

func (r *Converter) convert(rmnNum byte) int {
	num := &number{
		romanNumber: romanNumber{
			value: rmnNum,
		},
	}
	return num.convert(num)
}

func (r *Converter) Convert(input string) int {
	result := 0
	inputLength := len(input)
	romanNumber := strings.ToUpper(input)

	for i := 0; i < inputLength; i++ {
		resI := r.convert(romanNumber[i])

		if i+1 < inputLength {
			resNxt := r.convert(romanNumber[i+1])

			if resI < resNxt {
				result += resNxt - resI
				i++
				continue
			}
		}

		result += resI
	}

	return result
}
