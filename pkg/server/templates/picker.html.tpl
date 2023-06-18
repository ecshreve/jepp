{{define "picker"}}
<form
  method="POST"
  action="/"
>
  <label for="cars">Clues:</label>
  <select name="clue-sel" id="clue-sel">
    {{range .}}
      <option value={{.ClueID}} {{if .Selected}}selected{{end}}>{{.ClueID}}</option>
    {{end}}
  </select>
  <br><br>
  <input type="submit" value="Submit">
</form>
{{end}}