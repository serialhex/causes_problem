// Levenshtein distance in Go

// i'm just going to plop this func into the main file for now... being lazy!

package editdist

func EditDist(a, b string) (min int) {
	switch {
		case len(a) == 0:
			return len(b)
		case len(b) == 0:
			return len(a)
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
			j = 1 + EditDist(a, b[1:])
			
			// return the min
			min := 0
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


func Friend(a, b string) (ret bool) {
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
	return
}