package painter

import (
	"golang.org/x/exp/shiny/screen"
	"testing"
)

type MockOperation1 struct{}

func (mk1 MockOperation1) Do(t screen.Texture) (ready bool) {
	return false
}

type MockOperation2 struct{}

func (mk1 MockOperation2) Do(t screen.Texture) (ready bool) {
	return false
}

func TestMessageQueue(t *testing.T) {

	mq := &messageQueue{}
	if !mq.empty() {
		t.Errorf("expected empty message queue, got non-empty")
	}

	op1 := new(MockOperation1)
	mq.push(op1)
	if mq.empty() {
		t.Errorf("expected 1 operation in message queue, got empty")
	}
	_ = mq.pull()

	if !mq.empty() {
		t.Errorf("expected empty message queue, got non-empty")
	}

	op2 := new(MockOperation1)
	mq.push(op1)
	mq.push(op2)
	if len(mq.ops) != 2 {
		t.Errorf("expected 2 operation in message queue, got other number")
	}

	pushedOp1 := mq.pull()
	pushedOp2 := mq.pull()
	if pushedOp1 != pushedOp2 {
		t.Errorf("exepced equal operations pulled from message queue, got modified")
	}

	op3 := new(MockOperation2)
	mq.push(op1)
	mq.push(op2)
	mq.push(op3)
	if len(mq.ops) != 3 {
		t.Errorf("expected 3 operation in message queue, got other number")
	}

	pushedOp1 = mq.pull()
	pushedOp2 = mq.pull()
	if len(mq.ops) != 1 {
		t.Errorf("expected 1 operation in message queue, got other number")
	}

	pushedOp3 := mq.pull()
	if !mq.empty() {
		t.Errorf("expected empty message queue, got non-empty")
	}
	if pushedOp2 == pushedOp3 {
		t.Errorf("expected different operations pulled from message queue, got equal")
	}
}
