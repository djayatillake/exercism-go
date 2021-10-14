// package secret contains a function which implements a secret handshake
package secret

// Handshake returns a secret handshake based on a decimal input
func Handshake(input uint) (ret []string) {
	input %= 32
	actions := []string{"jump", "close your eyes", "double blink", "wink"}
	reverse := input/16 == 1
	input %= 16
	bools := make([]bool, 4)
	i := 0
	for j := 8; j > 0; j /= 2 {
		bools[i] = input/uint(j) == 1
		input %= uint(j)
		i++
	}
	for i := 0; i < 4; i++ {
		if bools[i] {
			if reverse {
				ret = append(ret, actions[i])
			} else {
				ret = append([]string{actions[i]}, ret...)
			}
		}
	}
	return
}
