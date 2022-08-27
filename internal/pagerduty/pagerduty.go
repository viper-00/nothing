package pagerduty

type Body struct {
	Type    string `json:"type"`
	Details string `json:"details"`
}

type Service struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type PagerDutyIncident struct {
	Id      string  `json:"id"`
	Type    string  `json:"type"`
	Title   string  `json:"title"`
	Urgency string  `json:"urgency"`
	Status  string  `json:"status"`
	Body    Body    `json:"body"`
	Service Service `json:"service"`
}

type Incident struct {
	Incident PagerDutyIncident `json:"incident"`
}

func CreateIncident(incident Incident) (string, error) {
	return "", nil
}

func UpdateIncident(id string) error {
	return nil
}
