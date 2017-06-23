package foundation

import (
	"bufio"
	"os"
	"strings"
)

const (
	IArgActionNext        = IArgAction(0)
	IArgActionStopInput   = IArgAction(1)
	IArgActionRepet       = IArgAction(2)
	IArgActionProcessExit = IArgAction(9)
)

type IArgAction int8

type IArag struct {
	args []*tArgDes
	idx  int
}

type tArgDes struct {
	alert   string
	inputer func(arg string) (IArgAction)
}

func (s *IArag) Input(alert string, inputer func(arg string) (IArgAction)) (*IArag) {
	if inputer == nil {
		return s
	}
	aArgDes := new(tArgDes)
	aArgDes.alert = alert
	aArgDes.inputer = inputer
	if s.args == nil {
		s.args = make([]*tArgDes, 0, 5)
	}
	s.args = append(s.args, aArgDes)
	return s
}

func (s *IArag) Run() {
	if s.idx >= len(s.args) {
		return
	}
	aArgDes := s.args[s.idx]
	print(aArgDes.alert)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadBytes('\n')
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
	inputStr := ""
	if input != nil {
		inputStr = string(input[0: len(input)-1])
	}
	action := aArgDes.inputer(strings.TrimSpace(inputStr))
	switch action {
	case IArgActionProcessExit:
		os.Exit(0)
	case IArgActionNext:
		s.idx ++
	case IArgActionStopInput:
		s.idx = len(s.args)
	}
	s.Run()
}