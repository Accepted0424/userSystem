package vertify

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
)

type Verification struct {
	Email    string `json: "email"`
	Code     string `json: "code"`
	Verified bool   `json: "verified"`
}

var VerificationStore = make(map[string]Verification) // 用于保存验证码

// 生成验证码并发送邮件
func InitVerification(w http.ResponseWriter, r *http.Request) {
	code := genCode() // 生成验证码

	var ver Verification
	err := json.NewDecoder(r.Body).Decode(&ver)
	if err != nil {
		log.Fatal(err)
	}

	VerificationStore[ver.Email] = Verification{Email: ver.Email, Code: code, Verified: false}
	err = sendEmail(ver.Email, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("email已发送")
}

func genCode() string {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return code
}

// 使用SMTP发送验证码到邮箱
func sendEmail(to string, code string) error {
	// SMTP配置
	smtpHost := "smtp.163.com" // 替换为你使用的SMTP服务器
	smtpPort := "25"
	from := "13905587045@163.com"  // 发送者的邮箱
	password := "EMTvUGszDUJjW3mS" // 邮箱密码

	// 邮件内容
	subject := "邮箱验证码"
	body := fmt.Sprintf("您的验证码是：%s", code)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	// 连接SMTP服务器
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Println("发送邮件失败:", err)
		return err
	}

	log.Println("邮件发送成功")
	return nil
}

// 用户输入验证码，进行验证
func verifyCode(email string, userCode string) bool {
	verification, exists := VerificationStore[email]
	if !exists {
		log.Println("找不到该邮箱的验证信息")
		return false
	}

	if verification.Code == userCode {
		// 验证成功，更新为已验证
		VerificationStore[email] = Verification{Email: email, Code: userCode, Verified: true}
		log.Println("邮箱验证成功")
		return true
	}
	log.Println("验证码错误")
	return false
}
