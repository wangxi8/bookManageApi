package controllers

import (
	"bookManageApi/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type BookController struct {
	beego.Controller
}


// @Title Get
// @Description get book by id
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.Book
// @Failure 403 :id is empty
// @router /:id [get]
func (u *BookController) Get() {
	id := u.GetString(":id")
	if id != "" {
		book, err := models.Query(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = book
		}
	}
	u.ServeJSON()
}

// @Title GetList
// @Description get bookList
// @Param	name		query 	string	true		"书名"
// @Param	isbn		query 	string	true		"isbn"
// @Param	location    query 	string	true		"位置"
// @Param	page		query 	int	true		"页数"
// @Success 200 {object} models.Book
// @router /getList [get]
func (u *BookController) GetList() {
	name     := u.GetString(":name")
	isbn     := u.GetString(":isbn")
	location := u.GetString(":location")
	page     := u.GetInt(":page")

	var book models.Book

	book.Name     = name
	book.Isbn     = isbn
	book.Location = location

	if id != "" {
		bookArr, err := models.QueryBookList(book, page)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = bookArr
		}
	}
	u.ServeJSON()
}

// @Title in
// @Description 插入一条图书
// @Param	id		    query 	int	    true		"id"
// @Param	name		query 	string	true		"书名"
// @Param	isbn		query 	string	true		"isbn"
// @Param	location    query 	string	true		"位置"
// @Param	author		query 	int	    true		"作者"
// @Param	number		query 	int	    false		"编号"
// @Param	mobile		query 	int	    false		"手机号"
// @Param	wechat		query 	int	    true		"微信号"
// @Param	desc		query 	int	    false		"描述"
// @Success 200 {object} models.Book
// @router /insert [post]
func (u *BookController) Insert() {
	name     := u.GetString(":name")
	isbn     := u.GetString(":isbn")
	location := u.GetString(":location")
	author   := u.GetInt(":author")
	number   := u.GetInt(":number")
	mobile   := u.GetInt(":mobile")
	wechat   := u.GetInt(":wechat")
	desc     := u.GetInt(":desc")

	var book models.Book

	book.Name     = name
	book.Isbn     = isbn
	book.Location = location
	book.Author   = author
	book.Number   = number
	book.Mobile   = mobile
	book.Wechat   = wechat
	book.Desc     = desc

	num, err := models.Insert(book)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = 'ok'
	}

	u.ServeJSON()
}

// @Title Get
// @Description 删除书
// @Param	id		    query 	int	    true		"id"
// @Success 200 {object} models.Book
// @Failure 403 :id is empty
// @router /remove [post]
func (u *BookController) Remove() {
	id := u.GetString(":id")
	if id != "" {
		num, err := models.Remove(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = 'ok'
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description 修改图书
// @Param	id		    query 	int	    true		"id"
// @Param	name		query 	string	true		"书名"
// @Param	isbn		query 	string	true		"isbn"
// @Param	location    query 	string	true		"位置"
// @Param	author		query 	int	    true		"作者"
// @Param	number		query 	int	    false		"编号"
// @Param	mobile		query 	int	    false		"手机号"
// @Param	wechat		query 	int	    true		"微信号"
// @Param	desc		query 	int	    false		"描述"
// @Success 200 {object} models.Book
// @Failure 403 :id is empty
// @router /update [post]
func (u *BookController) Update() {
	id       := u.GetString(":id")
	name     := u.GetString(":name")
	isbn     := u.GetString(":isbn")
	location := u.GetString(":location")
	author   := u.GetInt(":author")
	number   := u.GetInt(":number")
	mobile   := u.GetInt(":mobile")
	wechat   := u.GetInt(":wechat")
	desc     := u.GetInt(":desc")

	var book models.Book

	book.Id       = id
	book.Name     = name
	book.Isbn     = isbn
	book.Location = location
	book.Author   = author
	book.Number   = number
	book.Mobile   = mobile
	book.Wechat   = wechat
	book.Desc     = desc

	if id != "" {
		num, err := models.Update(book)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = 'ok'
		}
	}
	u.ServeJSON()
}


