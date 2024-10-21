package main

import "testing"

//func TestFields(t *testing.T) {
//	if Fields("hello noor") != ([]string{"hello", "noor"} ){
//		t.Error(`Fields("hello noor") = false`)
//	}
//}

func TestHex(t *testing.T) {
	if Hex("1E") != 30 {
		t.Error(`Hex("1E") = false`)
	}
	if Hex("2B") != 43 {
		t.Error(`Hex("2B") = false`)
	}
}

func TestBin(t *testing.T) {
	if Bin("101011") != 43 {
		t.Error(`Hex("1E") = false`)
	}
	if Bin("10") != 2 {
		t.Error(`Hex("2B") = false`)
	}
}

func TestConvertString(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},

		{"There it was. A amazing rock!", "There it was. An amazing rock!"}, // non-palindrome
		{"Hey (low)", "hey"}, // semi-palindrome
		//{"baba",true},
	}

	for _, test := range tests {
		if got := ConvertString(test.input); got != test.output {
			t.Errorf("ConvertString(%q) = %v", test.input, got)
		}
	}
}
