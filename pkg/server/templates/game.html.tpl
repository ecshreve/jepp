{{define "game"}}
<div style="display: flex; flex-direction: row;justify-content: space-between; align-items: center;">
  <div style="display: flex;">
    <h2>Game</h2>
  </div>
  <form
    method="POST"
    action="/"
  >
    <div style="display: flex;">
      <button id="game-roll" name="game-roll" type="submit" class="btn-custom" value="GAME">roll 🎲</button>
    </div>
  </form>
</div>
<p><strong>GameID</strong>: <code>{{ .GameID }}</code></p>
<p><strong>ShowNum</strong>: <code>{{ .ShowNum }}</code></p>
<p><strong>GameDate</strong>: <code>{{ .GameDate }}</code></p>
{{end}}
