// Levenshtein distance in Go

// i'm just going to plop this func into the main file for now... being lazy!

package editdist
import "fmt"

func EditDist(str1, str2 string) (min int) {
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


// wholesale copy/paste of the above func...
// not very DRY, but whatev!
// it's gotta be fast, i dont care if it's wet too :P

// this dosn't work yet.. moving on for now
func Friends(str1, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)

	if len1 == 0 { return len2 <= 1 }
	if len2 == 0 { return len1 <= 1 }

	d := make([]int, len2+1)
	for i := 0; i <= len2; i++ {
		d[i] = i
	}

	var min int

	friendloop: for i := 0; i < len1; i++ {
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
		fmt.Println("min:", min, "e:", e)
		if e > 1 { 
			break friendloop
		}
		d[len2] = min
	}
	return min <= 1
}