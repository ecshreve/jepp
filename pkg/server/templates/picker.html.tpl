{{define "picker"}}
<form
  method="POST"
  action="/{{.ClueID}}"
>
  <label for="cats">Categories:</label>
  <select name="cat-sel" id="cat-sel">
    {{range .CategoryOptions}}
      <option value={{.OptionKey}} {{if .Selected}}selected{{end}}>{{.OptionVal}}</option>
    {{end}}
  </select>

  <input type="submit" value="Submit">
</form>

<a id="prev-clue" href="http://10.35.220.99:8880/{{.Links.PrevClue}}"><< prev</a>
<a id="next-clue" href="http://10.35.220.99:8880/{{.Links.NextClue}}">next >></a>

{{end}}