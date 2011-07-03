package wordsort


func WordSort(dict []string) (map[int] map[byte] []string) {
	// maybe use nested slices instead of maps? might be faster...
	words := make(map[int] map[byte] []string)
	for _, str := range dict {
		 if words[len(str)] == nil {
		 	words[len(str)] = make(map[byte] []string)
		 }
		words[len(str)][str[0]] = append(words[len(str)][str[0]], str)
	}
	return words
}

// cooler version of the above from jteeuwen
func SortMap(dict []string) (m map[int] map[byte] []string) {
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

// designed by jteeuwen
// returns the index of the word in []dict
func SortSlice(dict []string) (m [][][]int) {
    var i, l int
    var c byte

    for i = range dict {
        if l = len(dict[i]); l == 0 {
            continue
        }

        if len(m) <= l {
            t := make([][][]int, l+1)
            copy(t, m)
            m = t
        }

        c = dict[i][0]
        if len(m[l]) <= int(c) {
            t := make([][]int, c+1)
            copy(t, m[l])
            m[l] = t
        }

        m[l][c] = append(m[l][c], i)
    }

    return
}