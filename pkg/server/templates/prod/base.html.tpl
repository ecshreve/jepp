<head>
  <title>Jepp</title>
  <link rel="stylesheet" href="/style.css">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
  <script src="https://unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
  <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3/swagger-ui.css"/>
  <script>
    window.onload = function () {
        // Begin Swagger UI call region
        const ui = SwaggerUIBundle({
            url: "/swagger/doc.json", //Location of Open API spec in the repo
            dom_id: '#swagger-ui',
            deepLinking: false,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIBundle.SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
        })
        window.ui = ui
    }
</script>
</head>
<div style="display: flex; justify-content: space-between; align-items: center; padding-left: 10px; padding-right: 10px;">
  <a href="/" style="text-decoration: none; color: black;"><h1>Jepp</h1></a>
  <p>access to <strong>{{.NumClues}}</strong> jeopardy clues via rest api</p>
  <a href="https://github.com/ecshreve/jepp">view project on github <i class="fa fa-brands fa-github"></i></a>
</div>
<hr>
<div class="card-container">
  <div class="card-boring">
    <div class="container">
      <div id="swagger-ui"></div> <!-- Div to hold the UI component -->
    </div>
  </div>
  <div>
   <div class="card">
      <h2 style="text-align: center;">Example: Random Clue</h2>
      <div class="rand-cont">
        <pre>GET /api/clue?random</pre>
        <a class="rand-clue" onclick="window.location.reload()"><i class="fa fa-refresh"></i></a>
      </div>
      <hr>
      <div class="container">
        <table class="rand-tbl">
          <tr>
            <th>ClueID</th>
            <td><a href="/api/clue?id={{.Clue.ClueID}}" target="_blank">{{.Clue.ClueID}}</a></td>
          </tr>
          <tr>
            <th>GameID</th>
            <td><a href="/api/game?id={{.Clue.GameID}}" target="_blank">{{.Clue.GameID}}</a></td>
          </tr>
          <tr>
            <th>CategoryID</th>
            <td><a href="/api/category?id={{.Clue.CategoryID}}" target="_blank">{{.Clue.CategoryID}}</a></td>
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
      </div>
      <div class="container">
        <h3>JSON</h3>
        <div class="pretext">
          <pre>{{.ClueJSON}}</pre>  
        </div>
     </div>
    </div>
  </div>
</div>
     