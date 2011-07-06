package main

import "os"
import "io"

func main() {

}

func getItemFromQueue() string {
	
}

func addToQueue(word string) {
	
}

func pushToWorker(word string) {
	
}

//Receives an array from the worker with friends of the word
func receiveFromWorker(word string, friends [] string) {
	s := "Base: " + word + "\n"
	if len(friends) > 0 {
		for i := 0; i < len(friends); i++ {
			s += "    " + friends[i] + "\n"
		}
	}
}

//Used to write the file results to a file
func writeToFinal(fileName string, word string) os.Error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0x775)
	if err != nil {
			return err 
		}
		ba := []byte(word)
		n, err := f.Write(ba)
		f.Close()
		if err == nill && n < len(ba) {
			err = io.ErrShortWrite
		}
		return err
}
