package message

type Message struct {
	MsgType 	MsgType_ 	`json:"msgtype"`
	Text 		Text_		`json:"text"`
	Link 		Link_		`json:"link"`
	Markdown	Markdown_ 	`json:"markdown"`
	ActionCard  ActionCard_ `json:"actionCard"`
	FeedCard	FeedCard_	`json:"feedCard"`
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
	Content 	string 	`json:"content"`
	At 			At_ 	`json:"at"`
}
type At_ struct {
	AtMobiles 	[]string 	`json:"atMobiles"`
	IsAtAll 	bool 		`json:"isAtAll"`
}

/*
	link类型
*/
type Link_ struct {
	Text 		string 		`json:"text"`
	Title 		string 		`json:"title"`
	PicUrl 		string 		`json:"picUrl"`
	MessageUrl 	string 		`json:"messageUrl"`
}


/*
	markdown类型
*/
type Markdown_ struct {
	Title 	string 		`json:"title"`
	Text 	string 		`json:"text"`
	At 		At_ 		`json:"at"`
}

/*
	整体跳转actionCard
*/
type ActionCard_ struct {
	Title 			string 		`json:"title"`
	Text 			string 		`json:"text"`
	BtnOrientation 	string 		`json:"btnOrientation"`
	SingleTitle	 	string 		`json:"singleTitle"`
	SingleURL 		string 		`json:"singleUrl"`
	HideAvatar 		string 		`json:"hideAvatar"`
	BtnS 			[]Btn_ 		`json:"btns"`
}
type Btn_ struct {
	Title 		string 		`json:"title"`
	ActionURL 	string 		`json:"actionURL"`
}

/*
	FeedCard类型
*/
type FeedCard_ struct {
	Links 		[]Link_ 	`json:"links"`
}