package wktfile


import (
	"testing"
)

const testDataDir = "./testdata/"



func TestReadDefaultOptions(t *testing.T) {
	testFile := testDataDir + "test_points.wkt"
	testHeader := []string{"point_id", "wkt"}


	wkt, err := Read(testFile)
	if err != nil {
		t.Errorf("Could not read file")
	}

	for i := range wkt.Header {
		
	}

}