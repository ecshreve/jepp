<head>
  <title>Clue Explorer</title>
  <link rel="stylesheet" href="/style.css">
</head>
<h1>Explorer</h1>
<div class="card-container">
  <div class="card-boring">
    <div class="container">
      <h2>Stats</h2>
      <hr>
      {{template "totals" .Stats}}
      <hr>
      {{template "cat-details" .Category}}
      <hr>
      {{template "clue-json" .ClueJSON}}
    </div>
  </div>
  <div class="card">
    <div class="container">
      <div class="details">
        {{template "game" .Game}}
      </div>
      <div class="details">
        {{template "category" .Category}}
      </div>
      <div class="details">
        {{template "clue" .Clue}}
      </div>
    </div>
  </div> 
</div>