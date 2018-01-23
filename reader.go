// Package wkt provides an easy-to-use wrapper around a WKT file, with options for parsing common formats.

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


// Read reads a standard text-based WKTFile file with certain common options available.
// It returns a WKTFile representation of the input file.
func Read(filepath string, options ...func(*WKTFile) error) (*WKTFile, error) {
	wktFile := &WKTFile{}

	for _, f := range defaults { // Set default WKTFile options
		err := f(wktFile)
		if err != nil {
			return nil, err
		}
	}

	for _, f := range options { // User options override defaults
		err := f(wktFile)
		if err != nil {
			return nil, err
		}
	}

	fileContents, err := readfile(filepath, wktFile.delimiter)
	if err != nil {
		return nil, err
	}



	if wktFile.hasheader {
		wktFile.Header = fileContents[0]
		wktFile.Rows = fileContents[1:]
	} else {
		wktFile.Rows = fileContents
	}

	return wktFile, nil
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

// PipeDelimiter is an option for wkt.Read() which specifies that the WKT file is delimited by '|' values.
func PipeDelimiter(_wkt *WKTFile) error {
	_wkt.delimiter = '|'
	return nil
}

// CommaDelimiter is an option for wkt.Read() which specifies that the WKT file is delimited by ',' values.
func CommaDelimiter(_wkt *WKTFile) error {
	_wkt.delimiter = ','
	return nil
}

// TabDelimiter is an option for wkt.Read() which specifies that the WKT file is delimited by '\t' values.
func TabDelimiter(_wkt *WKTFile) error {
	_wkt.delimiter = '\t'
	return nil
}

// CustomDelimiter is an option for wkt.Read() which allows the user to specify a custom delimiter for a file.
func CustomDelimiter(delimiter rune) func(_wkt *WKTFile) error {
	return func(_wkt *WKTFile) error {
		_wkt.delimiter = delimiter
		return nil
	}
}

// FileHasHeader is an option for wkt.Read() which specifies that the file has a header.
func FileHasHeader(_wkt *WKTFile) error {
	_wkt.hasheader = true
	return nil
}

// FileNotHasHeader is an option for wkt.Read() which specifies that the file does not have a header.
func FileNotHasHeader(_wkt *WKTFile) error {
	_wkt.hasheader = false
	return nil
}
