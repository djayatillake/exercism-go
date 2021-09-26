// package twelve has functions that outputs the verses and whole song of the twelve
// days of christmas
package twelve

import (
	"fmt"
	"strings"
)

var order = []string{
	"first", "second", "third", "fourth", "fifth", "sixth", "seventh",
	"eighth", "ninth", "tenth", "eleventh", "twelfth",
}

var items = []string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Verse(no int) (ret string) {
	ret += fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", order[no-1])
	for i := no - 1; i > -1; i-- {
		switch {
		case no > 1 && i == 0:
			ret += fmt.Sprintf(", and %s", items[i])
		case no == 1 || no-1 == i:
			ret += items[i]
		default:
			ret += fmt.Sprintf(", %s", items[i])
		}
	}
	return
}

func Song() (ret string) {
	for i := 1; i < 13; i++ {
		ret += Verse(i) + "\n"
	}
	ret = strings.TrimSpace(ret)
	return
}
