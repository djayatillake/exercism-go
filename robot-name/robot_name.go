// package robotname generates a random robotname
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot stores the name of a robot
type Robot struct {
	name string
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// declaring namestore
var namestore = map[string]bool{}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// randomname generates a random name in the right format
func randomname() string {
	return RandStringBytes(2) + fmt.Sprintf("%03d", rand.Intn(999))
}

// Name name a robot
func (r *Robot) Name() (string, error) {
	for {
		if r.name != "" {
			break
		}
		name := randomname()
		if namestore[name] {
			continue
		}
		r.name = name
		namestore[name] = true
	}
	return r.name, nil
}

func (r *Robot) Reset() error {
	r.name = ""
	r.Name()
	return nil
}
