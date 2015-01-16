package humanhash

import (
	"encoding/hex"
	"testing"
)

func TestCompress(t *testing.T) {

	fixtures := []struct {
		input  []byte
		result []byte
	}{
		{
			input:  []byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151},
			result: []byte{205, 128, 156, 96},
		},
	}

	for _, fixture := range fixtures {
		res := Compress(fixture.input, 4)
		//		log.Printf("res %v", res)
		if hex.Dump(res) != hex.Dump(fixture.result) {
			t.Errorf("humanize didn't match expected=%v actual=%v", fixture.result, res)
		}
	}

}

func TestHumanize(t *testing.T) {

	fixtures := []struct {
		input  []byte
		result string
	}{
		{
			input:  []byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151},
			result: "sodium-magnesium-nineteen-hydrogen",
		},
	}

	for _, fixture := range fixtures {
		res := Humanize(fixture.input, 4)
		//		log.Printf("res %v", res)
		if res != fixture.result {
			t.Errorf("humanize didn't match expected=%s actual=%s", fixture.result, res)
		}
	}

}
