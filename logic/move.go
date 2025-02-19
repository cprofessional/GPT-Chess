package logic

import (
        "encoding/json"
        "fmt"

        "github.com/cprofessional/gptchess/openai"
)

type Move struct {
        Player *Player

        From string `json:"from"`
        To   string `json:"to"`
}

func (r *Round) ValidateMove(m *Move) bool {
        botCtx := openai.NewMessage("system", "You are  chess validater that only responds in JSON format ONLY; DO NOT USE CODE WRAPPERS; Example: {\"valid\": true} or {\"valid\": false}")
        userCtx := openai.NewMessage("user", r.toString(m))

        req := openai.NewRequest("gpt-4o", []*openai.Message{botCtx, userCtx})

        return req.Call().Choices[0].Message.Content == "{\"valid\": true}"
}

func (r *Round) ValidateCheck(m *Move) bool {
        botCtx := openai.NewMessage("system", "You are  chess check validater that only responds in JSON format ONLY; DO NOT USE CODE WRAPPERS; Example: {\"check\": true} or {\"check\": false}")
        userCtx := openai.NewMessage("user", r.toString(m))

        req := openai.NewRequest("gpt-4o", []*openai.Message{botCtx, userCtx})

        return req.Call().Choices[0].Message.Content == "{\"check\": true}"
}

func (r *Round) GenerateMove() *Move {
        botCtx := openai.NewMessage("system", "You are chess move generator that only responds in JSON format ONLY; YOU ARE ON THE BLACK TEAM, MEANING YOU CAN ONLY MOVE PEICES THAT START WITH B; DO NOT USE CODE WRAPPERS; Example: {\"from\": \"a2\", \"to\": \"a4\"}")
        userCtx := openai.NewMessage("user", r.ChessboardToString())

        req := openai.NewRequest("gpt-4o", []*openai.Message{botCtx, userCtx})

        var m *Move
        json.Unmarshal([]byte(req.Call().Choices[0].Message.Content), &m)

        m.Player = r.Black

        if r.ValidateMove(m) {
                return m
        } else {
                return r.GenerateMove()
        }
}

func (r *Round) MakeMove(from string, to string) (*Move, error) {
        m := &Move{
                Player: r.White,
                From:   from,
                To:     to,
        }

        if r.ValidateMove(m) {
                return m, nil
        } else {
                return m, fmt.Errorf("invalid move")
        }
}