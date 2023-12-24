package repository

import (
	"bytes"
	"context"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
	"html/template"
	"os"
	"stori/cmd/logger"
	"stori/cmd/models"
	"strings"
)

const (
	HtmlMail  = "mail.html"
	SmtpGmail = "smtp.gmail.com"
)

type MailRepository struct{}

func NewMailRepository() MailRepository {
	return MailRepository{}
}

func (r *MailRepository) Send(ctx context.Context, summary models.Summary) error {
	logger.Info(ctx, "starting...")

	temp := template.New(HtmlMail)

	var errParse error
	temp, errParse = temp.ParseFiles(HtmlMail)
	if errParse != nil {
		logger.Error(ctx, errParse.Error())
		return errParse
	}

	var buffer bytes.Buffer
	if errExec := temp.Execute(&buffer, buildMail(summary)); errExec != nil {
		logger.Error(ctx, errExec.Error())
		return errExec
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("MAIL_FROM"))
	msg.SetHeader("To", os.Getenv("MAIL_TO"))
	msg.SetHeader("Subject", "Stori Challenge")
	msg.SetBody("text/html", buffer.String())

	dialer := gomail.NewDialer(SmtpGmail, 587, os.Getenv("MAIL_FROM"), os.Getenv("MAIL_FROM_PASSWORD"))

	dialer.TLSConfig = &tls.Config{
		ServerName:         SmtpGmail,
		InsecureSkipVerify: false,
	}

	if errSend := dialer.DialAndSend(msg); errSend != nil {
		logger.Error(ctx, errSend.Error())
		return errSend
	}

	return nil
}

func buildMail(summary models.Summary) mail {
	monthlyTransactions := strings.Split(strings.Trim(summary.MonthlyTransactions, "[]"), " ")

	return mail{
		Balance:   summary.Balance,
		January:   monthlyTransactions[0],
		February:  monthlyTransactions[1],
		March:     monthlyTransactions[2],
		April:     monthlyTransactions[3],
		May:       monthlyTransactions[4],
		June:      monthlyTransactions[5],
		July:      monthlyTransactions[6],
		August:    monthlyTransactions[7],
		September: monthlyTransactions[8],
		October:   monthlyTransactions[9],
		November:  monthlyTransactions[10],
		December:  monthlyTransactions[11],
		DebitAvg:  summary.AvgDebit,
		CreditAvg: summary.AvgCredit,
	}
}
