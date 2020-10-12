package stack

import (
	"testing"
)

func TestPushIsInsert(t *testing.T) {

	var check = []string{"a"}

	res := Push("a")

	if compare(check, res) != true {
		t.Errorf("Getting incorrect, got: %stack.[]string, want %stack.[]string", res, check)
	}
}
func TestPushInFull(t *testing.T) {
	var res []string
	var check = []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}

	for i := 0; i < 11; i++ {
		res = Push("a")
	}
	if compare(check, res) != true {
		t.Errorf("Getting incorrect value, got: %stack.[]string, want %stack.[]string", res, check)

	}
}

func TestPopIsEmpty(t *testing.T) {
	Stack = []string{}
	res := Pop()
	if isEmpty(res) != true {
		t.Errorf("Getting incorrect value, got: %stack.[]string, want %stack.[]string", res, Stack)
	}
}
func TestPopIsDeleted(t *testing.T) {
	Stack = []string{"a", "b", "c"}
	res := Pop()
	if IsDeleted(Stack, res) != true {
		t.Errorf("Getting incorrect value, got: %stack.[]string, want %stack.[]string", res, Stack)
	}
}

func compare(check, res []string) bool {

	if len(check) != len(res) {
		return false
	}
	for i, v := range check {
		if v != res[i] {
			return false
		}

	}
	return true
}

func isEmpty(res []string) bool {
	if len(res) == 0 {
		return true
	}
	return false
}

func IsDeleted(check, res []string) bool {
	l := len(check)
	if len(check[:l]) != len(res) {
		return false
	}
	for i, v := range check {
		if v != res[i] {
			return false
		}
	}
	return true
}
