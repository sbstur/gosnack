package main

import "fmt"

var TennisScore = [5]string{"Love", "Fifteen", "Thirty", "Forty", "Advantage"}

type Game struct {
	playerName  [2]string
	playerScore [2]int
	game        string
}

func (g *Game) PlayerGetPoint(p int) {
	var aPlayer, bPlayer int = 1, 0
	if p == 0 {
		aPlayer, bPlayer = 0, 1
	}
	if g.playerScore[aPlayer] == 3 {
		if g.playerScore[bPlayer] == 4 {
			g.playerScore[bPlayer] = 3
		} else if g.playerScore[bPlayer] < 3 {
			g.game = g.playerName[aPlayer]
		} else {
			g.playerScore[aPlayer]++
		}
	} else if g.playerScore[aPlayer] == 4 {
		g.game = g.playerName[aPlayer]
	} else {
		g.playerScore[aPlayer]++
	}
}

func (g Game) CurrentScore(server int) {
	var aPlayer, bPlayer int = 1, 0
	if server == 0 {
		aPlayer, bPlayer = 0, 1
	}
	if g.game != "" {
		fmt.Println("=== Game " + g.game)
	} else if g.playerScore[aPlayer] == 3 && g.playerScore[bPlayer] == 3 {
		fmt.Println("Deuce")
	} else if g.playerScore[aPlayer] == 4 {
		fmt.Println(TennisScore[g.playerScore[aPlayer]] + " " + g.playerName[aPlayer])
	} else if g.playerScore[bPlayer] == 4 {
		fmt.Println(TennisScore[g.playerScore[bPlayer]] + " " + g.playerName[bPlayer])
	} else if g.playerScore[aPlayer] == g.playerScore[bPlayer] {
		fmt.Println(TennisScore[g.playerScore[aPlayer]] + " all")
	} else {
		fmt.Println(TennisScore[g.playerScore[aPlayer]] + "-" + TennisScore[g.playerScore[bPlayer]])
	}
}

func main() {
	scoring := [][]int{{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1}, {0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0}, {1, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0}}

	for i := 0; i < len(scoring); i++ {
		gp := &Game{[2]string{"McEnroe", "Borg"}, [2]int{0, 0}, ""}
		for j := 0; gp.game == "" && j < len(scoring[i]); j++ {
			gp.PlayerGetPoint(scoring[i][j])
			gp.CurrentScore(0)
		}
	}
}
