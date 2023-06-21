{{define "quiz-clue"}}
<div class="quizclue" onclick="document.getElementById('answer').style.opacity=100">
  <h2>{{.Category.Name}}</h2>
  <hr>
  <p><strong>Question</strong>: <code>{{.Clue.Question}}</code></p>
</div>
<div class="answer" id="answer" style="opacity: 0">
  <p><strong>Answer</strong>: <code>{{.Clue.Answer}}</p></code></p>
  <form
    method="POST"
    action="/quiz"
  >
    <input name="correct" type="submit" value="correct">
    <input name="incorrect" type="submit" value="incorrect">
  </form>
</div>
{{end}}