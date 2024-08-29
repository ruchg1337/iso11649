package iso11649

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {

	inFile, err := os.Open("data/test.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		want := scanner.Text()
		got, err := GenerateReference(want[4:])
		if err != nil {
			t.Errorf("GenerateReference return an error%s:", err)
		}
		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	}
}
