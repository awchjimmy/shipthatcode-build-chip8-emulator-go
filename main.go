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
		VFCarry(line)
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
