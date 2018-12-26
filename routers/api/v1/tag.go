package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/heavenlwf/go-blog/pkg/e"
				"net/http"
	"github.com/heavenlwf/go-blog/models"
	"github.com/heavenlwf/go-blog/pkg/util"
	"github.com/heavenlwf/go-blog/pkg/config"
	"github.com/astaxie/beego/validation"
	)

func GetTags(c *gin.Context)  {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	tags := models.GetTags(util.GetPage(c), config.Conf.PageSize, maps)
	count := models.GetTagsTotal(maps)

	// TODO 输出处理

	data["list"] = tags
	data["count"] = count
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context)  {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createBy, "create_by").Message("创建人不能为空")
	valid.MaxSize(createBy, 100, "create_by").Message("创建人最长为100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTag(c *gin.Context)  {

}

func DeleteTag(c *gin.Context)  {

}