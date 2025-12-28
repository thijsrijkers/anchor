package runner

import (
	"fmt"
	"os"
	"os/exec"
)

func AssembleAndRun(debug bool) {
	asmPath := "./dist/output.s"
	exePath := "./dist/output"

	compile := exec.Command("gcc", "-nostdlib", "-no-pie", asmPath, "-o", exePath)
	out, err := compile.CombinedOutput()
	if err != nil {
		fmt.Println("Compilation failed:", err)
		fmt.Println(string(out))
		os.Exit(1)
	}

	run := exec.Command(exePath)
	out, err = run.CombinedOutput()
	if err != nil {
		fmt.Println("Execution failed:", err)
		os.Exit(1)
	}

	fmt.Println(string(out))

	if debug {
		fmt.Println("Program output:")
		fmt.Println(string(out))
	}
}
