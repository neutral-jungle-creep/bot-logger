package domain

type Message struct {
	MessageId     string
	Date          string
	IsEdit        bool
	V4Data        string
	MessageSender User
	MessageText   MessageText
}

type MessageText struct {
	Query    string
	Problem  string
	Cause    string
	Solution string
	Source   string
}
