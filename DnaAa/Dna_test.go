package dna

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
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

	test, _ := translate("ATGGAG")
	output, _ := os.ReadFile("DNAoutput.txt")
	fmt.Println(output)
	fmt.Println(test)

	if output[0] != 'M' {
		t.Errorf("Incorrect DNA to Amino conversion")
	}
	if output[1] != 'E' {
		t.Errorf("Incorrect DNA to Amino conversion")
	}

	test, _ = translate("TAATGGAG")
	output, _ = os.ReadFile("DNAoutput.txt")
	fmt.Println(output)
	fmt.Println(test)

	/*
		if test[0] != 'M' {
			t.Errorf("Incorrect DNA to Amino conversion")
		}
		if test[1] != 'E' {
			t.Errorf("Incorrect DNA to Amino conversion")
		}

		if len(test) > 2 {
			t.Errorf("Not stopping at STOP.")
		}

		test, _ = translate("ATGAAA")
		if len(test) == 0 {
			t.Errorf("STOP Identified out of alignment with Start")
		}

		test, _ = translate("ATGAAA")
		if len(test) == 0 {
			t.Errorf("STOP Identified out of alignment with Start")
		}
	*/
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

	//Time Performance
	start1 := time.Now()
	seq1 := "ATG" + strings.Repeat("A", 10000)
	searchSequence(seq1, "Asp")
	duration1 := time.Since(start1)
	if duration1 > 5 {
		t.Errorf("Slow.")
	}

	start2 := time.Now()
	seq2 := "ATG" + strings.Repeat("A", 100000000) //After 1 billion bp it crashes the terminal. Why is this? How can I make this perform better?
	searchSequence(seq2, "Asp")
	duration2 := time.Since(start2)
	if duration2 > 5 {
		t.Errorf("Slow.")
	}

}
