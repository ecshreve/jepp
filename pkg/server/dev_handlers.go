package server

import (
	"strconv"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
)

// registerDevHandlers registers route handlers for development.
func registerDevHandlers() {
	r := gin.Default()

	r.LoadHTMLGlob("pkg/server/templates/**/*")

	r.GET("/debug/:clueID", DebugUIHandler)
	// s.Router.GET("/:clueID", s.ClueUIHandler)
	// s.Router.POST("/:clueID", s.ClueUIPOSTHandler)
	// s.Router.GET("/quiz", s.QuizHandler)
	// s.Router.POST("/quiz", s.QuizHandler)
}

func DebugUIHandler(c *gin.Context) {
	clueID, err := strconv.ParseInt(c.Param("clueID"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid clueID"})
		return
	}

	clue, _ := mods.GetClue(clueID)
	game, _ := mods.GetGame(clue.GameID)
	category, _ := mods.GetCategory(clue.CategoryID)

	dat := struct {
		*mods.Clue
		*mods.Game
		*mods.Category
	}{
		Clue:     clue,
		Game:     game,
		Category: category,
	}
	c.HTML(200, "debug.html.tpl", dat)
}

// func QuizHandler(c *gin.Context) {
// 	if c.Request.Method == "POST" {
// 		cor := c.PostForm("correct")
// 		inc := c.PostForm("incorrect")

// 		correct := len(cor) > len(inc)
// 		if correct {
// 			s.QZ.Correct++
// 		} else {
// 			s.QZ.Incorrect++
// 		}
// 		s.QZ.Total++
// 	}

// 	clue, err := mods.GetRandomClue(nil)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch random clue"})
// 		return
// 	}

// 	cat, err := mods.GetCategory(clue.CategoryID)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch category for clue"})
// 		return
// 	}

// 	game, err := mods.GetGame(clue.GameID)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch game for clue"})
// 		return
// 	}

// 	s.QZ.Clues = append(s.QZ.Clues, clue)

// 	c.HTML(200, "quiz.html.tpl", gin.H{
// 		"Clue":     clue,
// 		"Category": cat,
// 		"Session":  s.QZ,
// 		"Game":     game,
// 		"Viz":      Viz(*s.QZ),
// 	})
// }

// // ClueUIPOSTHandler godoc
// func ClueUIPOSTHandler(c *gin.Context) {
// 	clueIDStr := c.Param("clueID")
// 	clueID, err := strconv.ParseInt(clueIDStr, 10, 64)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "invalid clueID"})
// 		return
// 	}

// 	clue, err := mods.GetClue(clueID)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch clue for clueID"})
// 		return
// 	}

// 	categoryIDStr := c.PostForm("cat-sel")
// 	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "invalid categoryID"})
// 		return
// 	}

// 	clues, err := mods.ListClues(models.CluesParams{GameID: clue.GameID, CategoryID: categoryID})
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch clues for categoryID"})
// 		return
// 	}
// 	log.Infof("clues: %s", pretty.Sprint(clues))

// 	newClueIDStr := fmt.Sprintf("%d", clues[0].ClueID)
// 	c.Params = gin.Params{gin.Param{Key: "clueIDStr", Value: newClueIDStr}}
// 	c.Request.Method = "GET"
// 	c.Redirect(302, "/"+newClueIDStr)
// }

// // ClueUIHandler godoc
// func ClueUIHandler(c *gin.Context) {
// 	if c.Request.Method == "POST" {
// 		s.ClueUIPOSTHandler(c)
// 		return
// 	}

// 	clueID, err := strconv.ParseInt(c.Param("clueID"), 10, 64)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "invalid clueID"})
// 		return
// 	}

// 	clue, err := mods.GetClue(clueID)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch clue for clueID"})
// 		return
// 	}

// 	game, err := mods.GetGame(clue.GameID)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch game for clue"})
// 		return
// 	}

// 	category, err := mods.GetCategory(clue.CategoryID)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch category for clue"})
// 		return
// 	}

// 	clueJSON := s.jsonHelper(clue)

// 	categoriesForGame, err := mods.GetCategoriesForGame(clue.GameID)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch categories for game"})
// 		return
// 	}

// 	catOpts := []models.Option{}
// 	for _, c := range categoriesForGame {
// 		catOpts = append(catOpts, models.Option{
// 			OptionKey: fmt.Sprintf("%d", c.CategoryID),
// 			OptionVal: c.Name,
// 			Selected:  c.CategoryID == clue.CategoryID,
// 		})
// 	}

// 	cluesForGame, err := mods.ListClues(models.CluesParams{GameID: clue.GameID})
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "couldn't fetch clues for game"})
// 		return
// 	}

// 	nextClue := models.Clue{}
// 	prevClue := models.Clue{}
// 	for i, c := range cluesForGame {
// 		if c.ClueID == clue.ClueID {
// 			if i > 0 {
// 				prevClue = *cluesForGame[i+1]
// 			}
// 			if i < len(cluesForGame)-1 {
// 				nextClue = *cluesForGame[i-1]
// 			}
// 		}
// 	}

// 	navLinks := models.NavLinks{
// 		NextClue: fmt.Sprintf("%d", nextClue.ClueID),
// 		PrevClue: fmt.Sprintf("%d", prevClue.ClueID),
// 	}

// 	options := &models.Options{
// 		ClueID:          clue.ClueID,
// 		Links:           navLinks,
// 		CategoryOptions: catOpts,
// 	}

// 	debug := struct {
// 		Clue     *models.Clue
// 		Game     *models.Game
// 		Category *models.Category
// 		Stats    *models.Stats
// 		Options  *models.Options
// 		ClueJSON string
// 	}{
// 		Clue:     clue,
// 		Game:     game,
// 		Category: category,
// 		Stats:    s.Stats,
// 		Options:  options,
// 		ClueJSON: pretty.Sprint(clueJSON),
// 	}
// 	c.HTML(200, "clue-base.html.tpl", debug)
// }

// func jsonHelper(clue *models.Clue) map[string]interface{} {
// 	cat, _ := mods.GetCategory(clue.CategoryID)
// 	game, _ := mods.GetGame(clue.GameID)

// 	return map[string]interface{}{
// 		"clueID":       clue.ClueID,
// 		"categoryID":   clue.CategoryID,
// 		"categoryName": cat.Name,
// 		"gameID":       clue.GameID,
// 		"gameDate":     game.GameDate,
// 		"question":     clue.Question,
// 		"answer":       clue.Answer,
// 	}
// }

// // adapted from
// // https://github.com/go-echarts/go-echarts/blob/master/templates/base.go
// // https://github.com/go-echarts/go-echarts/blob/master/templates/header.go
// var baseTpl = `
// <div class="container">
//     <div class="item" id="{{ .ChartID }}" style="width:250;height:250;"></div>
// </div>
// {{- range .JSAssets.Values }}
//    <script src="{{ . }}"></script>
// {{- end }}
// <script type="text/javascript">
//     "use strict";
//     let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
//     let option_{{ .ChartID | safeJS }} = {{ .JSON }};
//     goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
//     {{- range .JSFunctions.Fns }}
//     {{ . | safeJS }}
//     {{- end }}
// </script>
// `

// type snippetRenderer struct {
// 	c      interface{}
// 	before []func()
// }

// func newSnippetRenderer(c interface{}, before ...func()) chartrender.Renderer {
// 	return &snippetRenderer{c: c, before: before}
// }

// func (r *snippetRenderer) Render(w io.Writer) error {
// 	const tplName = "chart"
// 	for _, fn := range r.before {
// 		fn()
// 	}

// 	tpl := template.
// 		Must(template.New(tplName).
// 			Funcs(template.FuncMap{
// 				"safeJS": func(s interface{}) template.JS {
// 					return template.JS(fmt.Sprint(s))
// 				},
// 			}).
// 			Parse(baseTpl),
// 		)

// 	err := tpl.ExecuteTemplate(w, tplName, r.c)
// 	return err
// }

// func renderToHtml(c interface{}) template.HTML {
// 	var buf bytes.Buffer
// 	r := c.(chartrender.Renderer)
// 	err := r.Render(&buf)
// 	if err != nil {
// 		log.Printf("Failed to render chart: %s", err)
// 		return ""
// 	}

// 	return template.HTML(buf.String())
// }

// func Viz(qz QuizSession) template.HTML {
// 	// initialize chart
// 	pie := charts.NewPie()
// 	pie.Renderer = newSnippetRenderer(pie, pie.Validate)

// 	// preformat data
// 	pieData := []opts.PieData{
// 		{Name: "Correct", Value: qz.Correct, ItemStyle: &opts.ItemStyle{Color: "#55bf32"}},
// 		{Name: "Incorrect", Value: qz.Incorrect, ItemStyle: &opts.ItemStyle{Color: "#eb4034"}},
// 	}

// 	// put data into chart
// 	pie.AddSeries("Answers", pieData).SetSeriesOptions(
// 		charts.WithLabelOpts(opts.Label{Show: false, Formatter: "{b}: {c}"}),
// 	)

// 	var htmlSnippet template.HTML = renderToHtml(pie)
// 	return htmlSnippet
// }
