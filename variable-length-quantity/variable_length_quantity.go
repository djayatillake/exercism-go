// package variablelengthquantity implements variable length encoding and decoding
package variablelengthquantity

import "fmt"

// EncodeVarint encodes a uint into []byte using variable length encoding
func EncodeVarint(input []uint32) (ret []byte) {
	for _, no := range input {
		if no == 0 {
			ret = append(ret, byte(0))
		}
		ret_i := []byte{}
		for ; no > 0; no /= 128 {
			byt := no % 128
			if no/128 == 0 {
				byt = no
			}
			if len(ret_i) > 0 {
				byt = byt | 1<<7
			}
			ret_i = append([]byte{byte(byt)}, ret_i...)
		}
		ret = append(ret, ret_i...)
	}
	return
}

// DecodeVarint decodes a variable length encoded number as []byte into a uint
func DecodeVarint(input []byte) (ret []uint32, err error) {
	sum := uint32(0)
	incomplete := true
	for _, v := range input {
		if uint32(v)&(1<<7) == 0 {
			sum += uint32(v)
			ret = append(ret, sum)
			sum = uint32(0)
			incomplete = false
		} else {
			sum = sum << 7
			sum += (uint32(v) - 128) << 7
			incomplete = true
		}
	}
	if incomplete {
		err = fmt.Errorf("incomplete sequence")
	}
	return
}
