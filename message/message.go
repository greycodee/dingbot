package message

type Message struct {
	MsgType 	MsgType_ 	`json:"msgtype"`
	Text 		Text_		`json:"text"`
	Link 		Link_		`json:"link"`
}
type MsgType_ string

const(
	TextStr 			MsgType_ = "text"
	LinkStr 			MsgType_ = "link"
	MarkdownStr 		MsgType_ = "markdown"
	ActionCardStr 		MsgType_ = "actionCard"
	FeedCardStr 		MsgType_ = "feedCard"
)

/*
	text类型
*/
type Text_ struct {
	Content string `json:"content"`
	At At_ `json:"at"`
}
type At_ struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll bool `json:"isAtAll"`
}

/*
	link类型
*/
type Link_ struct {
	Text string `json:"text"`
	Title string `json:"title"`
	PicUrl string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}
