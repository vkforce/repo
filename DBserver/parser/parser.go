package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() {
	input := "$5\r\nAhmed\r\n"
	reader := bufio.NewReader(strings.NewReader(input))
	b, _ := reader.ReadByte()
	if b != '$' {
		fmt.Println("Error: Input Error")
		os.Exit(1)
	}
	size, _ := reader.ReadByte()
	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	reader.ReadByte()
	reader.ReadByte()

	name := make([]byte, strSize)
	reader.Read(name)
	fmt.Println(string(name))

}
