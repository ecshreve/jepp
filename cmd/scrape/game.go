package main

import (
	"fmt"

	mod "github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

// scrapeGame scrapes a game from j-archive.com
func scrapeGameClues(gameID int64) (map[int64]*mod.Clue, map[int64]string) {
	clueMap := map[int64]*mod.Clue{}
	clueStrings := map[int64]string{}
	cats := map[mod.Round][]string{}

	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)

	// collect and parse the clues
	c.OnHTML("td.clue", func(e *colly.HTMLElement) {
		cid := e.ChildAttr("td.clue_text", "id")
		if cid == "" {
			return
		}

		clueText := e.ChildText(fmt.Sprintf("td#%s", cid))
		clueAnswer := e.ChildText(fmt.Sprintf("td#%s_r em.correct_response", cid))
		clueId := mod.ParseClueID(cid, gameID)

		clueMap[clueId] = &mod.Clue{ClueID: clueId, GameID: gameID, Question: clueText, Answer: clueAnswer}
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

	catMap := map[int64]string{}

	for clueId, clueStr := range clueStrings {
		rd, col := utils.ParseRoundAndColumn(clueStr)
		catName := cats[mod.Round(rd)][col-1]
		catMap[clueId] = catName
	}

	return clueMap, catMap
}

func scrapeAndFillCluesForGame(db *mod.JeppDB, gid int64) int {
	clues, cats := scrapeGameClues(gid)

	for clueID, clue := range clues {
		actual, _ := db.GetCategoryByName(cats[clueID])
		if actual != nil {
			clue.CategoryID = actual.CategoryID
			continue
		}

		inserted, err := db.InsertCategory(cats[clueID])
		if err != nil {
			log.Fatal(err)
		}
		clue.CategoryID = inserted.CategoryID
	}

	for _, clue := range clues {
		if err := db.InsertClue(clue); err != nil {
			log.Fatal(err)
		}
	}

	log.Infof("inserted %d clues for game %d", len(clues), gid)
	return len(clues)
}
