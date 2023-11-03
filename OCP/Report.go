package main

type FinanceReport struct {
	report string
}

func (r *FinanceReport) Report() string {

	return r.report
}

//////////////////////////////////////

type Report interface {
}

//////////////////////////////////////

type ReportSender struct{}

func (s *ReportSender) SendReport(r Report, method string) {
	switch method {
	case "email":
		// send email
	case "sms":
		// send sms
	case "fax":
	}
}
