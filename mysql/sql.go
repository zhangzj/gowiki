package main

import (
  "fmt"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root:12345678@/test?charset=utf8")
  checkErr(err)

  // insert data
  stmt , err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
  checkErr(err)

  res, err := stmt.Exec("zhangzj", "客舍青青柳色新 Deaprtment", "2017-08-21")
  checkErr(err)

  id, err := res.LastInsertId()
  checkErr(err)

  // update data
  stmt, err = db.Prepare("UPDATE userinfo SET username=? where uid=?")
  checkErr(err)

  res, err = stmt.Exec("zhoujian", id)
  checkErr(err)

  // affect, err := res.RowsAffected()
  // checkErr(err)

  fmt.Println(err)

  fmt.Println(id)

  // select data
  rows, err := db.Query("SELECT * FROM userinfo")
  checkErr(err)

  for rows.Next() {
    var uid int
    var username string
    var department string
    var created string
    err = rows.Scan(&uid, &username, &department, &created)
    checkErr(err)
    fmt.Println(uid, username, department, created)
  }

  // delete data
  // stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
  // checkErr(err)
  // res, err= stmt.Exec(id)
  // checkErr(err)
  // affect, err = res.RowsAffected()
  // checkErr(err)
  // fmt.Println(affect)

  // close database
  db.Close()
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
