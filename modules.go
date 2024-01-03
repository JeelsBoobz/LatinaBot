package latinabot

import (
	"os"

	A "github.com/LalatinaHub/LatinaApi/common/account"
	"github.com/LalatinaHub/LatinaBot/helper"
	"github.com/NicoNex/echotron/v3"
)

func Client() *bot {
	return &bot{
		API: echotron.NewAPI(botToken),
	}
}

func SendVPNToTopic(chatID int64, topicID int) {
	var (
		bug      = []string{"BUG.COM"}
		accounts = A.PopulateBugs(A.Get("WHERE VPN != 'shadowsocks' AND org NOT LIKE '%Alibaba%' ORDER BY RANDOM() LIMIT 1"), bug, bug)
	)

	if len(accounts) > 0 {
		message := helper.MakeVPNMessage(accounts[0])

		go Client().SendMessage(message, chatID, &echotron.MessageOptions{
			ParseMode:           "HTML",
			ReplyToMessageID:    int(topicID),
			DisableNotification: true,
		})
	}
}

func SendTextAsDocument(text, filename string, chatID int64) {
	filename = filename + ".txt"
	f, _ := os.Create(filename)
	defer os.Remove(filename)
	defer f.Close()

	f.WriteString(text)

	Client().SendDocument(echotron.NewInputFilePath(filename), chatID, nil)
}

func DocumentToAdmin(text, filename string) {
	go SendTextAsDocument(text, filename, adminID)
}
