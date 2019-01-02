package controllers

import "fmt"

type MainController struct {
	BaseController
}

func (m *MainController) Get() {
	fmt.Println("=======-" + m.controllerName + "-=====")
	m.Data["Website"] = "aliio.net"
	m.Data["Email"] = "ren@aliio.com"

	m.TplName = m.controllerName + "index.tpl"
}
