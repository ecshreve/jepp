package scraper

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	mod "github.com/ecshreve/jepp/pkg/models"
	"github.com/gocolly/colly/v2"
)

var re = regexp.MustCompile(`.*#([0-9]+) - (.*)$`)

func ScrapeMany(gameIDs []int64) ([]mod.Game, []mod.Clue, []mod.Category) {
	games := []mod.Game{}
	clues := []mod.Clue{}
	cats := []mod.Category{}

	for _, gameID := range gameIDs {
		g, c, cc := Scrape(gameID)
		games = append(games, g)
		clues = append(clues, c...)
		cats = append(cats, cc...)
	}

	return games, clues, cats
}

func Scrape(gameID int64) (mod.Game, []mod.Clue, []mod.Category) {
	var showNum int64
	var gameDate time.Time
	clueMap := map[int64]mod.Clue{}
	clueStrings := map[int64]string{}
	cats := map[mod.Round][]string{}
	cc := []string{}

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

		gd, err := time.Parse(mod.TIME_FORMAT, tokens[2])
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
		clueId := mod.GetClueID(cid, gameID)

		clueMap[clueId] = mod.Clue{ClueID: clueId, GameID: gameID, Question: clueText, Answer: clueAnswer}
		clueStrings[clueId] = cid
	})

	c.OnHTML("div[id=jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[mod.Jeopardy] = append(cats[mod.Jeopardy], cc...)
	})

	c.OnHTML("div[id=double_jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[mod.DoubleJeopardy] = append(cats[mod.DoubleJeopardy], cc...)
	})

	c.OnHTML("div[id=final_jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[mod.FinalJeopardy] = append(cats[mod.FinalJeopardy], cc...)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit(fmt.Sprintf("https://www.j-archive.com/showgame.php?game_id=%d", gameID))

	clues := []mod.Clue{}
	for clueId, clue := range clueMap {
		clue.Category = helper(clueStrings[clueId], cats)
		clues = append(clues, clue)
	}

	cc = append(cc, cats[mod.Jeopardy]...)
	cc = append(cc, cats[mod.DoubleJeopardy]...)
	cc = append(cc, cats[mod.FinalJeopardy]...)

	allCats := []mod.Category{}
	for _, cat := range cc {
		allCats = append(allCats, mod.Category{Name: cat, GameID: gameID})
	}

	g := mod.Game{
		GameID:   gameID,
		ShowNum:  showNum,
		GameDate: gameDate,
	}

	return g, clues, allCats
}

func helper(s string, cats map[mod.Round][]string) string {
	tokens := strings.Split(s, "_")
	if len(tokens) > 4 || len(tokens) == 0 {
		fmt.Println("Error parsing clueID")
		return ""
	}

	round := mod.RoundMap[tokens[1]]
	if round == mod.FinalJeopardy {
		if tokens[1] == "TB" {
			return cats[round][1]
		}
		return cats[round][0]
	}

	catIndex, _ := strconv.Atoi(tokens[2])
	return cats[round][catIndex-1]
}
