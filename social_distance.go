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
	"./file"
	"fmt"
)

// putting it here cause i'm lazy and want to compile less things...
func editDist(a, b string) int {
	min := 0
	switch {
		case len(a) == 0: 
			min = len(b)
		case len(b) == 0:
			min = len(a)
		default:
			var i, j, k int
			// setting i
			if a[0] == b[0] {
				i = editDist(a[1:], b[1:])
			} else {
				i = 1 + editDist(a[1:], b[1:])
			}
			// setting j
			j = 1 + editDist(a[1:], b)
			// setting k
			k = 1 + editDist(a, b[1:])
			
			// return the min
			if i < j {
				min  = i
			} else {
				min = j
			}
			if k < min {
				min = k
			} else {
				return min
			}
	}
	return min
}

// will initialize dict here...
func loadDict(fname string) []string {
	opn, _ := file.Open(fname)
	reader := bufio.NewReader(opn)

	theGuess := 100 // the file is ~ this big...
	dict := make([]string, 0, theGuess)

	line, _, _ := reader.ReadLine()

}

func main() {

	dict := loadDict("word.small")
	fmt.Printf(dict)

	/* temporarily...
	for i := 0; i < len(dict); i++ {
		for j := i+1; j < len(dict); j++ {
			fmt.Print(dict[i], ", ", dict[j], " edit distance: ", editDist(dict[i], dict[j]), "\n")
		}
	}
	*/
}