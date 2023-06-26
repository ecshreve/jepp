<!DOCTYPE html>
<head>
  <title>Jepp</title>
  <!-- Google tag (gtag.js) -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=G-J6ZJ2Y9HHQ"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());

    gtag('config', 'G-J6ZJ2Y9HHQ');
  </script>
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
<div style="display: flex; justify-content: space-between; align-items: center; padding-left: 10px; padding-right: 20px;">
  <a href="/" style="text-decoration: none; color: black;"><h1>Jepp</h1></a>
  <p>access to <strong>{{.NumClues}}</strong> jeopardy clues via rest api</p>
  <a class="ghcustom" href="https://github.com/ecshreve/jepp"><i class="fa fa-brands fa-github" style="font-size: xx-large; text-decoration: none; color: black;"></i></a>
</div>
<hr>
<br>
<div class="card-container">
  <div style="display: flex; flex-direction: column;">
    <div style="margin-bottom: 20px;">
      <div class="card">
        <div class="container">
        <h3>API fun with Jeopardy! Access historical Jeopardy clues scraped from <a href="https://www.j-archive.com/" target="_blank">J-ARCHIVE</a> via a simple api.</h2>
        <hr>
        <a href="https://github.com/ecshreve/jepp">
          <div>
            <p>
              view project on github <i class="fa fa-brands fa-github"></i>
            </p>
             <p>
              <img alt="GitHub Workflow Status (with event)" src="https://img.shields.io/github/actions/workflow/status/ecshreve/jepp/.github%2Fworkflows%2Fci.yml">
              <img alt="GitHub release (release name instead of tag name)" src="https://img.shields.io/github/v/release/ecshreve/jepp">
            </p>
          </div>
        </a>
        </div>
      </div>
    </div>
     <div style="margin-bottom: 25px;">
      <div class="card">
        <div class="container">
        <h4>Just interested in the data? Feel free to download the CSV files or SQL dump from the GitHub Repository</h4>
        <p>
        The most recent data files should be attached to the <a href="https://github.com/ecshreve/jepp/releases" target="_blank">most recent release here.</a>
        <br><br>
        More information on the data itself can be found in the  <a href="https://github.com/ecshreve/jepp/tree/main/data" target="_blank">`data` directory's `readme`.</a>
        </p>
        </div>
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
  <div class="card-boring">
    <div class="container">
      <div id="swagger-ui"></div> <!-- Div to hold the UI component -->
    </div>
  </div>
</div>
</html>