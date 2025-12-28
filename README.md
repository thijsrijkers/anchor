# Anchor

Anchor is a lightweight scripting language that compiles directly to (currently) ARM machine code. It is designed for simplicity and performance, allowing scripts to execute as native binaries without any runtime or virtual machine.

## Overview

Anchor bridges the gap between scripting and systems programming by providing a minimalistic language that maps directly to raw machine instructions. It is well-suited for educational purposes, experimentation with compiler design, and low-level systems programming.

## Goals

- Compile scripts directly to ARM machine code
- Eliminate external runtimes and dependencies
- Provide a minimal, readable syntax for low-level operations
- Offer a scripting experience with the performance of native execution

## Implementation

Anchor is implemented in Go and processes `.anchor` source files. The compiler reads source code, parses it, and emits assembly instructions targeting the (currently) ARM architecture. The output is a standalone executable file that can be run directly on Linux systems.

## Features

- Register-level arithmetic and data movement
- Direct system call access (e.g., write, exit)
- Minimal and predictable execution model
- Executable output without linking or intermediate formats

## Getting Started

### Requirements
- **Operating System:** Linux (currently only Linux is supported for generated assembly)  
- **Tools:**  
  - [Go](https://go.dev/) — to run the generator  
  - **GCC** — to assemble and link the generated assembly  

---

### Debugging
To see the calculated values from an `.anchor` file, use the `--debug` flag:

```bash
go run ./main.go ./fixtures/additionOperation.anchor --debug
```

This will print the intermediate and final computed values without executing the assembly.


### Execute Code
To compile and run the generated assembly, add the `--execute` flag:

```bash
go run ./main.go ./fixtures/additionOperation.anchor --debug --execute
```

After running the command, the following steps are performed:

1. **Generate Assembly**  
   The generator converts your `.anchor` file into a GAS/AT&T `.s` assembly file.

2. **Compile Assembly**  
   GCC assembles the `.s` file and links it into a Linux ELF executable.

3. **Run Executable**  
   The generated executable is executed, and any output produced by your program (e.g., via syscalls) is printed to the console.

4. **Test ARM**
   If you are in a situation where you cannot run ARM code on your device, you can visit https://godbolt.org/ and select **AArch64 binutils** to verify that the assembly is valid.
