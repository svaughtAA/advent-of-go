package day8

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        8,
		Part:       1,
		Calculator: pt1,
	}
}

type Interpreter struct {
	acc int
	prm []Instruction
	pc  int
	ex  map[int]bool
	err error
}

type Instruction struct {
	name  string
	value int
}

func ParseInput(input string) []Instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, len(lines))
	instructionRe := regexp.MustCompile(`^(?P<name>\w{3}) (?P<sgn>\+|\-)(?P<dgt>\d{1,})$`)
	for _, line := range lines {
		matches := instructionRe.FindStringSubmatch(line)
		name := matches[instructionRe.SubexpIndex("name")]
		sgn := matches[instructionRe.SubexpIndex("sgn")]
		dgt := matches[instructionRe.SubexpIndex("dgt")]
		value, _ := strconv.Atoi(dgt)
		if sgn == "-" {
			value *= -1
		}
		instructions = append(instructions, Instruction{
			name:  name,
			value: value,
		})
	}
	return instructions
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s-%d", i.name, i.value)
}

func (int *Interpreter) Step() {
	inst := int.prm[int.pc]
	visited := int.ex[int.pc]
	if visited {
		int.err = fmt.Errorf("CYCLE DETECTED")
		return
	} else {
		int.ex[int.pc] = true
	}

	switch inst.name {
	case "nop":
		int.pc++
		return
	case "acc":
		int.pc++
		int.acc += inst.value
		return
	case "jmp":
		int.pc += inst.value
		return
	}
}

func pt1(input string) (string, error) {
	interpreter := &Interpreter{
		prm: ParseInput(input),
		ex:  map[int]bool{},
	}
	for interpreter.err == nil {
		interpreter.Step()
	}
	fmt.Printf("acc %v\n", interpreter.acc)
	return "", errors.New("Unimplemented!")
}
