package main

import (
	"log"
	"net/smtp"
)

func template() string {
	temp := `<html>
		<body>
			<h1>Hi this is test</h1>
			<p style='color:#ccc;'>Hello jsjsjsjajajajaja</p>
			<span style='color:green;'>HHHiiiza</span>
		</body>
	</html>
	`
	return temp
}

func sendMail(body string) {
	from := "example@gmail.com"
	pass := "..."
	to := "receiver.somewhere@gmail.com"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := "From: mr.Example\n" +
		"To: somewhere\n" +
		"Subject: Hello hi and goodbye for test\n" +
		mime + body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	} else {
		log.Println("mail is send ja")
	}
}

func main() {
	sendMail(template())
}
