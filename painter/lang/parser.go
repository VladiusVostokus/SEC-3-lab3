package lang

import (
	"bufio"
	"fmt"
	"github.com/VladiusVostokus/SEC-3-lab3/painter"
	"io"
	"strings"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		commandLine := scanner.Text()
		commands := strings.Split(commandLine, " ")
		//fmt.Println(words)

		err := p.parse(commands)
		if err != nil {
			return nil, err
		}

		//op := parse(commandLine) // parse the line to get Operation
		//res = append(res, op)
	}

	// TODO: Реалізувати парсинг команд.
	res = append(res, painter.OperationFunc(painter.WhiteFill))
	res = append(res, painter.UpdateOp)

	return res, nil
}

func (p *Parser) parse(commands []string) error {
	command := commands[0]
	wordsLen := len(commands)
	switch command {
	case "white":
		fmt.Println("white background")
		if wordsLen != 1 {
			return fmt.Errorf("too many parameters for white command")
		}
		fmt.Println("white background")

	case "green":
		if wordsLen != 1 {
			return fmt.Errorf("too many parameters for green command")
		}
		fmt.Println("green background")

	case "update":
		if wordsLen != 1 {
			return fmt.Errorf("too many parameters for updare command")
		}
		fmt.Println("update texture")

	case "figure":
		if wordsLen != 3 {
			return fmt.Errorf("must be 2 parametrs for this command")
		}
		fmt.Println("new figure with coords", commands[1], commands[2])
	case "move":
		if wordsLen != 3 {
			return fmt.Errorf("must be 2 parametrs for this command")
		}
		fmt.Println("move figure to", commands[1], commands[2])
	case "bgrect":
		if wordsLen != 5 {
			return fmt.Errorf("must be 4 parametrs for this command")
		}
		fmt.Println("draw background with size", commands[1], commands[2], commands[3], commands[4])
	case "reset":
		fmt.Println("reset texture state")
	default:
		return fmt.Errorf("wrond command")
	}
	return nil
}
