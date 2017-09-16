// ubuntu smtp 搭建：
//   1. https://linux.cn/article-8071-1-rel.html
//   2. https://linux.cn/article-8077-1.html
// 源码参考:
//   1. https://github.com/golang/go/wiki/SendingMail
package main

import (
	"bytes"
	"log"
	"net/smtp"
)

func main() {
	c, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.Mail("sender@example.com")
	c.Rcpt("recipient@example.com")

	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	buf := bytes.NewBufferString("Subject: mail server test\r\nThis is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
