package base62

import (
	"math"
	"strings"
)

const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const length = 62

func Encode(num int) string {
	r := num % length
	res := string(base[r])
	div := num / length
	q := int(math.Floor(float64(div)))

	for q != 0 {
		r = q % length
		temp := q / length
		q = int(math.Floor(float64(temp)))
		res = string(base[r]) + res
	}
	return res
}

func Decode(str string) int {
	var res int

	for _, r := range str {
		res = (length * res) + strings.Index(base, string(r))
	}
	return res
}
