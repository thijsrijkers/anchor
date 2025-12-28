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

### Debugging
To see what has been calculated from an `.anchor` file, you can add the `--debug` flag to get the calculated values.