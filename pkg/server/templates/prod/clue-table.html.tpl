{{define "clue-table"}}
<table class="rand-tbl">
  <tr>
    <th>ClueID</th>
    <td><a href="/api/clues/{{.Clue.ClueID}}" target="_blank">{{.Clue.ClueID}}</a></td>
  </tr>
  <tr>
    <th>GameID</th>
    <td><a href="/api/games/{{.Clue.GameID}}" target="_blank">{{.Clue.GameID}}</a></td>
  </tr>
  <tr>
    <th>CategoryID</th>
    <td><a href="/api/categories/{{.Clue.CategoryID}}" target="_blank">{{.Clue.CategoryID}}</a></td>
  </tr>
  <tr>
    <th>Question</th>
    <td>{{.Clue.Question}}</td>
  </tr>
  <tr>
    <th>Answer</th>
    <td>{{.Clue.Answer}}</td>
  </tr>
</table>
{{end}}
{{define "clue-json"}}
<h3>JSON</h3>
<div class="pretext">
  <pre>{{.}}</pre>  
</div>
{{end}}