package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
)

func rootServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
<!DOCTYPE html>
<html>

<head>
  <title>Waypoint Go Example</title>
  <link rel="stylesheet" type="text/css" href="/static/stylesheets/main.css" />
</head>

<body>
  <div class="container">
    <header>
      <a href="https://waypointproject.io" class="logo">
        <img src="/static/images/logo.svg" alt="Logo" />
      </a>
    </header>
    <section class="content">
      <div class="language-icon">
        <img class="language-image" src="/static/images/language.svg" alt="Icon" />
      </div>
      <h1>` + os.Getenv("SAY_HI") + `</h1>
      <p>
        And that was done using a targeted runner!
      </p>
      <p>
        Read the <a href="https://waypointproject.io/docs">documentation</a> for more about Waypoint.
      </p>
    </section>
    <footer>
      <a href="https://hashicorp.com" class="hashi">
        <img src="/static/images/hashi.svg" alt="HashiCorp" />
      </a>
    </footer>
  </div>
</body>

</html>
	`))
}

func main() {
	assetServer := http.FileServer(http.Dir("./"))

	http.HandleFunc("/", http.HandlerFunc(rootServer))
	http.Handle("/static/", assetServer)

	fmt.Printf("Starting server at: http://0.0.0.0:3030\n")
	if err := http.ListenAndServe("0.0.0.0:3030", nil); err != nil {
		log.Fatal(err)
	}
}
