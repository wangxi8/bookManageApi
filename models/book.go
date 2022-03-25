package models

import (
    "github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/httplib"
    "errors"
)

type Book struct{
    Id int `orm:"column(id)"`
    Name string `orm:"column(name)"`
    Isbn string `orm:"column(isbn)"`
    Author string `orm:"column(author)"`
    Number string `orm:"column(number)"`
    Location string `orm:"column(location)"`
    Mobile int `orm:"column(mobile)"`
    Wechat string `orm:"column(wechat)"`
    Desc string `orm:"column(desc)"`
    Pic string `orm:"column(pic)"`
    Year string `orm:"column(year)"`
}

var (
	bookList map[string]*Book
)

const pageSize = 10

func init(){
	// set default database
    orm.RegisterDataBase("default", "mysql", "wx:wang...123@tcp(120.27.155.16:3306)/bookManage?charset=utf8", 30)

    // register model
    orm.RegisterModel(new(Book))

    orm.Debug = true

    bookList = make(map[string]*Book)
}

func Query(id int) (a *Book,err1 error){
	o := orm.NewOrm()

	book := Book{Id: id}
	err := o.Read(&book)

	return &book, err
}

func QueryBookList(b *Book, page int) (list []*Book, n int64, err error){
	o := orm.NewOrm()

	qs := o.QueryTable("book")
	cond := orm.NewCondition()

	if b.Name != "" {
		cond = cond.Or("name__contains", b.Name)
	}

	if b.Isbn != "" {
		cond = cond.Or("isbn__contains", b.Isbn)
	}

	if b.Location != "" {
		cond = cond.Or("location__contains", b.Location)
	}

	var book []*Book
	if page > 0 && pageSize > 0 {
	    qs = qs.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	num, err := qs.SetCond(cond).All(&book)

	return book, num, err
	
}

//增加数据
func Insert(book *Book) (n int64, err error){
    o := orm.NewOrm()

    num, err := o.Insert(book)

    return num, err
}

//删除数据
func Remove(book *Book) (n int64, err error) {
    o := orm.NewOrm()

    num, err := o.Delete(book)

    return num, err
}

//更新数据
func UpdateBook(book *Book) (n int64, err error) {
    o := orm.NewOrm()

    num, err := o.Update(book)

    return num, err
}

//从第三方douban获取图书信息
func getBookInfoByThird(isbn string) (b *Book, err error){
	url := "https://book.douban.com/j/subject_suggest?q=" + isbn
	req := httplib.Get(url)

	req.Header("Accept-Encoding","gzip,deflate,sdch")
	req.Header("Host","book.douban.com")
	req.Header("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

	req.ToJSON(&result)

	if result != nil {
		return nil, errors.New("dont have this book")
	}

	b := Book{Name:result[0].title, Author:result[0].author_name, Pic:result[0].pic, Year:result[0].year}

	return b, err
}

func QueryByIsbn(isbn string) (a *Book,err1 error){
	o := orm.NewOrm()

	book := Book{Isbn: isbn}
	err := o.Read(&book)

	return &book, err
}

//用户扫码查询图书信息
func getBookInfo(isbn string) (b *Book, err error){
	//优先查自己库
	book, err := this.QueryByIsbn(isbn)
	if book == nil {
		bookOther, errOther := this.getBookInfoByThird(isbn)
		return bookOther, errOther
	}

	return book, err
}

