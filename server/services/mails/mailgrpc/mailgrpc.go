package mailgrpc

import (
	"context"
	"errors"
	"io/ioutil"
	"os"

	"google.golang.org/grpc/grpclog"
	"gopkg.in/gomail.v2"
)

type Server struct{}

const (
	VERIFY string = "verify"
	ACTIVATED string = "activated"
	RECOVER string = "recover"
	UPDATE string = "update"
	UPDATEVERIFY string ="update verify"
)

func getDialer() *gomail.Dialer {
	return gomail.NewDialer("in-v3.mailjet.com",587,"61504278c13d9beebebb5da0b7b89fce","97f247ab677a6995df02d58a53479bae")
}

func getTemplate(path string) ([]byte,error) {
	template,err := os.Open(path)
		if err != nil {
			grpclog.Errorln(err.Error())
			return nil, err
		}
		bs,err := ioutil.ReadAll(template)
		if err != nil {
			grpclog.Errorln(err.Error())
			return nil, err
		}
		return bs,nil
}

func formatReq(m *gomail.Message,to string,subject string) *gomail.Message {
	m.SetHeader("From","lois@pixelogi.com")
	m.SetHeader("To",to)
	//m.SetAddressHeader("Cc","loisbibehenonga@gmail.com","lois")
	m.SetHeader("Subject",subject)
	return m
}

func NewMailHandler() *Server {
	return &Server{}
}

func (s *Server) Send(ctx context.Context,req *Mailreq) (*Mailres,error) {
	d := getDialer()
	m := gomail.NewMessage()
	switch req.GetSubject() {
	case VERIFY:
		m := formatReq(m,req.GetTo(),"E-commerce, Activate your account")
		bs,err := getTemplate("./services/mails/templates/activate.html")
		if err != nil {
			grpclog.Errorln(err)
			return nil,err
		}
		m.SetBody("text/html",string(bs))
		m.Attach("./services/mails/LICENSE")
	case RECOVER:
		m := formatReq(m,req.GetTo(),"E-commerce, Recover your password")
		m.SetBody("text/html","<h1>Hello man</h1> <p>this is your boss lois</p>")
	case UPDATE:
		m := formatReq(m,req.GetTo(),"E-commerce, Recover your password")
		m.SetBody("text/html","<h1>Hello man</h1> <p>this is your boss lois</p>")
	case ACTIVATED:
		m := formatReq(m,req.GetTo(),"E-commerce, Recover your password")
		m.SetBody("text/html","<h1>Hello man</h1> <p>this is your boss lois</p>")
	case UPDATEVERIFY:
		m := formatReq(m,req.GetTo(),"E-commerce, Update your email address needs reverify")
		m.SetBody("text/html","<h1>Hello man</h1> <p>Reverify your email address please</p>")
	default:
		return nil,errors.New("this action is not available yet")
	}
	if err := d.DialAndSend(m); err != nil {
		grpclog.Errorln(err.Error())
		return nil,err
	}
	return &Mailres{},nil
}