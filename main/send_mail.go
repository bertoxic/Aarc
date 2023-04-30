package main

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
	"github.com/bertoxic/aarc/models"
)

func listenForMail(){
    log.Println(" mmmmmmmmmmm", )
	go func() {
		for {
			msg := <-app.MailChan
            log.Println(" vvvvvvvv", )
		SendMail(msg)
		}
	}()
}



func SendMail(ms models.MailData) {

	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.SendTimeout = 10 * time.Second
	server.ConnectTimeout = 10 * time.Second

	client, err := server.Connect()
    if err != nil {
        log.Println("NNNNNNNNNN",err)
    }
    log.Println(" client made sent to")
    email := mail.NewMSG()

    email.SetFrom(ms.From).AddTo(ms.To).SetSubject(ms.Subject)
    if ms.Template == "" {
        log.Println(" template about to be  sent to")
        email.SetBody(mail.TextHTML,ms.Content)
        log.Println(" template sent to", )
    }else{
        email.SetBody(mail.TextHTML,ms.Content)
    }

        err = email.Send(client)
        if err != nil {
            log.Println("xzzzzzzzzzzzz error in send mail",err)
        }else {
            log.Println("email sent")
        }
}