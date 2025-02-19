# GPTChess API

GPTChess is a Go package that integrates OpenAI's GPT-4o to validate and generate legitimate chess moves. This API allows users to play chess with an AI that ensures all moves are legal while responding with appropriate counter-moves.

## Features
- Uses OpenAI GPT-4o to validate chess moves.
- Generates AI-driven responses for legal chess play.
- Provides a simple interface for handling user input.
- Displays the chessboard after each move.

## Installation
To use GPTChess, install the package using:
```sh
 go get github.com/cprofessional/gptchess
```

## Usage
Below is an example implementation of GPTChess in a Go program:

```go
package main

import (
	"fmt"

	"github.com/cprofessional/gptchess"
)

func main() {
	gptchess.Setup("OPENAI_API_KEY")
	r := gptchess.NewRound()

	for r.Running {
		from, to := handleInput()
		m, err := r.MakeMove(from, to)
		if err != nil {
			fmt.Println(err)
			continue
		}

		r.ApplyMove(m)
		r.ApplyMove(r.GenerateMove())

		for _, row := range r.Chessboard {
			fmt.Println(row)
		}
	}
}

func handleInput() (string, string) {
	var from string
	var to string

	fmt.Print("\nChoose a location to lift your piece\n$ ")
	fmt.Scanln(&from)

	fmt.Print("\nChoose a location to set your piece\n$ ")
	fmt.Scanln(&to)

	return from, to
}
```

## Setup
1. Obtain an OpenAI API Key from [OpenAI](https://openai.com).
2. Set up the API key in your program using:
   ```go
   gptchess.Setup("YOUR_OPENAI_API_KEY")
   ```

## How It Works
1. The user inputs a move by selecting a piece's starting and ending positions.
2. The move is validated using GPT-4o.
3. If the move is legal, it is applied to the chessboard.
4. The AI generates a counter-move and applies it.
5. The updated chessboard is printed to the console.
6. The game continues until a checkmate, draw, or termination condition.

## Requirements
- Go 1.18+
- An OpenAI API Key

