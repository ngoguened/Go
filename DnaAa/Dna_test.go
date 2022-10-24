package dna

import (
	"os"
	"testing"
)

func TestValidDNA(t *testing.T) {
	if !validNucleotide('a') {
		t.Errorf("Valid DNA marked invalid")
	}
	if validNucleotide('q') {
		t.Errorf("Invalid DNA marked valid")
	}
}

func TestValidAminoAcid(t *testing.T) {
	if !validAminoAcid('P') {
		t.Errorf("Valid AA marked invalid")
	}
	if validAminoAcid('B') {
		t.Errorf("Invalid AA marked valid")
	}
}

//Fix Translate Tests with file idiom.
func TestTranslate(t *testing.T) {

	os.Truncate("DNAinput.txt", 0)
	os.Truncate("DNAoutput.txt", 0)
	os.WriteFile("DNAinput.txt", []byte("ATGGAG"), 0666)
	translate("DNAinput.txt")
	output, _ := os.ReadFile("DNAoutput.txt")

	if output[0] != 'M' {
		t.Errorf("Incorrect DNA to Amino conversion")
	}
	if output[1] != 'E' {
		t.Errorf("Incorrect DNA to Amino conversion")
	}

	os.Truncate("DNAinput.txt", 0)
	os.Truncate("DNAoutput.txt", 0)
	os.WriteFile("DNAinput.txt", []byte("TAATGGAG"), 0666)
	translate("DNAinput.txt")
	output, _ = os.ReadFile("DNAoutput.txt")

	if output[0] != 'M' {
		t.Errorf("Incorrect DNA to Amino conversion")
	}
	if output[1] != 'E' {
		t.Errorf("Incorrect DNA to Amino conversion")
	}

	if len(output) > 2 {
		t.Errorf("Not stopping at STOP.")
	}

	os.Truncate("DNAinput.txt", 0)
	os.Truncate("DNAoutput.txt", 0)
	os.WriteFile("DNAinput.txt", []byte("ATGAAA"), 0666)
	translate("DNAinput.txt")
	output, _ = os.ReadFile("DNAoutput.txt")
	if len(output) == 0 {
		t.Errorf("STOP Identified out of alignment with Start")
	}

}

func TestSearchSequence(t *testing.T) {
	test, _ := searchSequence("ATGAAA", "Lys")
	if !test {
		t.Errorf("Did not find element when it should.")
	}
	test, _ = searchSequence("ATGAAA", "Gln")
	if test {
		t.Errorf("Found element when it should not have.")
	}

}
