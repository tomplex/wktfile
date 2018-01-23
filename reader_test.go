package wktfile

import (
	"testing"
)

const (
	testDataDir = "./testdata/"
)

var (
	testHeader  = []string{"point_id", "wkt"}
	testContent = [][]string{
		{"1", "POINT(-72.549738631281 43.724067830418)"},
		{"2", "POINT(-73.3208999067471 44.9213993671176)"},
		{"3", "POINT(-72.3267250008435 44.3847434287252)"},
	}
)

func headersMatch(t *testing.T, wkt *WKTFile) {
	for i, col := range wkt.Header {
		if col != testHeader[i] {
			t.Errorf("Headers do not match")
		}
	}
}

func contentsMatch(t *testing.T, wkt *WKTFile) {
	for rIdx, row := range wkt.Rows {
		for cIdx, value := range row {
			if value != testContent[rIdx][cIdx] {
				t.Errorf("File contents do not match")
			}
		}
	}
}

func TestReadDefault(t *testing.T) {
	testFile := testDataDir + "test_points.wkt"

	wkt, err := Read(testFile)
	if err != nil {
		t.Errorf("Could not read file")
	}

	headersMatch(t, wkt)
	contentsMatch(t, wkt)

}

func TestReadComma(t *testing.T) {
	testFile := testDataDir + "test_points_comma.wkt"

	wkt, err := Read(testFile, CommaDelimited)
	if err != nil {
		t.Errorf("Could not read file")
	}

	headersMatch(t, wkt)
	contentsMatch(t, wkt)

}

func TestReadTab(t *testing.T) {
	testFile := testDataDir + "test_points_tab.wkt"

	wkt, err := Read(testFile, TabDelimited)
	if err != nil {
		t.Errorf("Could not read file")
	}

	headersMatch(t, wkt)
	contentsMatch(t, wkt)

}

func TestReadCustom(t *testing.T) {
	testFile := testDataDir + "test_points_custom.wkt"

	wkt, err := Read(testFile, CustomDelimiter(';'))
	if err != nil {
		t.Errorf("Could not read file")
	}

	headersMatch(t, wkt)
	contentsMatch(t, wkt)

}

func TestReadNoHeader(t *testing.T) {
	testFile := testDataDir + "test_points_noheader.wkt"

	wkt, err := Read(testFile, HasNoHeader)
	if err != nil {
		t.Errorf("Could not read file")
	}

	contentsMatch(t, wkt)

}
