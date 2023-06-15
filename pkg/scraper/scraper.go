package scraper

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gocolly/colly/v2"
)

var re = regexp.MustCompile(`.*#([0-9]+) - (.*)$`)

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
	"TB": FinalJeopardy,
}

func ScrapeMany(gameIDs []int64) ([]models.Game, []models.Clue) {
	games := []models.Game{}
	clues := []models.Clue{}
	for _, gameID := range gameIDs {
		g, c := Scrape(gameID)
		games = append(games, g)
		clues = append(clues, c...)
	}

	return games, clues
}

func Scrape(gameID int64) (models.Game, []models.Clue) {
	var showNum int64
	var gameDate time.Time
	clues := []models.Clue{}
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

		gd, err := time.Parse(models.TIME_FORMAT, tokens[2])
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
		clues = append(clues, models.Clue{ClueID: cid, GameID: gameID, Question: clueText, Answer: clueAnswer})
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

	g := models.Game{
		GameID:   gameID,
		ShowNum:  showNum,
		GameDate: gameDate,
	}

	return g, clues
}

func helper(s string, cats map[Round][]string) string {
	tokens := strings.Split(s, "_")
	if len(tokens) > 4 || len(tokens) == 0 {
		fmt.Println("Error parsing clueID")
		return ""
	}

	round := roundMap[tokens[1]]
	if round == FinalJeopardy {
		if tokens[1] == "TB" {
			return cats[round][1]
		}
		return cats[round][0]
	}

	catIndex, _ := strconv.Atoi(tokens[2])
	return cats[round][catIndex-1]
}
