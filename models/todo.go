package models

import (
        "../config"
        "fmt"
        "time"
)

type Todo struct {
        Id        int64
        Name      string
        Completed bool
        Due       time.Time
}

type Todos []Todo

func TodosIndex() Todos {
        db := config.Db()
        rows, err := db.Query("SELECT * from todos")
        if err != nil {
                panic(err)
        }
        defer rows.Close()

        todos := make(Todos, 5)
        for rows.Next() {
                fmt.Println(rows.Scan())
        }
        return todos
}
