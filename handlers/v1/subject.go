package handler

import (
	handler "a6-api/handlers"
	model "a6-api/models/v1/photo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var subjectModel *model.Subject

//SubjectList return subject list
func SubjectList(c *gin.Context) {
	var params = model.SubjectListParams{}

	if c.ShouldBind(&params) != nil {
		handler.ErrorMsg(c, http.StatusUnprocessableEntity, "")
	}
	data, pagin, notFound := subjectModel.List(params)
	if notFound {
		handler.ErrorMsg(c, http.StatusNotFound, "not found record")
		return
	}
	handler.Ok(c, data, pagin)
}

//SubjectDetail return subject detail
func SubjectDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.ErrorMsg(c, http.StatusUnprocessableEntity, "The param \"id\" type of uint64")
		return
	}

	data, notFound := subjectModel.Detail(id)

	if notFound {
		handler.ErrorMsg(c, http.StatusNotFound, "Not found record, Pleases check params")
		return
	}

	handler.Ok(c, data, map[string]interface{}{})
}