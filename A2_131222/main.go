package main

import (
	"os"
	"strconv"
)

func main() {
	s := makeString(1000, 1)
	saveToFile(s, "cFiles/program1000.c")
	s = makeString(100000, 100)
	saveToFile(s, "cFiles/program100000.c")
	s = makeString(10000000, 10000)
	saveToFile(s, "cFiles/program10000000.c")
	s = makeString(2000, 2)
	saveToFile(s, "cFiles/program2000.c")
	println("CREATED C-FILES")
}
func makeString(endRange int, increment int) string {
	var s string
	s = "#include <stdio.h>\n"
	s = s + "\nvoid p(int x) {\n  switch (x) {\n"
	for i := 0; i < endRange; i = i + increment {
		s = s + "    case " + strconv.Itoa(i) + ": q" + strconv.Itoa(i) + "(); break;\n"
	}
	s = s + "  }\n}"
	return s
}
func saveToFile(s string, name string) {
	os.Create(name)
	os.WriteFile(name, []byte(s), 0644)
}
