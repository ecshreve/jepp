<!DOCTYPE html>

<head>
  <title>Jepp</title>
  <!-- Google tag (gtag.js) -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=G-J6ZJ2Y9HHQ"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag() {dataLayer.push(arguments);}
    gtag('js', new Date());

    gtag('config', 'G-J6ZJ2Y9HHQ');
  </script>
  <link rel="stylesheet" href="/style.css">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
  <script src="https://unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
  <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3/swagger-ui.css" />
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
  <script src="https://cdn.jsdelivr.net/npm/react@17.0.2/umd/react.production.min.js" crossorigin="anonymous"></script>
  <script src="https://cdn.jsdelivr.net/npm/react-dom@17.0.2/umd/react-dom.production.min.js"
    crossorigin="anonymous"></script>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphiql@2.0.7/graphiql.min.css" crossorigin="anonymous" />
  <script>
      const openView = function (evt, viewName) {
        var i, tabcontent, tablinks;
        tabcontent = document.getElementsByClassName("tabcontent");
        for (i = 0; i < tabcontent.length; i++) {
          tabcontent[i].style.display = "none";
        }
        tablinks = document.getElementsByClassName("tablinks");
        for (i = 0; i < tablinks.length; i++) {
          tablinks[i].className = tablinks[i].className.replace(" active", "");
        }
        document.getElementById(viewName).style.display = "block";
        evt.currentTarget.className += " active";

        if (viewName === "graphiql") {
          document.getElementById("newTab").className += " active";
        }
      }
    </script>
</head>

<body>
  <div
    style="display: flex; justify-content: space-between; align-items: center; padding-left: 10px; padding-right: 20px;">
    <a href="/" style="text-decoration: none; color: black;">
      <h1>Jepp</h1>
    </a>
    <p>access to <strong>{{.NumClues}}</strong> jeopardy clues via rest api</p>
    <a class="ghcustom" href="https://github.com/ecshreve/jepp"><i class="fa fa-brands fa-github"
        style="font-size: xx-large; text-decoration: none; color: black;"></i></a>
  </div>
  <hr>
  <br>
  <div class="card-container">
    <div style="display: flex; flex-direction: column;">
      <div style="margin-bottom: 20px;">
        <div class="card">
          <div class="container">
            <h3>API fun with Jeopardy! Access historical Jeopardy clues scraped from <a
                href="https://www.j-archive.com/" target="_blank">J-ARCHIVE</a> via a http or graphql.</h2>
              <p>
                Both the http and graphql interfaces are generated via <a href="https://entgo.io/docs/getting-started"
                  target="_blank">ent</a>. The graphql
                endpoint uses the entgql plugin, it generates a graphql schema from the ent schema. The http endpoint is
                generated via
                the entoas and ogent plugins, which generate an openapi spec from the ent schema. The openapi spec is
                then used to generate
                the swagger ui, as well as the go server implementation.
              </p>
              <hr>
              <a href="https://github.com/ecshreve/jepp">
                <div>
                  <p>
                    view project on github <i class="fa fa-brands fa-github"></i>
                  </p>
                  <p>
                    <img alt="GitHub Workflow Status (with event)"
                      src="https://img.shields.io/github/actions/workflow/status/ecshreve/jepp/.github%2Fworkflows%2Fci.yml">
                    <img alt="GitHub release (release name instead of tag name)"
                      src="https://img.shields.io/github/v/release/ecshreve/jepp">
                  </p>
                </div>
              </a>
          </div>
        </div>
      </div>
      <div style="margin-bottom: 25px;">
        <div class="card">
          <div class="container">
            <h4>Just interested in the data? Feel free to download the SQLite db from the GitHub Repository</h4>
            <p>
              The most recent sqlite db file should be attached to the <a
                href="https://github.com/ecshreve/jepp/releases" target="_blank">most recent release here.</a>
              <br><br>
            </p>
          </div>
        </div>
      </div>
      <div>
        <div class="card">
          <div class="rand-cont">
            <h2 style="text-align: center;">Example Random Clue</h2>
            <a class="rand-clue" onclick="window.location.reload()"><i class="fa fa-refresh"></i></a>
          </div>
          <hr>
          <div class="container">
            <table class="rand-tbl">
              <tr>
                <th>ClueID</th>
                <td><a href="/clues/{{.Clue.ID}}" target="_blank">{{.Clue.ID}}</a></td>
              </tr>
              <tr>
                <th>GameID</th>
                <td><a href="/games/{{.Clue.GameID}}" target="_blank">{{.Clue.GameID}}</a></td>
              </tr>
              <tr>
                <th>CategoryID</th>
                <td><a href="/categories/{{.Clue.CategoryID}}" target="_blank">{{.Clue.CategoryID}}</a></td>
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
        </div>
      </div>
    </div>
    <div class="card-boring">
      <div class="tab">
        <button id="defaultOpen" class="tablinks" onclick="openView(event, 'swagger-ui')">Swagger</button>
        <button class="tablinks" onclick="openView(event, 'graphiql')">GraphQL</button>
        <button id="newTab" class="tablinks" onclick="window.open('/graphql')">
          <svg xmlns="http://www.w3.org/2000/svg" height="1em"
            viewBox="0 0 512 512"><!--! Font Awesome Free 6.4.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
            <path
              d="M352 0c-12.9 0-24.6 7.8-29.6 19.8s-2.2 25.7 6.9 34.9L370.7 96 201.4 265.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0L416 141.3l41.4 41.4c9.2 9.2 22.9 11.9 34.9 6.9s19.8-16.6 19.8-29.6V32c0-17.7-14.3-32-32-32H352zM80 32C35.8 32 0 67.8 0 112V432c0 44.2 35.8 80 80 80H400c44.2 0 80-35.8 80-80V320c0-17.7-14.3-32-32-32s-32 14.3-32 32V432c0 8.8-7.2 16-16 16H80c-8.8 0-16-7.2-16-16V112c0-8.8 7.2-16 16-16H192c17.7 0 32-14.3 32-32s-14.3-32-32-32H80z" />
          </svg>
        </button>
      </div>
      <div class="container">
        <div id="graphiql" class="tabcontent">
         looading
        </div>
        <div id="swagger-ui" class="tabcontent"></div>
      </div>
    </div>
  </div>
  <div>
    
    <script src="https://cdn.jsdelivr.net/npm/graphiql@2.0.7/graphiql.min.js" crossorigin="anonymous"></script>

    <script>
      const url = "/query"
      const subscriptionUrl = "";

      const fetcher = GraphiQL.createFetcher({url, subscriptionUrl});
      ReactDOM.render(
        React.createElement(GraphiQL, {
          fetcher: fetcher,
          isHeadersEditorEnabled: true,
          shouldPersistHeaders: true,
        }),
        document.getElementById('graphiql'),
      );
    </script>
  </div>
</body>
<script>
  document.getElementById("defaultOpen").click();
</script>
</html>