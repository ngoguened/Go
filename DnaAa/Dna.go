package dna

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const validNucleotides = "atcg"
const validAminoAcids = "RHKDESTNQCGPAVILMFYW$"

// Instead of creating a DNA and AA data type, we validate characters based off of constants which match a representation of DNA and AA.

func validNucleotide(d byte) bool {
	return strings.Contains(validNucleotides, string(d)) // Maybe char occurs in string function.
}

func validAminoAcid(a byte) bool {
	return strings.Contains(validAminoAcids, string(a))
}

/*_______________________________________________________________________________________________________________________*/
// Representation conversions

func aminoByteToThreeLetter(a byte) (string, error) {
	switch a {
	case 'R':
		return "Arg", nil
	case 'H':
		return "His", nil
	case 'K':
		return "Lys", nil
	case 'D':
		return "Asp", nil
	case 'E':
		return "Glu", nil
	case 'S':
		return "Ser", nil
	case 'T':
		return "Thr", nil
	case 'N':
		return "Asn", nil
	case 'Q':
		return "Gln", nil
	case 'C':
		return "Cys", nil
	case 'G':
		return "Gly", nil
	case 'P':
		return "Pro", nil
	case 'A':
		return "Ala", nil
	case 'V':
		return "Val", nil
	case 'I':
		return "Ile", nil
	case 'L':
		return "Leu", nil
	case 'M':
		return "Met", nil
	case 'F':
		return "Phe", nil
	case 'Y':
		return "Tyr", nil
	case 'W':
		return "Trp", nil
	default:
		return "", fmt.Errorf("Invalid input.")
	}
}

func aminoThreeLetterToByte(a string) (byte, error) {
	switch a {
	case "Arg":
		return 'R', nil
	case "His":
		return 'H', nil
	case "Lys":
		return 'K', nil
	case "Asp":
		return 'D', nil
	case "Glu":
		return 'E', nil
	case "Ser":
		return 'S', nil
	case "Thr":
		return 'T', nil
	case "Asn":
		return 'N', nil
	case "Gln":
		return 'Q', nil
	case "Cys":
		return 'C', nil
	case "Gly":
		return 'G', nil
	case "Pro":
		return 'P', nil
	case "Ala":
		return 'A', nil
	case "Val":
		return 'V', nil
	case "Ile":
		return 'I', nil
	case "Leu":
		return 'L', nil
	case "Met":
		return 'M', nil
	case "Phe":
		return 'F', nil
	case "Tyr":
		return 'Y', nil
	case "Trp":
		return 'W', nil
	default:
		return 0, fmt.Errorf("Invalid input.")
	}
}

func sequenceToAmino(s string) (byte, error) {
	switch s {
	case "ttt", "ttc":
		return 'F', nil
	case "tta", "ttg", "ctt", "ctc", "cta", "ctg":
		return 'L', nil
	case "att", "atc", "ata":
		return 'I', nil
	case "atg":
		return 'M', nil
	case "gtt", "gtc", "gta", "gtg":
		return 'V', nil
	case "tct", "tcc", "tca", "tcg", "agt", "agc":
		return 'S', nil
	case "cct", "ccc", "cca", "ccg":
		return 'P', nil
	case "act", "acc", "aca", "acg":
		return 'T', nil
	case "gct", "gcc", "gca", "gcg":
		return 'A', nil
	case "tat", "tac":
		return 'Y', nil
	case "cat", "cac":
		return 'H', nil
	case "caa", "cag":
		return 'Q', nil
	case "aat", "aac":
		return 'N', nil
	case "aaa", "aag":
		return 'K', nil
	case "gat", "gac":
		return 'N', nil
	case "gaa", "gag":
		return 'E', nil
	case "tgt", "tgc":
		return 'C', nil
	case "tgg":
		return 'W', nil
	case "cgt", "cgc", "cga", "cgg", "aga", "agg":
		return 'R', nil
	case "ggt", "ggc", "gga", "ggg":
		return 'G', nil
	case "taa", "tag", "tga":
		return '$', nil //$ Represents stop. Despite not being a direct translation, this function is private so it cannot be used maliciously.
	default:
		return 0, fmt.Errorf("Invalid input")
	}
}

/*_______________________________________________________________________________________________________________________*/
// Translation

func translate(inputFilepath string) (*os.File, error) {
	// Idiomize to file
	file, err := os.Open(inputFilepath)
	if err != nil {
		return nil, fmt.Errorf("file could not be read")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ds := ""
	for scanner.Scan() {
		ds = ds + scanner.Text()
	}

	ds = strings.ToLower(ds)

	//Validate sequence
	for i := 0; i < len(ds); i++ {
		if !validNucleotide(ds[i]) {
			return nil, fmt.Errorf("Invalid input sequence.")
		}
	}

	// Find start codon
	start := strings.Index(ds, "atg")
	if start == -1 {
		return nil, fmt.Errorf("No start sequence")
	}

	// Tidy sequence
	ds = ds[start:]
	lenDs := len(ds) // Turn 4 linear operations into 1.
	ds = ds[:lenDs-(lenDs%3)]

	// Create output file. If filename exists, writes to existing file.
	outputfile, err := os.OpenFile("DNAoutput.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer outputfile.Close()

	// Compare sequence and convert to amino

	for i := 0; i < len(ds); i = i + 3 {
		a, err := sequenceToAmino(ds[i : i+3])
		if err != nil {
			return nil, fmt.Errorf("Conversion from sequence to amino acid failed starting at index %d", i+start)
		}
		if a == '$' { // Represents stop codon. Stops translating.
			break
		}
		outputfile.Write([]byte{a})
	}
	return outputfile, nil
}

/*_______________________________________________________________________________________________________________________*/
// Searching

func searchSequence(inputFilepath string, q string) (bool, error) { //Must be three letter amino input code
	qByte, err := aminoThreeLetterToByte(q)
	if err != nil {
		return false, fmt.Errorf("Amino input could not be read.")
	}

	translateFile, err := translate(inputFilepath)
	if err != nil {
		return false, fmt.Errorf("Sequence could not be translated.")
	}

	as, err := os.ReadFile(translateFile.Name())

	for _, a := range as {
		if a == qByte {
			return true, nil //Found match
		}
	}
	return false, nil //No match
}
