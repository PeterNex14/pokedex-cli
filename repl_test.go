package main

import (
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	 string
		expected []string
	} {
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "we		lu tolol 	banget sih",
			expected: []string{"we","lu","tolol","banget","sih"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "PikAcHu CHArizard",
			expected: []string{"pikachu", "charizard"},
		},
	}
	
	passCount := 0
	failCount := 0

	

	for _, c := range cases {
		actual := cleanInput(c.input)
		wordMismatch := false

		if len(actual) != len(c.expected) {
			t.Errorf("Length of the actual and expected is different")
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			if expectedWord != word {
				wordMismatch = true
			} 

			
		}


		if wordMismatch {
			failCount++
			t.Errorf(`---------------------------------
Expected: 	%v
Actual: 	%v
Fail
`, c.expected, actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Expected: 	%v
Actual: 	%v
PASS
`, c.expected, actual)
		}

		
		
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	

}