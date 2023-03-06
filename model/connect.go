package model

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"net"
)

//定义一个包含ssh.Client指针，以及一个成员函数
type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(context context.Context, addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}
func Connect() {
	// 一个ClientConfig指针,指向的对象需要包含ssh登录的信息
	config := &ssh.ClientConfig{
		User: "root", //我使用的是root用户，就写"root"
		Auth: []ssh.AuthMethod{
			ssh.Password("dengzuxuan20010404!"), //用户的密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个我也不太清楚，大概是做服务端验证的，按这么写就行
	}

	client, err := ssh.Dial("tcp", "8.140.38.47:22", config) //xxx那段替换为你服务器的的IP地址
	if err != nil {
		panic("连接失败") //抛出异常
	} else {
		fmt.Println("连接成功")
	}

	//用ssh连接 client指针作为参数注册ViaSSHDialer
	mysql.RegisterDialContext("mysql+tcp", (&ViaSSHDialer{client}).Dial)

}
