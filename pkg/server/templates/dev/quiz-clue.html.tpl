{{define "quiz-clue"}}
<div class="quizclue" onclick="document.getElementById('answer').style.opacity=100">
  <h2 style="margin-bottom: 0px;">{{.Category.Name}}</h2>
  <small style="color: lightgray;">{{.Game.GameID}} - {{.Game.GameDate}}</small>
  <hr>
  <p><strong>Question</strong>: <code>{{.Clue.Question}}</code></p>
</div>
{{end}}