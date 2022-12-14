package controller

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"userManagementV2/model"
	"userManagementV2/service"
)

//负责转发到Service层
//显示主菜单
//接收用户输入
//调用userSerivce完成用户管理

type UserController struct {
	arg     string
	service service.UserService
}

func NewUserController(arg string, service service.UserService) *UserController {
	return &UserController{
		arg:     arg,
		service: service,
	}
}

func (c *UserController) add() {
	fmt.Println("添加用户")
	fmt.Print("用户名：")
	var name string
	fmt.Scanln(&name)
	fmt.Print("电话：")
	var phone string
	fmt.Scanln(&phone)
	fmt.Print("地址：")
	var address string
	fmt.Scanln(&address)
	user := model.User{
		Name:    name,
		Phone:   phone,
		Address: address,
	}
	if c.service.CreateUser(user) == nil {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}
}
//todo 删查改
//func (c *UserController) delete() {
//	if c.userService.GetUserNum() == 0 {
//		fmt.Println("用户管理列表为空")
//		return
//	}
//	fmt.Println("删除用户")
//	fmt.Print("用户编号：")
//	var userId int
//	fmt.Scanln(&userId)
//	idx, err := c.userService.HasUser(userId)
//	if err != nil {
//		fmt.Printf("删除失败：%s\n", err.Error())
//		return
//	}
//	user := c.userService.GetUser(idx)
//	c.tableFormat(user)
//	fmt.Print("确认删除：（y/n）")
//	var cfm string
//	fmt.Scanln(&cfm)
//	if cfm == "y" {
//		if c.userService.Delete(idx) {
//			fmt.Println("删除成功")
//		} else {
//			fmt.Println("删除失败")
//		}
//	} else {
//		fmt.Println("不删除")
//	}
//}
//
//func (c *UserController) modify() {
//	if c.userService.GetUserNum() == 0 {
//		fmt.Println("用户管理列表为空")
//		return
//	}
//	fmt.Println("修改用户")
//	fmt.Print("用户编号：")
//	var userId int
//	fmt.Scanln(&userId)
//	idx, err := c.userService.HasUser(userId)
//	if err != nil {
//		fmt.Printf("修改失败：%s\n", err.Error())
//		return
//	}
//	user := c.userService.GetUser(idx)
//	v.tableFormat(user)
//	fmt.Print("确认修改：（y/n）")
//	var cfm string
//	fmt.Scanln(&cfm)
//	if cfm == "y" {
//		fmt.Print("用户名：")
//		var name string
//		fmt.Scanln(&name)
//		fmt.Print("电话：")
//		var phone string
//		fmt.Scanln(&phone)
//		fmt.Print("地址：")
//		var address string
//		fmt.Scanln(&address)
//		mUser := model.User{
//			Name:    name,
//			Phone:   phone,
//			Address: address,
//		}
//		if c.userService.Modify(idx, mUser) {
//			fmt.Println("修改成功")
//		} else {
//			fmt.Println("修改失败")
//		}
//	} else {
//		fmt.Println("不修改")
//	}
//}
//
//func (c *UserController) query() {
//	if c.userService.GetUserNum() == 0 {
//		fmt.Println("用户管理列表为空")
//		return
//	}
//	fmt.Println("搜索用户")
//	fmt.Print("请输入关键字：")
//	var keyword string
//	fmt.Scanln(&keyword)
//	matchUsers := c.userService.Query(keyword)
//	if len(matchUsers) != 0 {
//		c.tableFormat(matchUsers)
//	} else {
//		fmt.Println("无匹配用户")
//	}
//}
//
func (c *UserController) list() {
	users := c.service.ListUser()
	c.tableFormat(users)
}

func (c *UserController) menu() {
	fmt.Println("*********用户管理系统*********")
	fmt.Printf("%15s\n", "a）添加用户")
	fmt.Printf("%15s\n", "m）修改用户")
	fmt.Printf("%15s\n", "d）删除用户")
	fmt.Printf("%15s\n", "l）用户列表")
	fmt.Printf("%15s\n", "q）搜索用户")
	fmt.Printf("%15s\n", "h）帮助信息")
	fmt.Printf("%15s\n", "exit）退出系统")
	fmt.Println("****************************")
}

func (c *UserController) tableFormat(data []*model.User) {
	fmtData := [][]string{}
	for _, v := range data {
		fmtData = append(fmtData, []string{strconv.Itoa(v.Id), v.Name, v.Phone, v.Address})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"用户编号", "用户名", "电话号码", "联系地址"})
	for _, v := range fmtData {
		table.Append(v)
	}
	table.Render()
}

func (c *UserController) MainMenu() {
	c.menu()
	for {
		fmt.Print("输入菜单选项：")
		fmt.Scanln(&c.arg)
		switch c.arg {
		case "a":
			c.add()
		//case "m":
		//	v.modify()
		//case "d":
		//	v.delete()
		case "l":
			c.list()
		//case "q":
		//	v.query()
		case "h":
			c.menu()
		case "exit":
			fmt.Println("退出用户管理")
			return
		default:
			fmt.Println("非法输入")
		}
	}
}
