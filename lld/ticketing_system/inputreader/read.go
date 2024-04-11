package inputreader

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Read reads input from the provided io.Reader(file or stdin)
// stops reading when user types exit or  EOF is received
func (i *InputReader) Read(r io.Reader, decorator string, exitText string) {
	reader := bufio.NewReader(r)
	for {
		fmt.Print(decorator)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		text = strings.ToLower(strings.Replace(text, "\n", "", -1))
		if text == "" {
			continue
		}
		if text == exitText {
			return
		}

		// 1. parse the instructions.
		// 2. verify the instructions.
		// 3. excute the instructions and display the results.
		fmt.Println("EXTRACTED TEXT", text)
	}
}
