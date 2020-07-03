package xsmtp

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

type TSMTP struct {
	Address  string
	Username string
	Password string
}

func New(address, username, password string) *TSMTP {
	return &TSMTP{
		Address:  address,
		Username: username,
		Password: password,
	}
}

func (s *TSMTP) SendMail(from, tos, subject, body string, contentType ...string) error {
	if s.Address == "" {
		return fmt.Errorf("address is necessary")
	}

	hp := strings.Split(s.Address, ":")
	if len(hp) != 2 {
		return fmt.Errorf("address format error")
	}

	arr := strings.Split(tos, ";")
	count := len(arr)
	safeArr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		if arr[i] == "" {
			continue
		}
		safeArr = append(safeArr, arr[i])
	}

	if len(safeArr) == 0 {
		return fmt.Errorf("tos invalid")
	}

	tos = strings.Join(safeArr, ";")
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

	header := make(map[string]string)
	header["From"] = from
	header["To"] = tos
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(subject)))
	header["MIME-Version"] = "1.0"

	ct := "text/plain; charset=UTF-8"
	if len(contentType) > 0 && contentType[0] == "html" {
		ct = "text/html; charset=UTF-8"
	}

	header["Content-Type"] = ct
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))

	auth := smtp.PlainAuth("", s.Username, s.Password, hp[0])
	return smtp.SendMail(s.Address, auth, from, strings.Split(tos, ";"), []byte(message))
}
