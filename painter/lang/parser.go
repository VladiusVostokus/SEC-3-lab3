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
		words := strings.Split(commandLine, " ")
		//fmt.Println(words)
		command := words[0]
		switch command {
		case "white":
			fmt.Println("white background")
		case "green":
			fmt.Println("green background")
		case "update":
			fmt.Println("update texture")
		case "figure":
			fmt.Println("new fugire with coords")
		case "move":
			fmt.Println("move figure to")
		case "bgrect":
			fmt.Println("draw background with size")
		}
		//op := parse(commandLine) // parse the line to get Operation
		//res = append(res, op)
	}

	// TODO: Реалізувати парсинг команд.
	res = append(res, painter.OperationFunc(painter.WhiteFill))
	res = append(res, painter.UpdateOp)

	return res, nil
}
