package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	var db databaseCurrent
	db.Databases = append(db.Databases, &database{Item: "shoes", Price: 40})
	db.Databases = append(db.Databases, &database{Item: "socks", Price: 60})
	db.TotalCount = 2
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

type dollars float32

type database struct {
	Item  string
	Price dollars
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type databaseCurrent struct {
	TotalCount int
	Databases  []*database
}

func (db databaseCurrent) list(w http.ResponseWriter, req *http.Request) {
	var dbList = template.Must(template.New("dblist").Parse(`
		<h1>{{.TotalCount}} db</h1>
		<table>
		<tr stype='text-align: left'>
			<th>item</th>
			<th>price</th>
		</tr>
		{{range .Databases}}
		<tr>
			<td>{{.Item}}</td>
			<td>{{.Price}}</td>
		</tr>
		{{end}}
		</table>
		`))
	if err := dbList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}
