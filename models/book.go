package models

import (
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

const pageSize = 10

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

func QueryBookList(b *Book, page int) (list []*Book, n int, err error){
	o := orm.NewOrm()

	qs = o.QueryTable("book")
	cond := orm.NewCondition()

	if b.Name != "" {
		cond.Or("name__iexact", b.Name)
	}

	if b.Isbn != "" {
		cond.Or("isbn__iexact", b.Isbn)
	}

	if b.Location != "" {
		cond.Or("location__iexact", b.Location)
	}

	var book []*Book
	if page > 0 && pageSize > 0 {
	    qs.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	num, err := qs.SetCond(cond).All(&book)

	return book, num, err
	
}

//增加数据
func Insert(book *Book) (n int, err error){
    o := orm.NewOrm()

    num, err := o.Insert(&book)

    return num, err
}

//删除数据
func Remove(book *Book) (n int, err error) {
    o := orm.NewOrm()

    num, err := o.Delete(&book)

    return num, err
}

//更新数据
func Update(book *Book) (n int, err error) {
    o := orm.NewOrm()

    num, err := o.Update(&book)

    return num, err
}

