package lang_test

import (
	"strings"
	"testing"

	"github.com/VladiusVostokus/SEC-3-lab3/painter"
	"github.com/VladiusVostokus/SEC-3-lab3/painter/lang"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedOps   []painter.Operation
		expectedError bool
	}{
		{
			name:          "white command",
			input:         "white\n",
			expectedOps:   []painter.Operation{painter.OperationFunc(painter.WhiteFill)},
			expectedError: false,
		},
		{
			name:          "green command",
			input:         "green\n",
			expectedOps:   []painter.Operation{painter.OperationFunc(painter.WhiteFill), painter.OperationFunc(painter.GreenFill)},
			expectedError: false,
		},
		{
			name:          "update command",
			input:         "update\n",
			expectedOps:   []painter.Operation{painter.OperationFunc(painter.WhiteFill), painter.UpdateOp, painter.UpdateOp},
			expectedError: false,
		},
		{
			name:          "figure command",
			input:         "figure 10 20\n",
			expectedOps:   []painter.Operation{painter.OperationFunc(painter.WhiteFill), painter.UpdateOp, &painter.Cross{X: 10, Y: 20}},
			expectedError: false,
		},
		{
			name:          "move command",
			input:         "figure 10 20\nmove 5 5\n",
			expectedOps:   []painter.Operation{painter.OperationFunc(painter.WhiteFill), painter.UpdateOp, &painter.Cross{X: 10, Y: 20}, &painter.Move{X: 5, Y: 5, AllCrosses: []*painter.Cross{{X: 10, Y: 20}}}},
			expectedError: false,
		},
		{
			name:          "bgrect command",
			input:         "bgrect 1 2 3 4\n",
			expectedOps:   []painter.Operation{painter.OperationFunc(painter.WhiteFill), painter.UpdateOp, &painter.BackGroundRect{X1: 1, Y1: 2, X2: 3, Y2: 4}},
			expectedError: false,
		},
		{
			name:          "reset command",
			input:         "reset\n",
			expectedOps:   []painter.Operation{painter.OperationFunc(painter.WhiteFill), painter.UpdateOp, painter.OperationFunc(painter.Reset)},
			expectedError: false,
		},
		{
			name:          "unknown command",
			input:         "unknown\n",
			expectedOps:   nil,
			expectedError: true,
		},
		{
			name:          "too many parameters for white command",
			input:         "white extra\n",
			expectedOps:   nil,
			expectedError: true,
		},
		{
			name:          "too many parameters for green command",
			input:         "green extra\n",
			expectedOps:   nil,
			expectedError: true,
		},
		{
			name:          "too many parameters for update command",
			input:         "update extra\n",
			expectedOps:   nil,
			expectedError: true,
		},
		{
			name:          "incorrect parameters for figure command",
			input:         "figure 10\n",
			expectedOps:   nil,
			expectedError: true,
		},
		{
			name:          "incorrect parameters for move command",
			input:         "move 10\n",
			expectedOps:   nil,
			expectedError: true,
		},
		{
			name:          "incorrect parameters for bgrect command",
			input:         "bgrect 1 2 3\n",
			expectedOps:   nil,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &lang.Parser{}
			reader := strings.NewReader(tt.input)
			_, err := p.Parse(reader)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}

		})
	}
}
