## Go RSS aggregator project
[Taked to](https://www.youtube.com/watch?v=un6ZyFkqFKo)
7:07
## RSS aggregator project
# create new module
go mod 

go build

# run
go build && ./rss-aggregator-go

# Chi
Install terminal
go get github.com/joho/godotenv
go mod vendor  //create a new module folder. required install tidy
go mod tidy
go get github.com/go-chi/chi
go get github.com/go-chi/cors


import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)



docgen: auto-generate routing documentation in JSON or Markdown for a `chi` Router from your app source


in terminal
export PORT=8000



# optionals
go get -u github.com/go-chi/chi/v5
