package dialog

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

const (
	yes = "y"
	no  = "n"
)

type Dialog struct {
	reader *bufio.Reader
}

func NewDialog() *Dialog {
	return &Dialog{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (d *Dialog) GetAnswerForYesNoQuestion(
	_ context.Context,
	question string,
) (bool, error) {
	var err error
	answer := ""

	fmt.Println(question + " (y/n):")

	for {
		answer, err = d.reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("failed to read answer: %w", err)
		}

		answer = strings.TrimSpace(answer)

		if answer != yes && answer != no {
			fmt.Println("Expected answer should be y/n. Provided: " + answer)
		} else {
			break
		}
	}

	return answer == yes, nil
}
