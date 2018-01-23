// Package wktfile provides an easy-to-use wrapper around a WKT file, with options for parsing common formats.
package wktfile

import (
	"os"
	"encoding/csv"
)

var defaults = [2]func(*WKTFile) error{PipeDelimiter, FileHasHeader}

// WKTFile is the main type provided by package wkt. It provides a thin wrapper over a file, reading the contents into
// a slice of []string's representing rows. It also provides access to the source file path and the header (if it has one).
type WKTFile struct {
	FilePath string // the full path of the input file
	Header []string // the file's header
	Rows [][]string // the file's data: a slice of

	// Private helpers defining WKTFile parse behavior

	delimiter rune // the delimiter for this file. Will default to '|'
	hasheader bool
}


// Read reads and parses a standard text-based WKT file, allowing the caller to specify some common options for parsing.
// It returns a WKTFile representation of the input file.
func Read(filepath string, options ...func(*WKTFile) error) (*WKTFile, error) {
	wkt := &WKTFile{}

	for _, f := range defaults { // Set default WKTFile options
		err := f(wkt)
		if err != nil {
			return nil, err
		}
	}

	for _, f := range options { // User options override defaults
		err := f(wkt)
		if err != nil {
			return nil, err
		}
	}

	fileContents, err := readfile(filepath, wkt.delimiter)
	if err != nil {
		return nil, err
	}

	if wkt.hasheader {
		wkt.Header = fileContents[0]
		wkt.Rows = fileContents[1:]
	} else {
		wkt.Rows = fileContents
	}

	return wkt, nil
}

// readfile is a helper function to get all rows in a file.
func readfile(filepath string, delimiter rune) ([][]string, error) {

	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.Comma = delimiter

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil

}

// PipeDelimiter is a functional parameter to Read() which specifies that the WKT file is delimited by '|' values.
func PipeDelimiter(wkt *WKTFile) error {
	wkt.delimiter = '|'
	return nil
}

// CommaDelimiter is a functional parameter to Read() which specifies that the WKT file is delimited by ',' values.
func CommaDelimiter(wkt *WKTFile) error {
	wkt.delimiter = ','
	return nil
}

// TabDelimiter is a functional parameter to Read() which specifies that the WKT file is delimited by '\t' values.
func TabDelimiter(wkt *WKTFile) error {
	wkt.delimiter = '\t'
	return nil
}

// CustomDelimiter isa functional parameter to Read() which allows the user to specify a custom delimiter for a file.
func CustomDelimiter(delimiter rune) func(wkt *WKTFile) error {
	return func(wkt *WKTFile) error {
		wkt.delimiter = delimiter
		return nil
	}
}

// FileHasHeader is a functional parameter to Read() which specifies that the file has a header.
func FileHasHeader(wkt *WKTFile) error {
	wkt.hasheader = true
	return nil
}

// FileNoHeader is a functional parameter to Read() which specifies that the file does not have a header.
func FileNoHeader(wkt *WKTFile) error {
	wkt.hasheader = false
	return nil
}
