package wordsort

// trying to use the benchmarking capabilites of the test suite
import "testing"

type dictionary struct {
	words []string
}

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
		if err == os.EOF { break }
		dict = append(dict, string(line))
	}
	return dict, err
}

var dict dictionary

// benchmarking
func BenchmarkWordSort(b *testing.B) {
	b.StopTimer()
	dict.words = loadDict("test")
	b.StartTimer()
	wordSort(dict)
}

func BenchmarkSortMap(b *testing.B) {
	b.StopTimer()
	dict.words = loadDict("test")
	b.StartTimer()
	sortMap(dict)
}

func BenchmarkSortSlice(b *testing.B) {
	b.StopTimer()
	dict.words = loadDict("test")
	b.StartTimer()
	sortSlice(dict)
}
