{{define "category-header"}}
<div style="display: flex; flex-direction: row;justify-content: space-between;align-items: center;">
  <div style="display: flex;">
    <h2>Category</h2>
  </div>
  <div style="display: flex;">
    <form
        method="POST"
        action="/"
      >
        <button id="cat-roll" name="cat-roll" type="submit" class="btn-custom" value={{ . }}>roll 🎲</button>
    </form>
  </div>
</div>
{{end}}
{{define "category-content"}}
<p><strong>CategoryID</strong>: <code>{{.CategoryID}}</code></p>
<p><strong>CategoryName</strong>: <code>{{.Name}}</code></p>
{{end}}
