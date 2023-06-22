{{define "quiz-answer"}}
<div class="answer" id="answer" style="opacity: 0">
  <p><strong>Answer</strong>: <code>{{.Clue.Answer}}</p></code></p>
  <form
    method="POST"
    action="/quiz"
  >
    <div style="display: flex;">
      <input class="score-but" name="correct" type="submit" value="correct">
      <input class="score-but" name="incorrect" type="submit" value="incorrect">
    </div>
  </form>
</div>
{{end}}