package utils

import (
	"testing"
)

func TestCapitalizeFirstLetter(t *testing.T) {
	input := "There is no greater a agony than bearing a untold story inside you and a umbrella."
	expected := "There is no greater an agony than bearing an untold story inside you and an umbrella."

	result := LoopText(input)
	 if result != expected {
		t.Errorf("capitalizeFirstletter testing failed")
	 }
	if !checkForVowels("apple") {
		t.Errorf("word 'lapple' not a vowel")
	}
	if capitalizeFirstLetter("long") != "Long" {
		t.Errorf("Failed to convert")
	}
}



