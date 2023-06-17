<head>
  <title>Random Clue</title>
  <link rel="stylesheet" href="/style.css">
</head>
<h1>RANDOM</h1>
<div class="card-container">
  <div class="card" onClick="window.location.reload();">
    <div class="container">
      <div>
        <h2>CLUE</h2>
        <p><strong>ClueID</strong>: <code>{{.Clue.ClueID}}</code></p>
        <p><strong>GameID</strong>: <code>{{.Clue.GameID}}</code></p>
        <p><strong>CategoryID</strong>: <code>{{.Clue.CategoryID}}</code></p>
        <p><strong>Question</strong>: <code>{{.Clue.Question}}</code></p>
        <p><strong>Answer</strong>: <code>{{.Clue.Answer}}</code></p>
      </div>
      <div>
        <h2>Game</h2>
        <p><strong>GameID</strong>: <code>{{.Game.GameID}}</code></p>
        <p><strong>ShowNum</strong>: <code>{{.Game.ShowNum}}</code></p>
        <p><strong>GameDate</strong>: <code>{{.Game.GameDate}}</code></p>
      </div>
      <div>
        <h2>Category</h2>
        <p><strong>CategoryID</strong>: <code>{{.Category.CategoryID}}</code></p>
        <p><strong>CategoryName</strong>: <code>{{.Category.Name}}</code></p>
      </div>
    </div>
  </div> 
</div>