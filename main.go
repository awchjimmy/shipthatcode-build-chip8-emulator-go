package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO (what-is-chip8): implement per the lesson description.

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		DecodeOpcode(line)
	}
	err := sc.Err()
	if err != nil {
		log.Println("Error in scanner.", err)
	}
}

func MemoryLayout(hexAddress string) {
	address, err := strconv.ParseInt(hexAddress, 16, 32)
	if err != nil {
		log.Println("Error could not convert hex address.", err)
		return
	}
	if address >= 0x000 && address <= 0x1FF {
		fmt.Println("INTERPRETER")
	} else if address >= 0x200 && address <= 0xFFF {
		fmt.Println("PROGRAM")
	} else {
		fmt.Println("INVALID")
	}
}

func VFCarry(line string) {
	numbers := strings.Split(line, " ")
	VX, err := strconv.ParseInt(numbers[0], 10, 32)
	if err != nil {
		log.Println("Error converting int.", err)
	}
	VY, err := strconv.ParseInt(numbers[1], 10, 32)
	if err != nil {
		log.Println("Error converting int.", err)
	}
	VF := 0
	if VX+VY >= 256 {
		VF = 1
	}
	fmt.Printf("%d %d\n", (VX+VY)%256, VF)
}

func DecodeOpcode(line string) {
	opcode, err := strconv.ParseInt(line, 16, 32)
	if err != nil {
		log.Println("Error could not decode opcode.", err)
	}

	switch Decode(int(opcode)).High4 {
	case 0x0:
		switch opcode {
		case 0x00E0:
			fmt.Println("CLS")
		case 0x00EE:
			fmt.Println("RET")
		default:
			fmt.Println("UNKNOWN")
		}
	case 0x1:
		fmt.Println("JP")
	case 0x2:
		fmt.Println("CALL")
	case 0x3:
		fmt.Println("SE_IMM")
	case 0x4:
		fmt.Println("SNE_IMM")
	case 0x5:
		fmt.Println("SE_REG")
	case 0x6:
		fmt.Println("LD_IMM")
	case 0x7:
		fmt.Println("ADD_IMM")
	case 0x8:
		fmt.Println("ALU")
	case 0x9:
		fmt.Println("SNE_REG")
	case 0xA:
		fmt.Println("LD_I")
	case 0xB:
		fmt.Println("JP_V0")
	case 0xC:
		fmt.Println("RND")
	case 0xD:
		fmt.Println("DRW")
	case 0xE:
		switch Decode(int(opcode)).KK {
		case 0x9E:
			fmt.Println("SKP")
		case 0xA1:
			fmt.Println("SKNP")
		default:
			fmt.Println("UNKNOWN")
		}
	case 0xF:
		fmt.Println("FX")
	default:
		fmt.Println("UNKNOWN")
	}
}

type DecodeHelper struct {
	High4 int
	X     int
	Y     int
	N     int
	KK    int
	NNN   int
}

func Decode(opcode int) DecodeHelper {
	return DecodeHelper{
		High4: (opcode & 0xF000) >> 12,
		X:     (opcode & 0x0F00) >> 8,
		Y:     (opcode & 0x00F0) >> 4,
		N:     opcode & 0x000F,
		KK:    opcode & 0x00FF,
		NNN:   opcode & 0x0FFF,
	}
}
