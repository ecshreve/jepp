<head>
  <title>Jepp</title>
  <link rel="stylesheet" href="/style.css">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<h1>Jepp</h1>
<div>
  <p>access <strong>{{.Stats.TotalClues}}</strong> jeopardy clues via rest api</p>
  <a href="/swagger/index.html">view api documentation</a>
</div>
<div class="card-container" style="justify-content: left;">
   <div class="card" style="max-width: 500px;">
      <div class="rand-cont">
        <h2>Random Clue</h2>
        <a class="rand-clue" onclick="window.location.reload()"><i class="fa fa-refresh"></i></a>
      </div>
      <hr>
    <div class="container">
      {{template "clue-table" .}}
    </div>
  </div>
</div>
     