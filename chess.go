package gptchess

import (
        "github.com/cprofessional/gptchess/logic"
        "github.com/cprofessional/gptchess/openai"
)

func Setup(k string) {
        openai.SetAPIKey(k)
}

func NewRound() *logic.Round {
        return &logic.Round{
                Running: true,

                Chessboard: [8][8]string{
                        {"Br", "Bn", "Bb", "Bq", "Bk", "Bb", "Bn", "Br"},
                        {"Bp", "Bp", "Bp", "Bp", "Bp", "Bp", "Bp", "Bp"},
                        {"", "", "", "", "", "", "", ""},
                        {"", "", "", "", "", "", "", ""},
                        {"", "", "", "", "", "", "", ""},
                        {"", "", "", "", "", "", "", ""},
                        {"Wp", "Wp", "Wp", "Wp", "Wp", "Wp", "Wp", "Wp"},
                        {"Wr", "Wn", "Wb", "Wq", "Wk", "Wb", "Wn", "Wr"},
                },

                White: &logic.Player{
                        Color: "white",
                },
                Black: &logic.Player{
                        Color: "black",
                },
        }
}