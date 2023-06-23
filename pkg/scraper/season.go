package scraper

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"time"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

var SeasonRE = regexp.MustCompile(`Season ([0-9]+)`)
var SeasonDateRE = regexp.MustCompile(`([0-9]{4}-[0-9]{2}-[0-9]{2}) to ([0-9]{4}-[0-9]{2}-[0-9]{2})`)

func scrapeSeasons() ([]*mods.Season, error) {
	seasonURLs := map[string]string{}
	seasonDates := map[int64][]string{}

	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)

	c.OnHTML("div#content tr", func(e *colly.HTMLElement) {
		ss := ""
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			seasonNum := SeasonRE.FindStringSubmatch(el.Text)
			if seasonNum == nil {
				return
			}
			ss = seasonNum[1]
			seasonURLs[ss] = el.Request.AbsoluteURL(el.Attr("href"))
		})

		ssNum, _ := strconv.Atoi(ss)
		e.ForEach("td", func(_ int, el *colly.HTMLElement) {
			seasonDate := SeasonDateRE.FindStringSubmatch(el.Text)
			if seasonDate == nil {
				return
			}
			seasonDates[int64(ssNum)] = []string{seasonDate[1], seasonDate[2]}
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit("https://www.j-archive.com/listseasons.php")

	seasons := []*mods.Season{}
	for seasonNum, dates := range seasonDates {
		st, _ := time.Parse("2006-01-02", dates[0])
		et, _ := time.Parse("2006-01-02", dates[1])

		seasons = append(seasons, &mods.Season{
			SeasonID:  seasonNum,
			StartDate: st,
			EndDate:   et,
		})
	}

	sort.Slice(seasons, func(i, j int) bool {
		return seasons[i].SeasonID < seasons[j].SeasonID
	})

	return seasons, nil
}

func fillSeasons(db *mods.JeppDB) {
	seasons, err := scrapeSeasons()
	if err != nil {
		panic(err)
	}

	for _, season := range seasons {
		if err := mods.InsertSeason(season); err != nil {
			log.Fatal(err)
		}
	}
}

func scrapeSeasonGames(seasonID int64) ([]*mods.Game, error) {
	var gameRE = regexp.MustCompile(`game_id=([0-9]+)`)
	var metaRE = regexp.MustCompile(`#([0-9]+),.*aired.*([0-9]{4}-[0-9]{2}-[0-9]{2})`)
	var tapedRE = regexp.MustCompile(`Taped.*([0-9]{4}-[0-9]{2}-[0-9]{2})`)

	gameIDs := []int64{}
	showNums := map[int64]int64{}
	airedDates := map[int64]time.Time{}
	tapedDates := map[int64]time.Time{}

	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)

	c.OnHTML("div#content tr", func(e *colly.HTMLElement) {
		var gid int64
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			gameID := gameRE.FindStringSubmatch(el.Attr("href"))
			if gameID == nil {
				return
			}
			gid, _ = strconv.ParseInt(gameID[1], 10, 64)
			gameIDs = append(gameIDs, gid)

			taped := tapedRE.FindStringSubmatch(el.Attr("title"))
			if taped == nil {
				return
			}
			tapedDates[gid], _ = time.Parse("2006-01-02", taped[1])
		})

		e.ForEach("td", func(_ int, el *colly.HTMLElement) {
			meta := metaRE.FindStringSubmatch(el.Text)
			if meta == nil {
				return
			}
			showNums[gid], _ = strconv.ParseInt(meta[1], 10, 64)
			airedDates[gid], _ = time.Parse("2006-01-02", meta[2])
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit(fmt.Sprintf("https://www.j-archive.com/showseason.php?season=%d", seasonID))

	games := []*mods.Game{}
	for _, gid := range gameIDs {
		games = append(games, &mods.Game{
			GameID:    gid,
			SeasonID:  seasonID,
			ShowNum:   showNums[gid],
			GameDate:  airedDates[gid],
			TapedDate: tapedDates[gid],
		})
	}

	sort.Slice(games, func(i, j int) bool {
		return games[i].ShowNum < games[j].ShowNum
	})

	return games, nil
}

func fillSeasonGames(db *mods.JeppDB, seasonID int64) {
	games, err := scrapeSeasonGames(seasonID)
	if err != nil {
		log.Fatal(err)
	}

	for _, game := range games {
		if err := mods.InsertGame(game); err != nil {
			log.Fatal(err)
		}
	}

	log.Infof("Inserted %d games for season %d", len(games), seasonID)
}
