<head>
  <title>Jepp</title>
  <link rel="stylesheet" href="/style.css">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
  <script src="https://unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
  <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3/swagger-ui.css"/>
</head>
<div style="display: flex; justify-content: space-between; align-items: center; padding-left: 10px; padding-right: 10px;">
  <a href="/" style="text-decoration: none; color: black;"><h1>Jepp</h1></a>
  <p>access to <strong>{{.Stats.TotalClues}}</strong> jeopardy clues via rest api</p>
  <a href="https://github.com/ecshreve/jepp">view project on github <i class="fa fa-brands fa-github"></i></a>
</div>
<hr>
<div class="card-container">
  <div class="card-boring">
    <div class="container">
      {{template "swagger"}}
    </div>
  </div>
  <div>
   <div class="card">
      <h2 style="text-align: center;">Example: Random Clue</h2>
      <div class="rand-cont">
        <pre>GET /api/clues/random</pre>
        <a class="rand-clue" onclick="window.location.reload()"><i class="fa fa-refresh"></i></a>
      </div>
      <hr>
      <div class="container">
        {{template "clue-table" .}}
      </div>
      <div class="container">
        {{template "clue-json" .ClueJSON}}
      </div>
    </div>
  </div>
</div>
     