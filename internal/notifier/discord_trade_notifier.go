package notifier

import (
	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	"github.com/bwmarrin/discordgo"
)

type DiscordTradeNotifier struct {
	session   *discordgo.Session
	channelID string
}

func NewDiscordTradeNotifier(cfg *config.Config) (*DiscordTradeNotifier, error) {
	session, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		return nil, err
	}

	channel, err := session.UserChannelCreate(cfg.DiscordUserID)
	if err != nil {
		return nil, err
	}

	return &DiscordTradeNotifier{
		session:   session,
		channelID: channel.ID,
	}, nil
}

func (notifier *DiscordTradeNotifier) GetBotName() string {
	bot, _ := notifier.session.User("@me")
	return bot.Username
}

func (notifier *DiscordTradeNotifier) Notify(message string) error {
	_, err := notifier.session.ChannelMessageSend(notifier.channelID, message)
	if err != nil {
		return err
	}

	return nil
}
