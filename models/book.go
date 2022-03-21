package models

import (
    "database/sql"
    "fmt"
    "github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
)

type Book struct{
    Id int `orm:"column(id)"`
    Name string `orm:"column(name)"`
    Isbn string `orm:"column(isbn)"`
    Author string `orm:"column(author)"`
    Number int `orm:"column(number)"`
    Location string `orm:"column(location)"`
    Mobile int `orm:"column(mobile)"`
    Wechat string `orm:"column(wechat)"`
    Desc string `orm:"column(desc)"`
}

var (
	bookList map[string]*Book
)

func init(){
	// set default database
    orm.RegisterDataBase("bookManage", "mysql", "wx:wang...123@tcp(120.27.155.16:3306)/bookManage?charset=utf8", 30)

    // register model
    orm.RegisterModel(new(Book))

    bookList = make(map[string]*Book)
}

func Query(id int) (a *Book,err1 error){
	o := orm.NewOrm()

	book := Book{Id: id}
	err = o.Read(&b)

	return b, err
}

//增加数据
func insert(book *Book) (n int, err error){
    o := orm.NewOrm()

    num, err := o.Insert(&book)

    return num, err
}

//删除数据
func remove(book *Book) (n int, err error) {
    o := orm.NewOrm()

    num, err := o.Delete(&book)

    return num, err
}

//更新数据
func update(book *Book) (n int, err error) {
    o := orm.NewOrm()

    num, err := o.Update(&book)

    return num, err
}

