// Levenshtein distance in Go

// i'm just going to plop this func into the main file for now... being lazy!

package editdist

func EditDist(a, b string) (min int) {
	switch {
		case len(a) == 0:
			min = len(b)
		case len(b) == 0:
			min = len(a)
		default:
			var i, j, k int
			// setting i
			if a[0] == b[0] {
				i = EditDist(a[1:], b[1:])
			} else {
				i = 1 + EditDist(a[1:], b[1:])
			}
			// setting j
			j = 1 + EditDist(a[1:], b)
			// setting k
			k = 1 + EditDist(a, b[1:])
			
			// return the min
			if i < j {
				min  = i
			} else {
				min = j
			}
			if k < min {
				min = k
			}
	}
	return
}

// absolute value function
func abs(n int) int {
	if n < 0 {
		n * -1
	}
	return n
}

// trying a different algo for speed & convenience!
func MaxDist(str1, str2 string, max int) (dist int, finish bool) {

	// short-circuit for a == b
	if str1 == str2 { return 0, true }
	// short-circuit for max == 0 (see ^)
	// as if a != b && max == 0 this is what we will get
	if max == 0 { return 0, false }
	
	// getting the lengths of the strings...
	len1 = len(str1)
	len2 = len(str2)

	// short-circuit for len(a || b) == 0
	if len1 == 0 { return len2, len2 > max }
	if len2 == 0 { return len1, len1 > max }
	// short-circuit for strings that are simply too long to be max dist away from each other
	if abs(len1-len2) > max { return max, false }

	ary = make([]slice, len1)


	return // something :P
}

func Friend(a, b string) (ret bool, n int) {
	// our default is > 1 so we dont get a false positive
	dist := 70
	// testing to see if len(a) +- 1 == len(b)
	switch {
		case len(a) == len(b):
			dist = EditDist(a, b)
		case len(a)+1 == len(b):
			dist = EditDist(a, b)
		case len(a)-1 == len(b):
			dist = EditDist(a, b)
		default:
			ret = false
	}
	// switchin like a transistor yo!
	switch dist {
		case -1, 0, 1: ret = true
		default: ret = false
	}
	return ret, dist
}
