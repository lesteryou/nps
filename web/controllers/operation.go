package controllers

import (
	"ehang.io/nps/server"
	"fmt"
)

type OperationController struct {
	BaseController
}

func (s *OperationController) List() {
	if s.Ctx.Request.Method == "GET" {
		s.Data["menu"] = "operation"
		s.SetInfo("operation")
		s.display("operation/list")
		return
	}

	start, length := s.GetAjaxParams()
	list, cnt := server.GetOperationList(start, length, s.getEscapeString("search"), s.getEscapeString("sort"), s.getEscapeString("order"))
	cmd := make(map[string]interface{})
	s.AjaxTable(list, cnt, cnt, cmd)
	fmt.Println(list, cmd)
}
