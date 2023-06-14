package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/kr/pretty"
)

type Game struct {
	GameID   int64
	ShowNum  int64
	GameDate time.Time
	Clues    []Clue
}

type Clue struct {
	ClueID string
	Text   string
	Answer string
}

const timeFormat = "Monday, January 2, 2006"

var re = regexp.MustCompile(`.*#([0-9]+) - (.*)$`)

func (g Game) String() string {
	return fmt.Sprintf("ID: %d\t %s", g.GameID, g.GameDate.Format(timeFormat))
}

func main() {
	gid := int64(6821)
	game := scrapeGame(gid)
	dumpGameToFile(game)

	tmp := dumpFileToGame("game-6821.json")

	pretty.Print(tmp)
}

func scrapeGame(gameID int64) Game {
	var showNum int64
	var gameDate time.Time
	clues := []Clue{}

	c := colly.NewCollector()

	// collect and parse the gameID and gameDate
	c.OnHTML("div#game_title", func(e *colly.HTMLElement) {
		tokens := re.FindStringSubmatch(e.ChildText("h1"))
		if len(tokens) != 3 {
			log.Fatal("Error parsing gameID and gameDate")
		}

		sn, err := strconv.ParseInt(tokens[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		showNum = sn

		gd, err := time.Parse(timeFormat, tokens[2])
		if err != nil {
			log.Fatal(err)
		}
		gameDate = gd
	})

	// collect and parse the clues
	c.OnHTML("td.clue", func(e *colly.HTMLElement) {
		cid := e.ChildAttr("td.clue_text", "id")
		if cid == "" {
			return
		}

		clueText := e.ChildText(fmt.Sprintf("td#%s", cid))
		clueAnswer := e.ChildText(fmt.Sprintf("td#%s_r em.correct_response", cid))
		clues = append(clues, Clue{ClueID: cid, Text: clueText, Answer: clueAnswer})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit(fmt.Sprintf("https://www.j-archive.com/showgame.php?game_id=%d", gameID))

	return Game{
		GameID:   gameID,
		ShowNum:  showNum,
		GameDate: gameDate,
		Clues:    clues,
	}
}

func dumpGameToFile(game Game) {
	file, _ := json.MarshalIndent(game, "", " ")

	fname := fmt.Sprintf("game-%d.json", game.GameID)
	_ = ioutil.WriteFile(fname, file, 0644)
}

func dumpFileToGame(fname string) Game {
	file, _ := ioutil.ReadFile(fname)

	game := Game{}
	_ = json.Unmarshal([]byte(file), &game)

	return game
}

// On every a element which has href attribute call callback
// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 	link := e.Attr("href")
// 	val := e.Text
// 	if !strings.Contains(val, "next") {
// 		return
// 	}

// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
// 	c.Visit(e.Request.AbsoluteURL(link))
// })
