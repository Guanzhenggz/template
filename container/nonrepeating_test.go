package main

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct{
		s string
		ans int
	} {
		//Normal cases
		{"abcaaaaaaa",3},
		{"pwwkew",3},

		// Edge cases
		{"",0},
		{"b",1},
		{"cccccc",1},
	}

	for _,tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("Got %d for input %s;" + "expected %d",actual,tt.s,tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "abcaaaaaaa"
	for i := 0; i < 10;i ++ {
		s = s + s
	}
	b.Logf("len(s)= %d",len(s))
	ans := 3
	b.ResetTimer()

	for i := 0; i < b.N ; i ++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s;"+"expected %d", actual, s, ans)
		}
	}

}
