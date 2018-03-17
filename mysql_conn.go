package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ylx?charset=utf8")
	if err != nil {
		panic(err.Error())
		return
	}
	defer db.Close()

	err = db.Ping();
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("select id,username from ylx_member limit 5")    
    defer rows.Close()    
    for rows.Next() {
        var id int    
        var name string    
        err = rows.Scan(&id, &name)    
        fmt.Printf("rows id = %d, value = %s \n", id, name)    
	}
	err = rows.Err()  
    if err != nil {
        panic(err.Error())  
    }
}