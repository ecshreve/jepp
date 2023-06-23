package scraper

import (
	"fmt"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

// scrapeGameClues scrapes a game from j-archive.com.
func scrapeGameClues(gameID int64) (map[int64]*mods.Clue, map[int64]string) {
	clueMap := map[int64]*mods.Clue{}
	clueStrings := map[int64]string{}
	cats := map[mods.Round][]string{}

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
		clueId := mods.ParseClueID(cid, gameID)

		clueMap[clueId] = &mods.Clue{ClueID: clueId, GameID: gameID, Question: clueText, Answer: clueAnswer}
		clueStrings[clueId] = cid
	})

	// collect and parse the categories for single jepp
	c.OnHTML("div[id=jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[mods.Jeopardy] = append(cats[mods.Jeopardy], cc...)
	})

	// collect and parse the categories for double jepp
	c.OnHTML("div[id=double_jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[mods.DoubleJeopardy] = append(cats[mods.DoubleJeopardy], cc...)
	})

	// collect and parse the categories for final jepp
	c.OnHTML("div[id=final_jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[mods.FinalJeopardy] = append(cats[mods.FinalJeopardy], cc...)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit(fmt.Sprintf("https://www.j-archive.com/showgame.php?game_id=%d", gameID))

	catMap := map[int64]string{}

	for clueId, clueStr := range clueStrings {
		rd, col := utils.ParseRoundAndColumn(clueStr)
		catName := cats[mods.Round(rd)][col-1]
		catMap[clueId] = catName
	}

	return clueMap, catMap
}

func scrapeAndFillCluesForGame(db *mods.JeppDB, gid int64) int {
	clues, cats := scrapeGameClues(gid)

	for clueID, clue := range clues {
		actual, err := mods.GetCategoryByName(cats[clueID])
		if actual != nil {
			clue.CategoryID = actual.CategoryID
			continue
		}

		inserted, err := mods.InsertCategory(cats[clueID])
		if err != nil {
			log.Fatal(err)
		}
		clue.CategoryID = inserted.CategoryID
	}

	for _, clue := range clues {
		if err := mods.InsertClue(clue); err != nil {
			log.Fatal(err)
		}
	}

	log.Infof("inserted %d clues for game %d", len(clues), gid)
	return len(clues)
}
