// social distance in Go

/* social distance & thoughts
  
  * from:
  http://jobs.github.com/positions/181fb26e-71cc-11e0-96a9-8752f87b91a0
  jobs@causes.com
	
	Two words are friends if they have a Levenshtein distance of 1. That is, you can add, remove, or substitute exactly one letter in word X to create word Y. A word’s social network consists of all of its friends, plus all of their friends, and all of their friends’ friends, and so on. Write a program to tell us how big the social network for the word “causes” is, using this word list

  this is my implementation of social distance for causes.com

  * helpers...
  http://en.wikipedia.org/wiki/Levenshtein_distance

  ---

  so, i figure i have a couple of options...
  
  1) i can either write an algo that computes the Levenshtein distance of two strings, and then just generate an infinite number of strings, and take only the strings that are found in the dictionary...
  - this is just wrong and stupid... horribly inefficient and just dumb (though brute force always works, i know i can do better!)
  
  2) i can iteratively generate strings from a primary string, only changing 1 feature, and seeing if the generated string is in the dictionary.
  - this is probably very search intensive, i mean, i'll be searching through the dict for every string i generate, instead of simply working on the dict itself...  this was the primary reason for thinking of having the dict be symbols instead of strings, but i may change that... 

  3) i can write an implementation of Levenshtein distance and then iterate through the dict looking for words where string.levenshtein_distance(str) == 1.  then iterate though *those* words and do the same and so on, and so forth...
  - this is better, though maybe i'll think of something even better!!

  - Justin Patera aka serialhex

*/

package main

import (
	"./editdist"
	"fmt"
	"bufio"
	"os"
)

// will initialize dict here...
func loadDict(fname string) ([]string, os.Error) {
	fmt.Print("Loading Dictionary\n")
	opn, err := os.Open(fname)
	if err != nil { return nil, err }
	defer opn.Close()
	reader := bufio.NewReader(opn)

	theGuess := 10 // the file is ~ this big...
	dict := make([]string, 0, theGuess)

	// i'm probably doing something wrong here...
	// getting the first line...

	for {
		line, _, err := reader.ReadLine()
		dict = append(dict, string(line))
		if err == os.EOF { break }
	}
	return dict, err
}

func main() {

	dict, err := loadDict("test")
	if err != nil { 
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print(true)
	fmt.Print(false)

	for i := 0; i < len(dict); i++ {
		for j := i+1; j < len(dict); j++ {
			fmt.Print(dict[i], ", ", dict[j], " friends? ", editdist.Friend(dict[i], dict[j]), "\n")
		}
	}
}