// Package humanhash provides some methods to reduce a hash or uuid bytes into
// an array of words which are concatenated together in a more memorable string.
// The aim is to represent a hash in a form which is easier to recognise than a
// hex or base64 encoded string.
//
// Example usage:
//   input :=  []byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151}
//
//   // take the input and map it to 4 words
//   result, _ := humanhash.Humanize(input, 4)
//
//   // prints "result = sodium-magnesium-nineteen-hydrogen"
//   log.Printf("result = %s", result)
//
package humanhash

import (
	"errors"
	"strings"
)

var (
	ErrTargetLessThanInput = errors.New("Fewer input bytes than requested target")
	ErrIncorrectSizeList   = errors.New("Word list must contain 256 words")
)

// The default word list.
var DefaultWordList = []string{
	"ack", "alabama", "alanine", "alaska", "alpha", "angel", "apart", "april",
	"arizona", "arkansas", "artist", "asparagus", "aspen", "august", "autumn",
	"avocado", "bacon", "bakerloo", "batman", "beer", "berlin", "beryllium",
	"black", "blossom", "blue", "bluebird", "bravo", "bulldog", "burger",
	"butter", "california", "carbon", "cardinal", "carolina", "carpet", "cat",
	"ceiling", "charlie", "chicken", "coffee", "cola", "cold", "colorado",
	"comet", "connecticut", "crazy", "cup", "dakota", "december", "delaware",
	"delta", "diet", "don", "double", "early", "earth", "east", "echo",
	"edward", "eight", "eighteen", "eleven", "emma", "enemy", "equal",
	"failed", "fanta", "fifteen", "fillet", "finch", "fish", "five", "fix",
	"floor", "florida", "football", "four", "fourteen", "foxtrot", "freddie",
	"friend", "fruit", "gee", "georgia", "glucose", "golf", "green", "grey",
	"hamper", "happy", "harry", "hawaii", "helium", "high", "hot", "hotel",
	"hydrogen", "idaho", "illinois", "india", "indigo", "ink", "iowa",
	"island", "item", "jersey", "jig", "johnny", "juliet", "july", "jupiter",
	"kansas", "kentucky", "kilo", "king", "kitten", "lactose", "lake", "lamp",
	"lemon", "leopard", "lima", "lion", "lithium", "london", "louisiana",
	"low", "magazine", "magnesium", "maine", "mango", "march", "mars",
	"maryland", "massachusetts", "may", "mexico", "michigan", "mike",
	"minnesota", "mirror", "mississippi", "missouri", "mobile", "mockingbird",
	"monkey", "montana", "moon", "mountain", "muppet", "music", "nebraska",
	"neptune", "network", "nevada", "nine", "nineteen", "nitrogen", "north",
	"november", "nuts", "october", "ohio", "oklahoma", "one", "orange",
	"oranges", "oregon", "oscar", "oven", "oxygen", "papa", "paris", "pasta",
	"pennsylvania", "pip", "pizza", "pluto", "potato", "princess", "purple",
	"quebec", "queen", "quiet", "red", "river", "robert", "robin", "romeo",
	"rugby", "sad", "salami", "saturn", "september", "seven", "seventeen",
	"shade", "sierra", "single", "sink", "six", "sixteen", "skylark", "snake",
	"social", "sodium", "solar", "south", "spaghetti", "speaker", "spring",
	"stairway", "steak", "stream", "summer", "sweet", "table", "tango", "ten",
	"tennessee", "tennis", "texas", "thirteen", "three", "timing", "triple",
	"twelve", "twenty", "two", "uncle", "undress", "uniform", "uranus", "utah",
	"vegan", "venus", "vermont", "victor", "video", "violet", "virginia",
	"washington", "west", "whiskey", "white", "william", "winner", "winter",
	"wisconsin", "wolfram", "wyoming", "xray", "yankee", "yellow", "zebra",
	"zulu"}

// SetWordList allows you to override the default word list used by the Humanize method.
// This list of words MUST contain 256 entries.
func SetWordList(words []string) error {

	if len(words) != 256 {
		return ErrIncorrectSizeList
	}

	DefaultWordList = words

	return nil
}

// Humanize takes a digest or some array of bytes, compresses it and selects a number of
// words to represent it. The selection of words will occur the same for the a matching
// hash but it isn't reversable to the hash.
func Humanize(digest []byte, words int) (string, error) {

	var w []string

	c, err := Compress(digest, words)

	if err != nil {
		return "", err
	}

	for _, b := range c {
		w = append(w, DefaultWordList[b])
	}

	return strings.Join(w, "-"), nil
}

// Compress an array of bytes to the target size using a simple xor.
func Compress(digest []byte, target int) ([]byte, error) {
	results := make([]byte, target)

	length := len(digest)

	if target > length {
		return results, ErrTargetLessThanInput
	}

	segmentSize := length / target

	segments := make([][]byte, target)

	for i := range segments {
		segments[i] = digest[i*segmentSize : (i+1)*segmentSize]
	}

	remainder := length % target

	// if there is some remaining bytes after dividing up the input
	// append them to the last segment
	if remainder > 0 {
		segments[len(segments)-1] = append(segments[len(segments)-1], digest[segmentSize*target:length]...)
	}

	// xor each segment into it's respective bucket
	for i := range segments {
		for b := range segments[i] {
			results[i] = results[i] ^ segments[i][b]
		}
	}

	return results, nil
}
