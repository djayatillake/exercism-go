package atbash

import "testing"

func TestAtbash(t *testing.T) {
	for _, test := range tests {
		actual := Atbash(test.s)
		if actual != test.expected {
			t.Errorf("Atbash(%s): expected %s, actual %s, len_exp %d, len_actual %d",
				test.s, test.expected, actual, len(test.expected), len(actual))
		}
	}
}

func BenchmarkAtbash(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Atbash(test.s)
		}
	}
}
