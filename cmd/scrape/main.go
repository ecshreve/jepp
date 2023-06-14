package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/kr/pretty"
)

type Round int // 0 = Jeopardy, 1 = Double Jeopardy, 2 = Final Jeopardy

const (
	Jeopardy Round = iota + 1
	DoubleJeopardy
	FinalJeopardy
)

var roundMap = map[string]Round{
	"J":  Jeopardy,
	"DJ": DoubleJeopardy,
	"FJ": FinalJeopardy,
}

type Game struct {
	GameID   int64
	ShowNum  int64
	GameDate time.Time
	Clues    []Clue
}

type Clue struct {
	GameID   int64
	ClueID   string
	Category string
	Text     string
	Answer   string
}

const timeFormat = "Monday, January 2, 2006"

var re = regexp.MustCompile(`.*#([0-9]+) - (.*)$`)

func (g Game) String() string {
	return fmt.Sprintf("ID: %d\t %s", g.GameID, g.GameDate.Format(timeFormat))
}

func main() {
	gid := int64(6821)
	game := scrape(gid)

	dumpGameToFile(game)
	tmp := dumpFileToGame("game-6821.json")
	pretty.Print(tmp)
}

func scrape(gameID int64) Game {
	var showNum int64
	var gameDate time.Time
	clues := []Clue{}
	cats := map[Round][]string{}

	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)

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
		clues = append(clues, Clue{ClueID: cid, GameID: gameID, Text: clueText, Answer: clueAnswer})
	})

	c.OnHTML("div[id=jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[Jeopardy] = append(cats[Jeopardy], cc...)
	})

	c.OnHTML("div[id=double_jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[DoubleJeopardy] = append(cats[DoubleJeopardy], cc...)
	})

	c.OnHTML("div[id=final_jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[FinalJeopardy] = append(cats[FinalJeopardy], cc...)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit(fmt.Sprintf("https://www.j-archive.com/showgame.php?game_id=%d", gameID))

	for i, clue := range clues {
		clues[i].Category = helper(clue.ClueID, cats)
	}

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

func helper(s string, cats map[Round][]string) string {
	// clue_J_1_1
	tokens := strings.Split(s, "_")
	if len(tokens) > 4 || len(tokens) < 0 {
		fmt.Println("Error parsing clueID")
		return ""
	}

	round := roundMap[tokens[1]]
	if round == FinalJeopardy {
		return cats[round][0]
	}

	catIndex, _ := strconv.Atoi(tokens[2])
	return cats[round][catIndex-1]
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
