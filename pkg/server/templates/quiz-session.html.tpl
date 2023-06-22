{{define "quiz-session"}}
<div class="quizsession">
  <h2>Session Stats</h2>
  <hr>
  <p><strong>Total</strong>: <code>{{.Session.Total}}</code></p>
  <p><strong>Correct</strong>: <code>{{.Session.Correct}}</code></p>
  <p><strong>Incorrect</strong>: <code>{{.Session.Incorrect}}</code></p>
  <div>{{.Viz}}</div>
</div>
{{end}}