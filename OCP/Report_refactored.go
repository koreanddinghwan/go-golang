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
/**
* 1. OCP 구현 => SRP저절로 지켜짐
 */

type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct{}

func (e *EmailSender) Send(r *Report) {
	// send email
}

type FaxSender struct {
}

func (f *FaxSender) Send(r *Report) {
	// send fax
}
