## Damerau–Levenshtein distance
This is a spelling corrector implementing the Damerau-Levenshtein distance. Takes a string value input from the user. Looks for an identical word on a list of words, if none is found, look for a similar word.

## Getting Started

* Go programming language [installation instructions](http://golang.org/doc/install).
* For setting up the Go workspace check out [How to Write Go Code](http://golang.org/doc/code.html#tmp_13) or [Get started with Go](http://youtu.be/2KmHtgtEZ1s).
* Check out the package documentation page: [damerau GoPkgDoc Documentation](http://go.pkgdoc.org/github.com/jhprks/damerau)

### To install the damerau package:

    $ go get github.com/jhprks/damerau
    
To use it in a program:

```Go
import "github.com/jhprks/damerau"
```

### To install it immediately for command-line use:

    $ go get github.com/jhprks/damerau/damerau

**The program must be given two arguments for command line use:**

1. Path to a text file containing a lot of words (e.g. the word.txt file in the damerau folder in this repository)
2. Single erroneous string value (e.g. korrecter, commanderr)

**Example command-line usage:**

    $ damerau -file="words.txt" -word="korrecter"

Output: `[correct correctly correct.]`