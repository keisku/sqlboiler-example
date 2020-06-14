package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/kskumgk63/sqlboiler-example/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	ctx      = context.Background()
	host     = os.Getenv("PSQL_HOST")
	user     = os.Getenv("PSQL_USER")
	password = os.Getenv("PSQL_PASSWORD")
	dbname   = os.Getenv("PSQL_DATABASE")
	sslmode  = os.Getenv("PSQL_SSLMODE")
	dns      = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbname, sslmode)
)

func main() {
	conn, err := sql.Open("postgres", dns)
	if err != nil {
		panic(err)
	}

	// check if success to connect
	if err = conn.Ping(); err != nil {
		panic(err)
	}

	// create new `article`
	http.HandleFunc("/article/new", func(w http.ResponseWriter, r *http.Request) {
		count, err := models.Articles().Count(ctx, conn)
		if err != nil {
			panic(err)
		}
		count++

		newArticle := models.Article{
			ID:      count,
			Title:   fmt.Sprintf("title %d", count),
			Content: fmt.Sprintf("content %d", count),
		}
		if err = newArticle.Insert(ctx, conn, boil.Infer()); err != nil {
			panic(err)
		}

		// render html
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, `<p style="color:red;">success to create new article</p>`)
		fmt.Fprintln(w, `<a href="/articles">go back to list</a>`)
	})

	// show list of `articles`
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		articles, err := models.Articles().All(ctx, conn)
		if err != nil {
			panic(err)
		}

		// render html
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<p>total = %d</p>", len(articles))
		for _, article := range articles {
			fmt.Fprintf(w, "<p>id: %v\ntitle: %v\ncontent: %v\n created_at: %v</p>", article.ID, article.Title, article.Content, article.CreatedAt)
		}
		fmt.Fprintln(w, `<a href="article/new">new</a>`)
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
