package slack

import "time"

type Payload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

type Attachment struct {
	Fallback   string   `json:"fallback"`
	Color      string   `json:"color"`
	Pretext    string   `json:"pretext"`
	AuthorName string   `json:"author_name"`
	AuthorLink string   `json:"author_link"`
	AuthorIcon string   `json:"author_icon"`
	Title      string   `json:"title"`
	TitleLink  string   `json:"title_link"`
	Text       string   `json:"text"`
	Fields     []Field  `json:"fields"`
	ImageURL   string   `json:"image_url"`
	ThumbURL   string   `json:"thumb_url"`
	Footer     string   `json:"footer"`
	FooterIcon string   `json:"footer_icon"`
	Timestamp  int64    `json:"ts"`
	MarkdownIn []string `json:"markdown_in"`
}
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func (ap *Payload) AddAttachment(a Attachment) {
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

func (a *Attachment) AddTimestamp(t time.Time) *Attachment {
	a.Timestamp = t.Unix()
	return a
}
