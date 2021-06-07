package algorithm

import "testing"

func TestDeepFirstSearch(t *testing.T) {
	ResetCondition(3)
	DeepFirstSearch(0)
}

func TestDeepFirstSearch2(t *testing.T) {
	ResetCondition(9)
	DeepFirstSearch2(0)
}

func TestDeepFirstSearch3(t *testing.T) {
	DeepFirstSearch3(0, 0, 0)
}
