
package worker

import (
	"./editdist"
	"sync"
)

typedef WordList struct {
	words map[int] map[byte] []string
	mu sync.RWMutex
}

func NewWordList() *WordList {
	return WordList(words: make(map[int] map[byte] []string))
}

// shut up about the dumb name!
func WordFriendFinder(word string, words []string) (friends []string) {
	
}