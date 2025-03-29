package models

type Channel struct{}

type JoinChannelQuery struct {
	Name      string `form:"name" required:"true"`
	IP        string `form:"-" swaggerignore:"true"`
	ChannelID string `form:"-" swaggerignore:"true"`
}
