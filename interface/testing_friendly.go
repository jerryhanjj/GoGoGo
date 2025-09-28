package interfacego

import "fmt"

// 实际应用场景2：让代码更易测试
type EmailSender interface {
	Send(to, subject, body string) error
}

type SMTPSender struct {
	server string
}

func (s SMTPSender) Send(to, subject, body string) error {
	fmt.Printf("通过SMTP服务器 %s 发送邮件给 %s\n", s.server, to)
	fmt.Printf("主题: %s\n", subject)
	// 实际发送邮件的代码...
	return nil
}

// 测试用的mock实现
type MockEmailSender struct {
	SentEmails []Email
}

type Email struct {
	To      string
	Subject string
	Body    string
}

func (m *MockEmailSender) Send(to, subject, body string) error {
	fmt.Printf("Mock: 记录邮件发送 -> %s\n", to)
	m.SentEmails = append(m.SentEmails, Email{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	return nil
}

// 业务服务
type NotificationService struct {
	emailSender EmailSender
}

func (n NotificationService) WelcomeUser(userEmail, userName string) {
	subject := "欢迎加入我们！"
	body := fmt.Sprintf("亲爱的 %s，欢迎！", userName)
	n.emailSender.Send(userEmail, subject, body)
}

func (n NotificationService) SendAlert(adminEmail, message string) {
	subject := "系统警报"
	n.emailSender.Send(adminEmail, subject, message)
}

func TestingFriendlyDemo() {
	fmt.Println("=== 接口让测试更容易 ===")

	// 生产环境使用
	fmt.Println("1. 生产环境:")
	smtpSender := SMTPSender{server: "smtp.company.com"}
	notificationService := NotificationService{emailSender: smtpSender}
	notificationService.WelcomeUser("user@example.com", "张三")

	fmt.Println("\n2. 测试环境:")
	// 测试时使用mock，不会真的发邮件
	mockSender := &MockEmailSender{}
	testService := NotificationService{emailSender: mockSender}

	testService.WelcomeUser("test@example.com", "测试用户")
	testService.SendAlert("admin@example.com", "系统出错了")

	// 验证测试结果
	fmt.Printf("测试验证：共发送了 %d 封邮件\n", len(mockSender.SentEmails))
	for i, email := range mockSender.SentEmails {
		fmt.Printf("  邮件%d: %s -> %s\n", i+1, email.Subject, email.To)
	}

	fmt.Println("\n优势：测试时不需要真的发邮件！")
}
