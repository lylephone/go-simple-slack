package slack

type AttachmentPayload struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	Pretext    string  `json:"pretext"`
	AuthorName string  `json:"author_name"`
	AuthorLink string  `json:"author_link"`
	AuthorIcon string  `json:"author_icon"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Text       string  `json:"text"`
	Fields     []Field `json:"fields"`
	ImageURL   string  `json:"image_url"`
	ThumbURL   string  `json:"thumb_url"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
	Ts         int     `json:"ts"`
}
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func (ap *AttachmentPayload) AddAttachment(a Attachment) {
	ap.Attachments = append(ap.Attachments, a)
}

func (a *Attachment) AddField(title, value string, short bool) *Attachment {
	f := Field{
		Title: title,
		Value: value,
		Short: short,
	}
	a.Fields = append(a.Fields, f)
	return a
}
