package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Username    string `fake:"{gamertag}"`
	Email       string `fake:"{email}"`
	Password    string `fake:"{password:true,true,false,false,false,8}"`
	PhoneNumber string `fake:"{phoneformatted}"`
	Description string `fake:"{hackerphrase}"`
	CreatedAt   time.Time
}

type Guild struct {
	GuildID    uuid.UUID
	GuildName  string `fake:"{company}"`
	GuildOwner uuid.UUID
	CreatedAt  time.Time
}

type Channel struct {
	ChannelName string `fake:"{hipstersentence:2}"`
	ChannelType int8
	GuildID     uuid.UUID
}

type Member struct {
	GuildID  uuid.UUID
	UserID   uuid.UUID
	JoinedAt time.Time
}

// type FriendList struct {
// 	UserID   uuid.UUID
// 	FriendID uuid.UUID
// }

// type DirectMessage struct {
// 	FirstUser  uuid.UUID
// 	SecondUser uuid.UUID
// }

// type Message struct {
// 	SenderUser     uuid.UUID
// 	MsgDestination uuid.UUID
// }
