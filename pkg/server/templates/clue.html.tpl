{{define "clue"}}
<div style="display: flex; flex-direction: row;justify-content: space-between;align-items: center;">
  <h2>CLUE</h2>
  <form
    method="POST"
    action="/"
  >
    <button id="clue-roll" name="clue-roll" type="submit" class="btn-custom" value={{ .ClueID }}>roll 🎲</button>
  </form>
</div>
<p><strong>ClueID</strong>: <code>{{.ClueID}}</code></p>
<p><strong>GameID</strong>: <code>{{.GameID}}</code></p>
<p><strong>CategoryID</strong>: <code>{{.CategoryID}}</code></p>
<p><strong>Question</strong>: <code>{{.Question}}</code></p>
<p><strong>Answer</strong>: <code>{{.Answer}}</code></p>
{{end}}