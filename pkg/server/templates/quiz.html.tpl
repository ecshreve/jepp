<head>
  <title>Quiz</title>
  <link rel="stylesheet" href="/style.css">
</head>
<h1>QUIZ MODE!</em></small></h1>
<div style="display: flex;">
  <div>
    <div class="card-container">
      <div class="card">
        {{template "quiz-clue" .}}
      </div>
    </div>
    <div>
      <div>
        {{template "quiz-answer" .}}
      </div>
    </div>
  </div>
  <div class="card-container">
    <div class="card-boring">
      <div class="container">
        {{template "quiz-session" .}}
      </div>
    </div>
  </div>
</div>