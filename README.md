
```sample.go
package main

import (
        "context"
        "database/sql"
        "fmt"

        _ "github.com/lib/pq"
        "github.com/m0cchi/textstack/model"
)

func main() {
        ctx := context.Background()
        db, err := sql.Open("postgres", "host=localhost user=textstack password=textstack sslmode=disable")
        if err != nil {
                fmt.Println("Opps")
                return
        }
        defer db.Close()

        q := model.New(db)

        textparams := model.CreateTextParams{
                Title: "test title",
                Body:  "Hey",
        }

        uuid, err := q.CreateText(ctx, textparams)
        if err != nil {
                fmt.Printf("%v", err)
                return
        }

        text, err := q.GetText(ctx, uuid)
        if err != nil {
                fmt.Printf("%v", err)
                return
        }
        fmt.Printf("%v", text)

}
```
