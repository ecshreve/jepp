<head>
  <title>Random Clue</title>
  <link rel="stylesheet" href="/style.css">
</head>
<h1>RANDOM</h1>
<div class="card-container">
  <div class="card-boring">
    <div class="container">
      <div>
        <h3>Games</h3>
        <p><strong>Total</strong>: <code>{{.Stats.TotalGames}}</code></p>
      </div>
      <div>
        <h3>Categories</h3>
        <p><strong>Total</strong>: <code>{{.Stats.TotalCategories}}</code></p>
      </div>
      <div>
        <h3>CLUES</h3>
        <p><strong>Total</strong>: <code>{{.Stats.TotalClues}}</code></p>
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
          <h2>CLUE</h2>
          <form
            method="POST"
            action="/"
          >
            <button id="clue-roll" name="clue-roll" type="submit" class="btn-custom" value={{ .Clue.ClueID }}>roll 🎲</button>
          </form>
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