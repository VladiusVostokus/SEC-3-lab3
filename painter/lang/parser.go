package lang

import (
	"bufio"
	"fmt"
	"github.com/VladiusVostokus/SEC-3-lab3/painter"
	"io"
	"strconv"
	"strings"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation

	res = append(res, painter.OperationFunc(painter.WhiteFill))
	res = append(res, painter.UpdateOp)

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		commandLine := scanner.Text()
		commands := strings.Split(commandLine, " ")

		op, err := p.parse(commands)
		if err != nil {
			return nil, err
		}

		res = append(res, op)
	}

	// TODO: Реалізувати парсинг команд.

	return res, nil
}

func (p *Parser) parse(commands []string) (painter.Operation, error) {
	command := commands[0]
	wordsLen := len(commands)
	var op painter.Operation
	switch command {
	case "white":
		if wordsLen != 1 {
			return nil, fmt.Errorf("too many parameters for white command")
		}
		op = painter.OperationFunc(painter.WhiteFill)

	case "green":
		if wordsLen != 1 {
			return nil, fmt.Errorf("too many parameters for green command")
		}
		op = painter.OperationFunc(painter.GreenFill)

	case "update":
		if wordsLen != 1 {
			return nil, fmt.Errorf("too many parameters for updare command")
		}
		op = painter.UpdateOp

	case "figure":
		if wordsLen != 3 {
			return nil, fmt.Errorf("must be 2 parametrs for this command")
		}
		//fmt.Println("new figure with coords", commands[1], commands[2])
		x, _ := strconv.Atoi(commands[1])
		y, _ := strconv.Atoi(commands[2])
		painter.SetMoveCoords(x, y)
		op = painter.OperationFunc(painter.DrawCross)

	case "move":
		if wordsLen != 3 {

			return nil, fmt.Errorf("must be 2 parametrs for this command")
		}
		fmt.Println("move figure to", commands[1], commands[2])

	case "bgrect":
		if wordsLen != 5 {
			return nil, fmt.Errorf("must be 4 parametrs for this command")
		}
		x1, _ := strconv.Atoi(commands[1])
		y1, _ := strconv.Atoi(commands[2])
		x2, _ := strconv.Atoi(commands[3])
		y2, _ := strconv.Atoi(commands[4])
		bgr := painter.BackGroundRect{x1, y1, x2, y2}
		op = &bgr
	case "reset":
		if wordsLen != 1 {
			return nil, fmt.Errorf("too many parameters for green command")
		}
		op = painter.OperationFunc(painter.Reset)

	default:
		return nil, fmt.Errorf("wrond command")
	}
	return op, nil
}
