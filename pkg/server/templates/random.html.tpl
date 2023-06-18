<head>
  <title>Random Clue</title>
  <link rel="stylesheet" href="/style.css">
</head>
<h1>RANDOM</h1>
<div class="card-container">
  <div class="card-boring">
    <div class="container">
      <h2>Stats</h2>
      <hr>
      <h3>Totals</h3>
      <div>
        <p><em>Games</em>: <code>{{.Stats.TotalGames}}</code></p>
      </div>
      <div>
       <p><em>Categories</em>: <code>{{.Stats.TotalCategories}}</code></p>
      </div>
      <div>
        <p><em>Clues</em>: <code>{{.Stats.TotalClues}}</code></p>
      </div>
      <hr>
        <h3>Category Details</h3>
        <p><em>CategoryID</em>: {{ .Clue.CategoryID }}</p>
        <p><em>CategoryName</em>: {{ .Category.Name }}</p>
        <p><em>CategoryGamesCount</em>: {{ .CategoryGamesCount }}</p>
        <p><em>CategoryCluesCount</em>: {{ .CategoryCluesCount }}</p>
      <hr>
        <h3>JSON</h3>
      <div class="pretext">
        <pre>{{.ClueJSON}}</pre>  
      </div>
    </div>
  </div>
  <!-- guess form -->
        <div class="col-4">
          
        </div>
  <div class="card" onClick="window.location.reload();">
    <div class="container">
      <div class="details">
        <div style="display: flex; flex-direction: row;justify-content: space-between; align-items: center;">
          <div style="display: flex;">
            <h2>Game</h2>
          </div>
          <div>
            <form
              method="POST"
              action="/"
            >
              <div style="display: flex;">
                <button id="game-roll" name="game-roll" type="submit" class="btn-custom" value="GAME">roll 🎲</button>
              </div>
            </form>
          </div>
        </div>
        <p><strong>GameID</strong>: <code>{{.Game.GameID}}</code></p>
        <p><strong>ShowNum</strong>: <code>{{.Game.ShowNum}}</code></p>
        <p><strong>GameDate</strong>: <code>{{.Game.GameDate}}</code></p>
      </div>
      <div class="details">
        <div style="display: flex; flex-direction: row;justify-content: space-between;align-items: center;">
          <div style="display: flex;">
            <h2>Category</h2>
          </div>
          <div>
            <form
              method="POST"
              action="/"
            >
              <select name="cat-select" id="cat-select">
                {{ $clue_cat := .Clue.CategoryID }}
                {{ range .GameCategories }}
                <option value={{ .CategoryID }} {{if eq .CategoryID $clue_cat}}selected{{end}}>{{ .Name }}</option>
                {{ end }}
              </select> 
            </form>
          </div>
          <div style="display: flex;">
          <form
              method="POST"
              action="/"
            >
              <button id="cat-roll" name="cat-roll" type="submit" class="btn-custom" value={{ .Clue.ClueID }}>roll 🎲</button>
            </form>
          </div>
        </div>
        <p><strong>CategoryID</strong>: <code>{{.Category.CategoryID}}</code></p>
        <p><strong>CategoryName</strong>: <code>{{.Category.Name}}</code></p>
      </div>
      <div class="details">
        <div style="display: flex; flex-direction: row;justify-content: space-between;align-items: center;">
          <div>
            <h2>CLUE</h2>
          </div>
           <div>
            <form
              method="POST"
              action="/"
            >
              <select name="clue-select" id="clue-select">
                {{ $clue_id := .Clue.ClueID }}
                {{ range .CategoryClues }}
                <option value={{ .ClueID }} {{if eq .ClueID $clue_id}}selected{{end}}>{{ .ClueID }}</option>
                {{ end }}
              </select> 
            </form>
          </div>
          <div>
            <form
              method="POST"
              action="/"
            >
              <button id="clue-roll" name="clue-roll" type="submit" class="btn-custom" value={{ .Clue.ClueID }}>roll 🎲</button>
            </form>
          </div>
        </div>
        <p><strong>ClueID</strong>: <code>{{.Clue.ClueID}}</code></p>
        <p><strong>GameID</strong>: <code>{{.Clue.GameID}}</code></p>
        <p><strong>CategoryID</strong>: <code>{{.Clue.CategoryID}}</code></p>
        <p><strong>Question</strong>: <code>{{.Clue.Question}}</code></p>
        <p><strong>Answer</strong>: <code>{{.Clue.Answer}}</code></p>
      </div>
    </div>
  </div> 
</div>