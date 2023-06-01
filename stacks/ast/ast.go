package ast

import (
	"errors"
	"fmt"

	"github.com/aggronmagi/gstacks/stacks/token"
)

type Grountine struct {
	Status *GStatus
	Stacks []*GStack
}

type GStatus struct {
	Status  string
	Minutes int
}
type GStack struct {
	Func string
	File string
}

func CheckGoroutines(all any) (any, error) {
	return all, nil
}

func NewGoroutines() ([]*Grountine, error) {
	return make([]*Grountine, 0, 1024), nil
}

func AppendGoroutines(s any, item any) ([]*Grountine, error) {
	arr, ok := s.([]*Grountine)
	if !ok {
		return nil, errors.New("invlaid stack arrays")
	}
	v, ok := item.(*Grountine)
	if !ok {
		return nil, errors.New("invlaid, not stack object")
	}
	arr = append(arr, v)
	return arr, nil
}

func NewGoroutine(id, status, stacks any) (*Grountine, error) {
	return &Grountine{
		Status: status.(*GStatus),
		Stacks: stacks.([]*GStack),
	}, nil
}

func NewStatus(num any, args ...any) (s *GStatus, err error) {
	s = &GStatus{}
	switch v := num.(type) {
	case int:
		s.Minutes = v
	case *token.Token:
		n, err := v.Int64Value()
		if err != nil {
			return nil, err
		}
		s.Minutes = int(n)
	default:
		fmt.Printf("%#v\n", num)
		return nil, errors.New("invalid status numbers")
	}
	for k, v := range args {
		if k > 0 {
			s.Status += " "
		}
		s.Status += v.(*token.Token).IDValue()
	}
	return
}

func NewStacks() ([]*GStack, error) {
	return make([]*GStack, 0, 4), nil
}

func AppendStack(s any, item any) (arr []*GStack, err error) {
	arr, ok := s.([]*GStack)
	if !ok {
		return nil, errors.New("invlaid stack arrays")
	}
	v, ok := item.(*GStack)
	if !ok {
		return nil, errors.New("invlaid, not stack object")
	}
	arr = append(arr, v)
	return
}

func NewStack(fun, file any) (stack *GStack, err error) {
	stack = &GStack{
		Func: fun.(string),
		File: file.(*token.Token).IDValue(),
	}
	return
}

func NewCall(items ...any) (fun string, err error) {
	for k, v := range items {
		if k > 0 {
			fun += " "
		}
		fun += v.(*token.Token).IDValue()
	}
	return
}
