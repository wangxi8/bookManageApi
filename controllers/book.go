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


