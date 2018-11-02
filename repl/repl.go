package repl

import (
	"bufio"
	"fmt"
	"io"
	"writingincompiler/compiler"
	"writingincompiler/vm"
	"writinginterpreter/lexer"
	"writinginterpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			// printParserErrors(out, p.Errors())
			// continue
		}
		comp := compiler.New()
		err := comp.Compile(program)

		machine := vm.New(comp.Bytecode())
		err = machine.Run()

		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.LastPoppedStackElem()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")

	}
}
