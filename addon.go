package latinabot

import (
	"fmt"
	"os"
	"strings"

	"github.com/LalatinaHub/LatinaApi/common/member"
	"github.com/NicoNex/echotron/v3"
)

var (
	bannerUrl = os.Getenv("BANNER_URL")
)

func (b *bot) menu(update *echotron.Update) {
	var (
		expired, password = member.GetMember(update.ChatID())
		message           = []string{}
	)

	if update.Message != nil {
		message = append(message, fmt.Sprintf("Halo %s !", update.Message.From.FirstName))
	} else if update.CallbackQuery != nil {
		message = append(message, fmt.Sprintf("Halo %s !", update.CallbackQuery.Message.From.FirstName))
	}

	message = append(message, fmt.Sprintf("\nID: <code>%d</code>", update.ChatID()))
	if expired <= 0 {
//		premiumData := member.GetPremiumAccount(password)
		message = append(message, "Status Akun: <b>Premium</b> 👑")
		message = append(message, fmt.Sprintf("Password: <code>%s</code>", password))
		message = append(message, fmt.Sprintf("Masa Aktif: %d Day(s)", expired))
//		message = append(message, fmt.Sprintf("Quota: %s", "Unlimited"))
/*
		message = append(message, "\nInfo Akun VPN:")
		message = append(message, fmt.Sprintf("Jenis: %s", premiumData.VPN))
		message = append(message, fmt.Sprintf("Quota: %d MB", premiumData.Quota))
		message = append(message, fmt.Sprintf("Password: <code>%s</code>", premiumData.Password))
		message = append(message, fmt.Sprintf("Domain: %s", premiumData.Domain))
		message = append(message, fmt.Sprintf("Path: /%s", premiumData.VPN))
*/
		message = append(message, "\nBatasan:")
		message = append(message, "- Tidak bisa mengambil akun VPN lebih dari 20")

		message = append(message, "\nCatatan:")
		message = append(message, "- Masa aktif akun tidak berlaku akumulasi")
		message = append(message, "- Ketahuan nakal = Premium hangus")
		message = append(message, "- Kirim /resetpass untuk reset password akun vpn")
	} else {
		message = append(message, "Status Akun: <b>Gratis</b> 👒")
		message = append(message, fmt.Sprintf("Password: <code>%s</code>", password))
		message = append(message, "Masa Aktif: Lifetime")

		message = append(message, "\nBatasan:")
		message = append(message, "- Tidak bisa mengambil akun VPN lebih dari 5")
		message = append(message, "- Hanya bisa mengambil akun VPN dari SG, ID dan US")
		message = append(message, "- Hanya bisa mengambil akun Trojan, VMess dan Vless")
/*
		message = append(message, "\nCara Order:")
		message = append(message, "1. Lakukan pembayaran pada e-wallet (lihat banner)")
		message = append(message, "2. Simpan bukti pembayaran berupa screenshot")
		message = append(message, "3. Kirimkan bukti pembayaran pada bot ini")
		message = append(message, "4. Duduk manis sambil menunggu aktivasi akun diproses")
*/
		message = append(message, "\nCatatan:")
	}

	message = append(message, "- Kirim /newpass untuk memperbarui password api")
	message = append(message, "- Segera ganti password apabila bocor ke publik")
//	message = append(message, "- Bisa full refund selama quota diatas 90%, dibawah itu menyesuaikan sisa quota")

//	message = append(message, "\nAmbil akun VPN gratis full speed dengan langkah mudah !")
	message = append(message, "\nOwner  : @JeelsBoobz")
	message = append(message, "Credits : @FoolVPN")

	if bannerUrl == "" {
		bannerUrl = "https://raw.githubusercontent.com/JeelsBoobz/LatinaBot/main/assets/Banner.png"
	}

	go b.SendPhoto(echotron.NewInputFileURL(bannerUrl), update.ChatID(), &echotron.PhotoOptions{
		Caption:   strings.Join(message, "\n"),
		ParseMode: "HTML",
		ReplyMarkup: echotron.InlineKeyboardMarkup{
			InlineKeyboard: [][]echotron.InlineKeyboardButton{
				{
					{
						Text: "😎 Ambil Akun",
						URL:  "https://darling-macaque-heroic.ngrok-free.app/get?format=raw&region=Asia&cdn=speed.cloudflare.com&sni=google.com&pass=" + password,
					},
					{
						Text: "🛒 Order Premium",
						URL: "https://t.me/casperinject",
					},
					/*{
						Text:         "Buat Akun (Premium) ✨",
						CallbackData: "create_account",
					},*/
				},
				{
					{
						Text: "📄 Dokumentasi",
						URL:  "https://darling-macaque-heroic.ngrok-free.app/",
					},
					{
						Text: "💬 Gabung Grup",
						URL:  "https://t.me/MiTM_Attack",
					},
				},
				{
					{
						Text: "💰 Sedekah",
						URL: "https://t.me/MiTM_Attack/31845/31964",
					},
				},
			},
		},
	})
}
