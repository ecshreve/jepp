{{define "cat-details"}}
<h3>Category Details</h3>
<p><em>CategoryID</em>: {{ .CategoryID }}</p>
<p><em>CategoryName</em>: {{ .Name }}</p>
{{end}}
{{define "clue-json"}}
<h3>JSON</h3>
<div class="pretext">
  <pre>{{.}}</pre>  
</div>
{{end}}