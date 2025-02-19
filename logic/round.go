package logic

import "fmt"

type Round struct {
        Running bool

        Chessboard [8][8]string

        White *Player
        Black *Player

        LatestMove *Move
}

type Player struct {
        Color string

        Captures []string
        Moves    []*Move
}

func (r *Round) ApplyMove(m *Move) {
        m.Player.Moves = append(m.Player.Moves, m)

        if r.Chessboard[8-int(m.From[1]-'0')][int(m.From[0]-'a')] != "" {
                m.Player.Captures = append(m.Player.Captures, r.Chessboard[8-int(m.From[1]-'0')][int(m.From[0]-'a')])
        }

        t := r.Chessboard[8-int(m.From[1]-'0')][int(m.From[0]-'a')]
        r.Chessboard[8-int(m.From[1]-'0')][int(m.From[0]-'a')] = ""
        r.Chessboard[8-int(m.To[1]-'0')][int(m.To[0]-'a')] = t
}

func (r *Round) toString(m *Move) string {
        return "Current Board: {" + fmt.Sprintf("%v", r.Chessboard) + "} Move: {" + m.From + "to" + m.To + "}"
}

func (r *Round) ChessboardToString() string {
        return fmt.Sprintf("%v", r.Chessboard)
}