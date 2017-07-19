package main

import "fmt"

type Game struct {
    Title string
    Year int
    Genre []string
}

func main() {

    fallout4 := Game{
        Title: "Fallout 4",
        Year: 2015,
    }
    fallout4.Genre = []string{
        "RPG", "Adventure",
    }

    fmt.Println(fallout4.DisplayTitle())
}

// Method for Game type. Game here is the reciever type
func (game Game) DisplayTitle() string {

    return fmt.Sprintf("%s (%d) genres %s %s", game.Title, game.Year, game.Genre[0], game.Genre[1])
}
