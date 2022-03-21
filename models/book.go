package models

import (
    "database/sql"
    "fmt"

    _"github.com/go-sql-driver/mysql"
)

//增加数据
func insert() {
    db, err := sql.Open("mysql", "wx:wang...123@tcp(20.27.155.16:3306)/bookManage?charset=utf8")
    checkErr(err)

    stmt, err := db.Prepare(`INSERT user (userid,username,userage,usersex) values (?,?,?,?)`)
    checkErr(err)
    res, err := stmt.Exec(1, "Mary", 20, 1)
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)
}

//删除数据
func remove() {
    db, err := sql.Open("mysql", "wx:wang...123@tcp(20.27.155.16:3306)/bookManage?charset=utf8")
    checkErr(err)

    stmt, err := db.Prepare(`DELETE FROM user WHERE userid=?`)
    checkErr(err)
    res, err := stmt.Exec(1)
    checkErr(err)
    num, err := res.RowsAffected()
    checkErr(err)
    fmt.Println(num)
}

//更新数据
func update() {
    db, err := sql.Open("mysql", "wx:wang...123@tcp(20.27.155.16:3306)/bookManage?charset=utf8")
    checkErr(err)

    stmt, err := db.Prepare(`UPDATE user SET userage=?,usersex=? WHERE userid=?`)
    checkErr(err)
    res, err := stmt.Exec(21, 2, 2)
    checkErr(err)
    num, err := res.RowsAffected()
    checkErr(err)
    fmt.Println(num)
}

//查询数据
func queryAll() {
    db, err := sql.Open("mysql", "wx:wang...123@tcp(20.27.155.16:3306)/bookManage?charset=utf8")
    checkErr(err)

    rows, err := db.Query("SELECT * FROM book")
    checkErr(err)

    //    //普通demo
    for rows.Next() {
        
    }
}
var (
	bookList map[string]*Book
)

func init(){
    bookList = make(map[string]*Book)
}
 type Book struct{
    id int
    name string
 }
func Query(id string) (a map[string]*Book,err1 error){
	db, err := sql.Open("mysql", "wx:wang...123@tcp(120.27.155.16:3306)/bookManage?charset=utf8")
    checkErr(err)

   //rows, err := db.Query("SELECT id,name FROM book where id = " + id)
rows, err := db.Query("SELECT id,name FROM book")
    checkErr(err)
    for rows.Next() {
        var id1 int
        var name1 string

        rows.Columns()
        err = rows.Scan(&id1, &name1)
        checkErr(err)
        b := Book{1,"111",}
	bookList[name1] = &b
    }
    return bookList,nil
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }

}
