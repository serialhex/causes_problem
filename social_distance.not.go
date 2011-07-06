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

  4) best idea yet: find a fast & efficient algo for finding the Levenshtein distance, and implement it.  then, scan through the dictionary and sort the words by length & beginning letter. finally, when actually _looking_ for the friends, only look for words that are +-1 len & begin with the same letter, or the same length (and begin with any letter) then use the algo to sort through _thoes_.  stick friends in a map & make sure to set `map[str1] << str2` && `map[str2] << str1` and check for that, because if words are friends, then they're friends right?

  - Justin Patera aka serialhex

*/

package main

import (
	"fmt"
	"bufio"
	"os"
)

// initialize dict here...
func loadDict(fname string) ([]string, os.Error) {
	opn, err := os.Open(fname)
	if err != nil { return nil, err }
	defer opn.Close()
	reader := bufio.NewReader(opn)

	theGuess := 10 // the file is ~ this big...
	dict := make([]string, 0, theGuess)

	for {
		line, _, err := reader.ReadLine()
		if err == os.EOF { break }
		dict = append(dict, string(line))
	}
	return dict, err
}

// finds the edit distance between two strings
func editDist(str1, str2 string) (min int) {
	len1 := len(str1)
	len2 := len(str2)

	if len1 == 0 { return len2 }
	if len2 == 0 { return len1 }

	d := make([]int, len2+1)
	for i := 0; i <= len2; i++ {
		d[i] = i
	}

	for i := 0; i < len1; i++ {
		e := i+1
		for j := 0; j < len2; j++ {
			cost := 0
			if str1[i] != str2[j] { cost = 1 }
			x := d[j+1] +1 		// insertion
			y := e+1			// deletion
			z := d[j] + cost	// substitution
			if x < y {
				if x < z {
					min = x
				} else {
					min = z
				}
			} else if y < z {
				min = y
			} else {
				min = z
			}
			d[j] = e
			e = min
		}
		d[len2] = min
	}
	return
}


func sortMap(dict []string) (m map[int] map[byte] []string) {
// sorts a slice of strings into a map of lenths, beginning letters and the words themselves
    var i, l int
    var c byte
    var ok bool

    m = make(map[int] map[byte] []string)
    for i = range dict {
        if l = len(dict[i]); l == 0 {
            continue  // like next, only different spelling :P
        }

        if _, ok = m[l]; !ok {
            m[l] = make(map[byte] []string)
        }

        c = dict[i][0]
      	m[l][c] = append(m[l][c], dict[i])
        
    }

    return
}

// setting up our uber WordList type...
type WordList struct {
	words map[int] map[byte] []string
}

func NewWordList(dict []string) *WordList {
	l := sortMap(dict)
	return &WordList{words: l}
}

// gets the words that are closest to 'word'
func (w *WordList) findNear(word string) (near []*[]string) {
	l := len(word)
	ch := word[0]
	for _, set := range w.words[l] {
		near = append(near, &set)
	}
	if tmp := w.words[l+1][ch]; tmp != nil {
		near = append(near, &tmp)
	}
	if tmp := w.words[l-1][ch]; tmp != nil {
		near = append(near, &tmp)
	}
	return
}

// finds the friends of a word
func findFriends(word string, list []*[]string) (friends []string) {
	for _, lwrds := range list {
		for i, lwrd := range lwrds {
			if editDist(word, lwrd) == 1 {
				friends = append(friends, lwrd)
				lwrds[i] = ""
			}
		}
	}
	return
}

func uberFriendFindingFunction(input map[string] bool, list WordList) (output map[string] bool) {
	output := make(map[string] bool)
	friends := make([]string, 10)
	
	for word, _ := range input {
		if input[word] {
			continue
		}
		nearby := list.findNear(word)
		tmp := findFriends(word, nearby)
		
		for _, wrd := range tmp {
			friends = append(friends, wrd)
		}
	}

	for _, word := range friends {
		output[word] = true
	}

	for _, word := range friends {
		tmp := uberFriendFindingFunction(output, list)
		for wrd, _ := range tmp {
			output[wrd] = true
		}
	}
	return
}

func main() {

	fmt.Println("loading dict...")
	// dict := []string
	dict, err := loadDict("test")
	if err != nil { 
		fmt.Print(err)
		os.Exit(1)
	}

	// sorting words
	fmt.Println("sorting words...")
	word_list := NewWordList(dict)
	
	// keep friends list in here
	friends := make(map[string] bool, 500)
	friends["causes"] = true

	fmt.Println("doing thingy: ")
	list := uberFriendFindingFunction(friends, wordlist)

	fmt.Println(list)
}