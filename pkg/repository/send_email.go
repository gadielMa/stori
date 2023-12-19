package repository

import (
	"bytes"
	"context"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
	"html/template"
	"stori/cmd/logger"
	"stori/cmd/models"
	"strings"
)

type MailRepository struct{}

func NewMailRepository() MailRepository {
	return MailRepository{}
}

func (r *MailRepository) Send(ctx context.Context, summary models.Summary) error {
	logger.Info(ctx, "starting...")

	temp := template.New("mail.html")

	var errParse error
	temp, errParse = temp.ParseFiles("mail.html")
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
	msg.SetHeader("From", "gmalagrino@frba.utn.edu.ar")
	msg.SetHeader("To", "gadiel.malagrino@gmail.com")
	msg.SetHeader("Subject", "Stori Challenge")
	msg.SetBody("text/html", buffer.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "gmalagrino@frba.utn.edu.ar", "zisa scgz irak rtuv") //TODO

	dialer.TLSConfig = &tls.Config{
		ServerName:         "smtp.gmail.com",
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
