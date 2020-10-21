package tools

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
)

type MailUser struct {
	Host string
	Server_Addr string
	User string
	Password string
}

func (u *MailUser) NewMailUser(Host, Server_Addr, User, Password string) *MailUser{
	return &MailUser{
		User : User,
		Password : Password,
		Host : Host,
		Server_Addr : Server_Addr,
	}
}

func (m *MailUser) SendEmail(to,html,fileName string)  bool{
	mime := bytes.NewBuffer(nil)
	//设置邮件
	boundary :="氕氘氚"
	mime.WriteString("From: 邮件名称<"+m.User+">\r\nTo: "+to+"\r\nSubject: 爬虫数据导出\r\nMIME-Version: 1.0\r\n")
	mime.WriteString("Content-Type: multipart/mixed; boundary="+boundary+"\r\n\r\n")

	mime.WriteString("--"+boundary+"\r\n")    //自定义邮件内容分隔符

	//邮件正文
	//html ="导出数据已通过邮件发送到您的邮箱,请下载后用excel打开"  //邮件正文
	mime.WriteString("Content-Type: text/html; charset=utf-8\r\n\r\n")  //text/html html text/plain 纯文本
	mime.WriteString(html)
	mime.WriteString("\r\n\r\n\r\n")

	//附件
	mime.WriteString("--"+boundary+"\r\n")
	mime.WriteString("Content-Type: application/vnd.ms-excel\r\n")   //application/octet-stream
	mime.WriteString("Content-Transfer-Encoding: base64\r\n")
	mime.WriteString("Content-Disposition: attachment; filename=\""+fileName+"\"")
	mime.WriteString("\r\n\r\n")

	//将文件转为base64

	//读取并编码文件内容
	//attaData, err := ioutil.ReadFile("../bapi/main.go")

	attaData, err := ioutil.ReadFile(fileName)
	if err!= nil {
		fmt.Print(err)

	}


	b :=make([]byte, base64.StdEncoding.EncodedLen(len(attaData)))
	base64.StdEncoding.Encode(b, attaData)
	mime.Write(b)
	mime.WriteString("\r\n")
	mime.WriteString("--"+boundary+"--")


	str3 := mime.String()
	auth:= smtp.PlainAuth("", m.User, m.Password, m.Host)
	errs := smtp.SendMail(m.Server_Addr,auth,m.User,[]string{to}, []byte(str3))
	if errs!= nil {
		fmt.Println(errs)
		return false
	}else{
		fmt.Println("邮件发送成功!")
	}

	return true
}
