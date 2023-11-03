package main

type Report interface {
	Report() string
}

type MarketingReport struct {
}

func (m *MarketingReport) Report() string {
	return "marketing report"
}

func SendReport(r Report) {
	if _, ok := r.(*MarketingReport); ok {
		panic("MarketingReport is not allowed")
	}
}

func main() {
	var report = &MarketingReport{}
	SendReport(report)
}
