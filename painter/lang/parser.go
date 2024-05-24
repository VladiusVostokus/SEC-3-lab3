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
	allCrosses []*painter.Cross
	bg         painter.Operation
	bgRects    *painter.BackGroundRect
	allMoves   []painter.Operation
	update     painter.Operation
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation
	p.update = nil
	res = append(res, painter.OperationFunc(painter.WhiteFill))
	res = append(res, painter.UpdateOp)

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		commandLine := scanner.Text()
		commands := strings.Split(commandLine, " ")

		err := p.parse(commands)
		if err != nil {
			return nil, err
		}

		res = p.returnOperations()
	}

	// TODO: Реалізувати парсинг команд.

	return res, nil
}

func (p *Parser) parse(commands []string) error {
	command := commands[0]
	wordsLen := len(commands)
	switch command {
	case "white":
		if wordsLen != 1 {
			return fmt.Errorf("too many parameters for white command")
		}
		p.bg = painter.OperationFunc(painter.WhiteFill)

	case "green":
		if wordsLen != 1 {
			return fmt.Errorf("too many parameters for green command")
		}
		p.bg = painter.OperationFunc(painter.GreenFill)

	case "update":
		if wordsLen != 1 {
			return fmt.Errorf("too many parameters for updare command")
		}
		p.update = painter.UpdateOp

	case "figure":
		if wordsLen != 3 {
			return fmt.Errorf("must be 2 parametrs for this command")
		}
		x, _ := strconv.Atoi(commands[1])
		y, _ := strconv.Atoi(commands[2])
		cross := painter.Cross{X: x, Y: y}
		p.allCrosses = append(p.allCrosses, &cross)

	case "move":
		if wordsLen != 3 {

			return fmt.Errorf("must be 2 parametrs for this command")
		}
		x, _ := strconv.Atoi(commands[1])
		y, _ := strconv.Atoi(commands[2])
		m := painter.Move{X: x, Y: y, AllCrosses: p.allCrosses}
		p.allMoves = append(p.allMoves, &m)

	case "bgrect":
		if wordsLen != 5 {
			return fmt.Errorf("must be 4 parametrs for this command")
		}
		x1, _ := strconv.Atoi(commands[1])
		y1, _ := strconv.Atoi(commands[2])
		x2, _ := strconv.Atoi(commands[3])
		y2, _ := strconv.Atoi(commands[4])
		bgr := painter.BackGroundRect{X1: x1, Y1: y1, X2: x2, Y2: y2}
		p.bgRects = &bgr
	case "reset":
		if wordsLen != 1 {
			return fmt.Errorf("too many parameters for green command")
		}
		p.bg = painter.OperationFunc(painter.Reset)
		p.update = nil
		p.allMoves = nil
		p.allCrosses = nil
		p.bgRects = nil

	default:
		return fmt.Errorf("wrond command")
	}
	return nil
}

func (p *Parser) returnOperations() []painter.Operation {
	var ops []painter.Operation
	ops = append(ops, p.bg)
	if len(p.allMoves) > 0 {
		for _, move := range p.allMoves {
			ops = append(ops, move)
		}
	}
	p.allMoves = nil
	if len(p.allCrosses) > 0 {
		for _, figure := range p.allCrosses {
			ops = append(ops, figure)
		}
	}
	if p.bgRects != nil {
		ops = append(ops, p.bgRects)
	}
	if p.update != nil {
		ops = append(ops, p.update)
	}
	return ops
}
