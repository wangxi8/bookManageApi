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
	id, _ := u.GetInt(":id")
	if id != 0 {
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
// @Param	name		query 	string	false		"书名"
// @Param	isbn		query 	string	false		"isbn"
// @Param	location    query 	string	false		"位置"
// @Param	page		query 	int	true		"页数"
// @Success 200 {object} models.Book
// @router /getList [get]
func (u *BookController) GetList() {
	name      := u.GetString(":name")
	isbn      := u.GetString(":isbn")
	location  := u.GetString(":location")
	page, _ := u.GetInt(":page")

	var book models.Book

	book.Name     = name
	book.Isbn     = isbn
	book.Location = location

	bookArr, _, err := models.QueryBookList(&book, page)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = bookArr
	}
	u.ServeJSON()
}

// @Title in
// @Description 插入一条图书
// @Param	name		query 	string	true		"书名"
// @Param	isbn		query 	string	true		"isbn"
// @Param	location    query 	string	true		"位置"
// @Param	author		query 	string	true		"作者"
// @Param	number		query 	string  false		"编号"
// @Param	mobile		query 	int	    false		"手机号"
// @Param	wechat		query 	string	true		"微信号"
// @Param	desc		query 	string	false		"描述"
// @Success 200 {object} models.Book
// @router /insert [post]
func (u *BookController) Insert() {
	name      := u.GetString(":name")
	isbn      := u.GetString(":isbn")
	location  := u.GetString(":location")
	author    := u.GetString(":author")
	number    := u.GetString(":number")
	mobile, _ := u.GetInt(":mobile")
	wechat    := u.GetString(":wechat")
	desc      := u.GetString(":desc")

	var book models.Book

	book.Name     = name
	book.Isbn     = isbn
	book.Location = location
	book.Author   = author
	book.Number   = number
	book.Mobile   = mobile
	book.Wechat   = wechat
	book.Desc     = desc

	_, err := models.Insert(&book)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = "ok"
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
	id, _ := u.GetInt(":id")

	var book models.Book

	book.Id = id
	if id != 0 {
		_, err := models.Remove(&book)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "ok"
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
// @Param	author		query 	string	true		"作者"
// @Param	number		query 	string  false		"编号"
// @Param	mobile		query 	int	    false		"手机号"
// @Param	wechat		query 	string	true		"微信号"
// @Param	desc		query 	string	false		"描述"
// @Success 200 {object} models.Book
// @Failure 403 :id is empty
// @router /update [post]
func (u *BookController) Update() {
	id, _     := u.GetInt(":id")
	name      := u.GetString(":name")
	isbn      := u.GetString(":isbn")
	location  := u.GetString(":location")
	author    := u.GetString(":author")
	number    := u.GetString(":number")
	mobile, _ := u.GetInt(":mobile")
	wechat    := u.GetString(":wechat")
	desc      := u.GetString(":desc")

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

	if id != 0 {
		_, err := models.UpdateBook(&book)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "ok"
		}
	}
	u.ServeJSON()
}


