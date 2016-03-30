# humanhash [![GoDoc](https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat)](http://godoc.org/github.com/wolfeidau/humanhash)

This golang library converts a arbitry array of bytes into a string composed of words selected from the `DefaultWordList` using the method in https://github.com/zacharyvoase/humanhash.

It is very handy if you have a digest and you want to reduce it into something memorable.

# Example

```go

input :=  []byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151}

// take the input and map it to 4 words
result, _ := humanhash.Humanize(input, 4)

// prints "result = sodium-magnesium-nineteen-hydrogen"
log.Printf("result = %s", result)

```

# Goroutine Safe Function
Use function HumanizeUsing if you need to generate human hashes from within multiple goroutines and you don't want to use the default word list or delimiter.

```go

input :=  []byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151}

// take the input and map it to 4 words from your own keyword list, using . as the word delimiter.
result, _ := humanhash.HumanizeUsing(input, 4, myKeywordList, ".")

```
Refer to human_hash_test.go for a working example.

# Sponsor

This project was made possible by [Ninja Blocks](http://ninjablocks.com).

# License

This code is Copyright (c) 2014 Mark Wolfe and licenced under the MIT licence. All rights not explicitly granted in the MIT license are reserved. See the included LICENSE.md file for more details.

