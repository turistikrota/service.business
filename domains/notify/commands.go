package notify

type SendSpecialEmailCmd struct {
	Email          string `json:"email"`
	Template       string `json:"template"`
	Subject        string `json:"subject"`
	Content        string `json:"content"`
	TemplateParams any    `json:"templateParams"`
	Translate      bool   `json:"translate"`
	Locale         string `json:"locale"`
}

type NotifySendToAllChannelsCmd struct {
	ActorName    string `json:"actorName"`
	Subject      string `json:"subject"`
	Content      string `json:"content"`
	Locale       string `json:"locale"`
	Template     string `json:"template,omitempty"`
	TemplateData any    `json:"templateData,omitempty"`
	Translate    bool   `json:"translate"`
}
