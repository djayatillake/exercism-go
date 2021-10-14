// package secret contains a function which implements a secret handshake
package secret

// Handshake returns a secret handshake based on a decimal input
func Handshake(input uint) (ret []string) {
	input %= 32
	actions := []string{"wink", "double blink", "close your eyes", "jump"}
	reverse := input/16 == 1
	input %= 16
	for i := 0; i < 4; i++ {
		if input&(1<<i) > 0 {
			if reverse {
				ret = append([]string{actions[i]}, ret...)
			} else {
				ret = append(ret, actions[i])
			}
		}
	}
	return
}
