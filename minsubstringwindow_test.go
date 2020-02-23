package minimumsubstringwindow

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	S      string `json:"s"`
	T      string `json:"t"`
	Output string `json:"output"`
}

func TestMinWindow(t *testing.T) {
	f, err := os.Open("tests.json")

	if err != nil {
		t.Error(err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)
	tests := make(map[string]Test)

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				t.Run(name, func(st *testing.T) {
					res := MinWindow(test.S, test.T)

					if res != test.Output {
						st.Errorf("result returned is %s", res)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			t.Error(err)
			break
		}
	}
}
