package helper

import (
	"example/go_fakeit/model"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateUserData(DB *gorm.DB, count *int) {
	for i := 0; i < *count; i++ {
		var user model.User
		gofakeit.Struct(&user)
		user.CreatedAt = gofakeit.DateRange(time.Date(2015, 05, 13, 0, 0, 0, 0, time.UTC), time.Now())
		createUserQuery := "INSERT INTO users (username, email, password, phone_number, description, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
		DB.Exec(createUserQuery, user.Username, user.Email, user.Password, user.PhoneNumber, user.Description, user.CreatedAt.Format("2006-01-02"))
	}

	ClearScreen()
	fmt.Println(*count, "users successfully added to the database!")
}

func GetRandomUser(DB *gorm.DB) (uuid.UUID, time.Time) {
	var userID uuid.UUID
	var userCreatedAt time.Time
	err := DB.Table("users").Select("user_id, created_at").Order("RANDOM()").Limit(1).Row().Scan(&userID, &userCreatedAt)
	if err != nil {
		panic(err)
	}
	return userID, userCreatedAt
}

func GenerateGuildData(DB *gorm.DB, count *int) {
	for i := 0; i < *count; i++ {
		var guild model.Guild
		var userCreatedAt time.Time
		gofakeit.Struct(&guild)
		guild.GuildOwner, userCreatedAt = GetRandomUser(DB)
		guild.CreatedAt = gofakeit.DateRange(userCreatedAt, time.Now())
		createGuildQuery := "INSERT INTO guilds (guild_name, guild_owner, created_at) VALUES ($1, $2, $3) RETURNING guild_id"
		addOwnerToGuild := "INSERT INTO members (guild_id, user_id, joined_at) VALUES ($1, $2, $3)"
		DB.Raw(createGuildQuery, guild.GuildName, guild.GuildOwner, guild.CreatedAt.Format("2006-01-02")).Scan(&guild.GuildID)
		DB.Exec(addOwnerToGuild, guild.GuildID, guild.GuildOwner, guild.CreatedAt.Format("2006-01-02"))
	}

	ClearScreen()
	fmt.Println(*count, "guilds successfully added to the database!")
}

func GetRandomGuild(DB *gorm.DB) (uuid.UUID, time.Time) {
	var guildID uuid.UUID
	var guildCreatedAt time.Time
	err := DB.Table("guilds").Select("guild_id, created_at").Order("RANDOM()").Limit(1).Row().Scan(&guildID, &guildCreatedAt)
	if err != nil {
		panic(err)
	}
	return guildID, guildCreatedAt
}

func GetRandomChannelType(DB *gorm.DB) int8 {
	var channelType int8
	err := DB.Table("channel_types").Select("channel_type_id").Order("RANDOM()").Limit(1).Row().Scan(&channelType)
	if err != nil {
		panic(err)
	}
	return channelType
}

func GenerateChannelData(DB *gorm.DB, count *int) {
	for i := 0; i < *count; i++ {
		var channel model.Channel
		gofakeit.Struct(&channel)
		channel.GuildID, _ = GetRandomGuild(DB)
		channel.ChannelType = GetRandomChannelType(DB)
		createChannelQuery := "INSERT INTO channels (channel_name, channel_type, guild_id) VALUES ($1, $2, $3)"
		DB.Exec(createChannelQuery, channel.ChannelName, channel.ChannelType, channel.GuildID)
	}

	ClearScreen()
	fmt.Println(*count, "channels successfully added to the database!")
}

func GetRandomChannel(DB *gorm.DB) (uuid.UUID, uuid.UUID) {
	var channelID, guildID uuid.UUID
	err := DB.Table("channels").Select("channel_id, guild_id").Order("RANDOM()").Limit(1).Row().Scan(&channelID, &guildID)
	if err != nil {
		panic(err)
	}
	return channelID, guildID
}

func GenerateMemberData(DB *gorm.DB, count *int) {
	for i := 0; i < *count; i++ {
		var member model.Member
		var guildCreatedAt time.Time
		member.GuildID, guildCreatedAt = GetRandomGuild(DB)
		member.UserID, _ = GetRandomUser(DB)
		member.JoinedAt = gofakeit.DateRange(guildCreatedAt, time.Now())
		createMemberQuery := "INSERT INTO members (guild_id, user_id, joined_at) VALUES ($1, $2, $3)"
		DB.Exec(createMemberQuery, member.GuildID, member.UserID, member.JoinedAt.Format("2006-01-02"))
	}

	ClearScreen()
	fmt.Println(*count, "members successfully added to the database!")
}

func GenerateMessageData(DB *gorm.DB, count *int) {
	for i := 0; i < *count; i++ {
		var message model.Message
		var guildID uuid.UUID
		var guildCreatedAt time.Time
		message.SenderUser, _ = GetRandomUser(DB)
		message.MsgDestination, guildID = GetRandomChannel(DB)
		DB.Raw("SELECT created_at FROM guilds WHERE guild_id = $1", guildID).Scan(&guildCreatedAt)
		message.SentAt = gofakeit.DateRange(guildCreatedAt, time.Now())
		createMessageQuery := "INSERT INTO messages (sender_user, msg_destination, sent_at) VALUES ($1, $2, $3)"
		DB.Exec(createMessageQuery, message.SenderUser, message.MsgDestination, message.SentAt)
	}

	ClearScreen()
	fmt.Println(*count, "messages successfully added to the database!")
}
