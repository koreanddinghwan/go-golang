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

func (s *ReportSender) SendReport(r Report) {

	// send report
}
