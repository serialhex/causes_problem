// Levenshtein distance in Go

// i'm just going to plop this func into the main file for now... being lazy!

package editdist

func EditDistance(a, b string) int {
	switch {
		case len(a) == 0 {
			return len(b)
		}
		case len(b) == 0 {
			return len(a)
		}
		default {
			var i, j, k int
			// setting i
			if a[0] == b[0] {
				i = Levenshtein(a[1:], b[1:])
			} else {
				i = 1 + Levenshtein(a[1:], b[1:])
			}
			// setting j
			j = 1 + Levenshtein(a[1:], b)
			// setting k
			j = 1 + Levenshtein(a, b[1:])
			
			// return the min
			min := 0
			if i < j {
				min  = i
			} else {
				min = j
			}
			if k < min {
				return k
			} else {
				return min
			}
		}
	}
}