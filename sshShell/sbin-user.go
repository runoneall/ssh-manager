package sshShell

import (
	"fmt"
	"slices"
	"ssh-manager/helper"
	"ssh-manager/vars"

	"github.com/akamensky/argparse"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func sbinUserSave(t *term.Terminal) {
	if umanager.SaveToJson(vars.FILE_USER_CONFIG) {
		fmt.Fprintln(t, "用户保存成功!")
	} else {
		fmt.Fprintln(t, "用户保存失败!")
	}
}

func sbinUserAdd(s ssh.Session, t *term.Terminal, arg []string) {

	// 验证管理员
	if !umanager.IsAdmin(s.User()) {
		fmt.Fprintln(t, "你不能这么做!")
		return
	}

	// 获取用户输入
	name := TokenAt(arg, 1)
	if name == "" {
		fmt.Fprintf(t, "请输入用户名: %s <username>\n", arg[0])
		return
	}

	// 检查用户
	if umanager.IsExist(name) {
		fmt.Fprintln(t, "用户已存在!")
		return
	}

	// 创建用户
	umanager.AddUser(name, "", false, []string{}, true)
	fmt.Fprintln(t, "用户创建成功!")
	sbinUserSave(t)

}

func sbinUserManage(s ssh.Session, t *term.Terminal, arg []string) {

	// 验证管理员
	if !umanager.IsAdmin(s.User()) {
		fmt.Fprintln(t, "你不能这么做!")
		return
	}

	// 创建解析器
	parser := argparse.NewParser("user", "管理用户")
	parser.DisableHelp()

	// 添加help选项
	isHelp := parser.Flag("h", "help", &argparse.Options{Help: "显示帮助信息"})

	// 获取目标用户
	nameStr := parser.String("t", "target", &argparse.Options{Required: true, Help: "目标用户名"})

	// 更改密码
	newPassword := parser.String(
		"p",
		"password",
		&argparse.Options{
			Required: false,
			Help:     "新的密码",
			Default:  "",
		},
	)

	// 更改管理员权限
	isAdmin := parser.Selector(
		"a",
		"admin",
		[]string{"y", "n"},
		&argparse.Options{
			Required: false,
			Help:     "是(y) 否(n) 为管理员",
		},
	)

	// 更改禁用状态
	isDisable := parser.Selector(
		"d",
		"disable",
		[]string{"y", "n"},
		&argparse.Options{
			Required: false,
			Help:     "是(y) 否(n) 禁用用户",
		},
	)

	// 添加服务器
	serverAddList := parser.StringList(
		"s",
		"sadd",
		&argparse.Options{
			Required: false,
			Help:     "添加服务器访问",
		},
	)

	// 删除服务器
	serverDelList := parser.StringList(
		"",
		"sdel",
		&argparse.Options{
			Required: false,
			Help:     "删除服务器访问",
		},
	)

	// 删除用户
	isDelete := parser.Flag(
		"D",
		"del",
		&argparse.Options{
			Required: false,
			Help:     "删除用户",
		},
	)

	// 打印用户信息
	isInfo := parser.Flag(
		"i",
		"info",
		&argparse.Options{
			Required: false,
			Help:     "打印用户信息",
		},
	)

	// 未知参数
	err := parser.Parse(arg)
	if err != nil {
		fmt.Fprint(t, parser.Usage(err))
		return
	}

	// 显示帮助信息
	if *isHelp {
		fmt.Fprint(t, parser.Usage(nil))
		return
	}

	// 获取目标用户
	user, ok := umanager.GetUser(*nameStr)
	if !ok {
		fmt.Fprintln(t, "用户不存在!")
		return
	}

	// 操作完后退出目标登录
	logoutAll := func(user string) {
		for _, store := range smanager.GetUserSessions(user) {
			store.Session.Exit(0)
		}
	}

	// 更改密码
	if *newPassword != "" {
		user.Password = *newPassword
		fmt.Fprintln(t, "密码修改成功!")
	}

	// 更改管理员权限
	if *isAdmin != "" {
		user.IsAdmin = *isAdmin == "y"
		fmt.Fprintln(t, "管理员权限修改成功!")
	}

	// 更改禁用状态
	if *isDisable != "" {
		user.IsDisable = *isDisable == "y"
		fmt.Fprintln(t, "禁用状态修改成功!")
	}

	// 添加服务器
	if len(*serverAddList) > 0 {
		for _, server := range *serverAddList {
			if !slices.Contains(user.Servers, server) {
				user.Servers = append(user.Servers, server)
			}
		}
		fmt.Fprintln(t, "服务器添加成功!")
	}

	// 删除服务器
	if len(*serverDelList) > 0 {
		for _, server := range *serverDelList {
			if slices.Contains(user.Servers, server) {
				index := slices.Index(user.Servers, server)
				user.Servers = helper.RemoveUnordered(
					user.Servers, index,
				)
			}
		}
		fmt.Fprintln(t, "服务器删除成功!")
	}

	// 删除用户
	if *isDelete {
		logoutAll(user.Name)
		umanager.DeleteUser(user.Name)
		fmt.Fprintln(t, "用户删除成功!")
		sbinUserSave(t)
		return
	}

	// 打印用户信息
	if *isInfo {
		fmt.Fprintln(t, "用户名:", user.Name)
		fmt.Fprintln(t, "密码:", user.Password)
		fmt.Fprintln(t, "管理员权限:", user.IsAdmin)
		fmt.Fprintln(t, "禁用状态:", user.IsDisable)
		fmt.Fprintln(t, "服务器访问:", user.Servers)
		return
	}

	// 保存用户信息
	umanager.AddUser(
		user.Name,
		user.Password,
		user.IsAdmin,
		user.Servers,
		user.IsDisable,
	)
	sbinUserSave(t)
	logoutAll(user.Name)

}

func sbinUserList(s ssh.Session, t *term.Terminal, arg []string) {

	// 验证管理员
	if !umanager.IsAdmin(s.User()) {
		fmt.Fprintln(t, "你不能这么做!")
		return
	}

	// 打印用户列表
	fmt.Fprintln(t, "用户名列表:")
	for _, name := range umanager.ListUser() {
		fmt.Fprintf(t, "  %s\n", name)
	}

}
