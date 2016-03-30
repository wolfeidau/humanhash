package humanhash

import (
	"encoding/hex"
	"fmt"
	_ "log"
	"sort"
	"strings"
	"sync"
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
		res, err := Compress(fixture.input, 4)
		if err != nil {
			t.Error(err)
		}
		//		log.Printf("res %v", res)
		if hex.Dump(res) != hex.Dump(fixture.result) {
			t.Errorf("humanize didn't match expected=%v actual=%v", fixture.result, res)
		}
	}

	fmt.Printf("len = %d\n", len(DefaultWordList))

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
		{
			input:  []byte{0, 255, 141, 13},
			result: "ack-zulu-mississippi-august",
		},
	}

	for _, fixture := range fixtures {
		res, err := Humanize(fixture.input, 4)
		if err != nil {
			t.Error(err)
		}

		//		log.Printf("res %v", res)
		if res != fixture.result {
			t.Errorf("humanize didn't match expected=%s actual=%s", fixture.result, res)
		}
	}

}

// This test case only makes sense when run with "go test -race"
func Test_HumanizeUsing(t *testing.T) {

	// Use a reverse keywords list in one of the goroutines
	reverseKeywords := make([]string, len(DefaultWordList))
	copy(reverseKeywords, DefaultWordList)
	sort.Sort(sort.Reverse(sort.StringSlice(reverseKeywords)))

	// Use an uppercase keywords list in another of the goroutines
	upperKeywords := make([]string, 0)
	for _, w := range DefaultWordList {
		upperKeywords = append(upperKeywords, strings.ToUpper(w))
	}

	fixtures := []struct {
		input     []byte
		result    string
		keywords  []string
		delimiter string
	}{
		{
			input:     []byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151},
			result:    "delta_magazine_india_november",
			keywords:  reverseKeywords,
			delimiter: "_",
		},
		{
			input:     []byte{0, 255, 141, 13},
			result:    "ack-zulu-mississippi-august",
			keywords:  DefaultWordList,
			delimiter: "-",
		},
		{
			input:     []byte{0, 255, 141, 13},
			result:    "ACK.ZULU.MISSISSIPPI.AUGUST",
			keywords:  upperKeywords,
			delimiter: ".",
		},
	}

	var wg sync.WaitGroup

	for _, fixture := range fixtures {
		wg.Add(1)

		go func(input []byte, result string, keywords []string, delimiter string) {
			defer wg.Done()

			res, err := HumanizeUsing(input, 4, keywords, delimiter)
			if err != nil {
				t.Error(err)
			}

			//log.Printf("res %v", res)
			if res != result {
				t.Errorf("humanize didn't match expected=%s actual=%s", result, res)
			}
		}(fixture.input, fixture.result, fixture.keywords, fixture.delimiter)
	}

	wg.Wait()
}
