{{define "clue-table"}}
<table class="rand-tbl">
  <tr>
    <th>ClueID</th>
    <td><a href="/debug/{{.Clue.ClueID}}">{{.Clue.ClueID}}</a></td>
  </tr>
  <tr>
    <th>GameID</th>
    <td>{{.Clue.GameID}}</td>
  </tr>
  <tr>
    <th>CategoryID</th>
    <td>{{.Category.CategoryID}}</td>
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