// package encode has functions which allow a string to be encoded and decoded
package encode

import (
	"fmt"
	"math"
	"strings"
)

// RunLengthEncode encodes a string into a compressed format
func RunLengthEncode(unencoded string) (encoded string) {
	count := 1
	var last_r rune
	for i, r := range unencoded {
		if i == 0 {
			last_r = r
		} else if last_r == r {
			count++
			if i == len(unencoded)-1 {
				encoded += fmt.Sprintf("%d%s", count, string(last_r))
			}
		} else if last_r != r && count == 1 {
			encoded += string(last_r)
			if i == len(unencoded)-1 {
				encoded += string(r)
			}
			last_r = r
		} else if last_r != r && count > 1 {
			encoded += fmt.Sprintf("%d%s", count, string(last_r))
			if i == len(unencoded)-1 {
				encoded += string(r)
			}
			last_r = r
			count = 1
		}
	}
	return
}

// RunLengthDecode decodes a string from a compressed format
func RunLengthDecode(encoded string) (decoded string) {
	is_num := false
	num_s := []rune{}
	num := 1
	for _, r := range encoded {
		is_num = r-'0' >= 0 && r-'0' < 10
		if !is_num {
			decoded += strings.Repeat(string(r), num)
			num = 1
			num_s = []rune{}
		} else {
			num_s = append([]rune{r}, num_s...)
			num = rune_s_to_int(num_s)
			continue
		}
	}
	return
}

func rune_s_to_int(rune_s []rune) (num int) {
	for i, v := range rune_s {
		num += int((v - '0')) * int(math.Pow10(i))
	}
	return
}
