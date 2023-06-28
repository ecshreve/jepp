# [jepp](https://jepp.app)

API fun with Jeopardy! Access >100k Jeopardy clues scraped from [j-archive] via a simple api.


[![CI](https://github.com/ecshreve/jepp/actions/workflows/ci.yml/badge.svg?branch=main&event=push)](https://github.com/ecshreve/jepp/actions/workflows/ci.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ecshreve/jepp)
[![Go Report Card](https://goreportcard.com/badge/github.com/ecshreve/jepp)](https://goreportcard.com/report/github.com/ecshreve/jepp)
[![GoDoc](https://godoc.org/github.com/ecshreve/jepp?status.svg)](https://godoc.org/github.com/ecshreve/jepp)
![GitHub release (release name instead of tag name)](https://img.shields.io/github/v/release/ecshreve/jepp)

---

![jepp](static/repo/jepp-ui.png)

# API

The api is backed by a go web server built with [gin] that exposes a few endpoints to access historical Jeopardy data.

## Types

The shape of the data returned from the api aligns with the db schema, this is accomplished via various struct tags on the type definitions.

### Struct tags quick reference

- `db` tag is used by the [sqlx] library to map the db columns to the struct fields
- `json` tag is used by the [gin] library to map the struct fields to the json response
- `example` tag is used by the [swaggo] library to generate example responses for the swagger docs
- `form` and `binding` tags are used by [gin] to map query arguments to a struct with some basic validation


for example, the `pkg.models.Clue` type is defined as follows:
```{golang}
type Clue struct {
	ClueID     int64  `db:"clue_id" json:"clueId" example:"804002032"`
	GameID     int64  `db:"game_id" json:"gameId" example:"8040"`
	CategoryID int64  `db:"category_id" json:"categoryId" example:"804092001"`
	Question   string `db:"question" json:"question" example:"This is the question."`
	Answer     string `db:"answer" json:"answer" example:"This is the answer."`
}
```


- Struct tags also appear on some helper structs like the `pkg.server.Filter` type:
```{golang}
// Filter describes the query parameters that can be used to filter the results of an API query.
type Filter struct {
	Random *bool  `form:"random"`
	ID     *int64 `form:"id"`
	Page   *int64 `form:"page,default=0" binding:"min=0"`
	Limit  *int64 `form:"limit,default=1" binding:"min=1,max=100"`
}
```


## Frontend / UI

- The ui is served from the `/` endpoint and is an html template that displays the swagger docs, some
	general information, and a sample request.
- The embedded swagger ui provides runnable request / response examples and type references.

## Swagger Docs

- Swagger documentation is generated with [swaggo] and embedded in the homepage as part of the html template.
- Figuring out the right build/deploy configuration was challenging here, I ran into some problems in my Taskfile task dependencies. The main problem seemed to be multiple tasks with the same set of files listed as `sources` causing a watched task to continuously rebuild because of some circular dependencies.
- These problems seem to be solved after breaking up and organizing the taskfiles better.

# DB

Currently the app uses a file based sqlite database. Below are some notes on the deprecated mysql setup.
All in all, the 15 seasons of data currently in the DB only end up as ~25 MB .sql file. Using
sqlite removed the need to run a mysql server and made the app easier to deploy and test.

## Notes on deprecated mysql setup

Getting the data into the database started as a manual process, and hasn't been automated yet because the data is all there and I haven't needed to import / export it recently.

Here's how I went about doing it initially:
- For local development I set the `DB_HOST`, `DB_USER`, `DB_PASS`, `DB_NAME` environment variables to target a `mariadb/mysql` server running in my home lab.
- Most of the time I play with that local copy of the data, but the public api uses a mysql db hosted on [digital ocean](https://www.digitalocean.com/products/managed-databases-mysql)
- Initially to populate the prod db I just manually created a backup of my local database and restored it to the prod database, both via an [adminer](https://hub.docker.com/_/adminer/) instance running in my home lab.
- Currently the `task sql:dump` command will create a dump of the database defined by the environment variables and write it to `data/dump.sql.gz`.
- Recent dumps of the prod database are available in the [data](data/) directory or as downloads on repository's [Releases](https://github.com/ecshreve/jepp/releases) page.



## Data Scraping

note: all the scraping was done against the mysql databse, not the current sqlite setup (though I did 
some brief testing and things seemed to still work for the most part _ymmv_)

The [scraper](pkg/scraper/) package contains the code to scrape [j-archive] for jeopardy clues and write the data to a mysql database. [Colly] is the package use to scrape the data and [sqlx] is used to write the data to the db. The scraping happened in a few passes, more or less following these steps:

Get all the seasons and populate the seasons table.

- This scrape targeted the season [summary page on j-archive](https://www.j-archive.com/listseasons.php) and pulled the season number, start date, end date for each season

Get all the games for each season and populate the game table.

- This scrape targets the individual [season show pages on j-archive](https://www.j-archive.com/showseason.php?season=1) and pulls the game number, air date, taped date for each season
 
Get all the clues for each game in each season and populate the category and clue tables

- This scrape targeted the individual [game pages on j-archive](https://www.j-archive.com/showgame.php?game_id=7040) and pulls the clue data from the `<table>` elements on the page


## references / prior art

- [jservice](https://jservice.io/)
- [jservice repo](https://github.com/sottenad/jService)
- [jeppy](https://github.com/ecshreve/jeppy)
- [illustrated sqlx](https://jmoiron.github.io/sqlx/)

[sqlx]: <https://github.com/jmoiron/sqlx>
[gin]: <https://github.com/gin-gonic/gin>
[swaggo]: <https://github.com/swaggo/swag>
[j-archive]: <https://www.j-archive.com/>
[colly]: <https://github.com/gocolly/colly>


<hr>
<hr>

![cf](https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white)
![do](https://img.shields.io/badge/Digital_Ocean-0080FF?style=for-the-badge&logo=DigitalOcean&logoColor=white)
![ga](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)
![mysql](https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white)
![mariadb](https://img.shields.io/badge/MariaDB-003545?style=for-the-badge&logo=mariadb&logoColor=white)
![dock](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
![swag](https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=Swagger&logoColor=white)
![golan](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)


<a href="https://www.buymeacoffee.com/ecshreve" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-blue.png" alt="Buy Me A Coffee" style="height: 25px !important;width: 100px !important;" ></a>
