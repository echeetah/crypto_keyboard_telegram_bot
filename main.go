package main

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const botToken string = "PASTE_YOUR_BOT_TOKEN_HERE"

const BITCOIN string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const ETHEREUM string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const BNBBEP20 string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const BNBBEP2 string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const POLKADOT string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const SOLANA string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const ATOM string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const TWTBEP20 string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const TWTBEP2 string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const TON string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const USDTTRC20 string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const USDTERC20 string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const NEAR string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const LITECOIN string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const DASH string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const ZEC string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const DOGE string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const RIPPLE string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const BCASH string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const ETC string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const BAT string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const UNI string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const DAI string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const CAKE string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const XMR string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const HBAR string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const TRON string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const KAVA string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const FIL string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const SXP string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const BAND string = "PASTE_YOUR_PUBLIC_KEY_HERE"
const ADA string = "PASTE_YOUR_PUBLIC_KEY_HERE"

var firstPageCryptoAssets = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔱 BTC"),
		tgbotapi.NewKeyboardButton("💠 ETH"),
		tgbotapi.NewKeyboardButton("🌆 BNB"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🕴 XMR"),
		tgbotapi.NewKeyboardButton("📘 DASH"),
		tgbotapi.NewKeyboardButton("🔗 ZEC"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🧿 DOT"),
		tgbotapi.NewKeyboardButton("🧬 SOL"),
		tgbotapi.NewKeyboardButton("⚛️ ATOM"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🛡 TWT"),
		tgbotapi.NewKeyboardButton("💎 TON"),
		tgbotapi.NewKeyboardButton("⏭"),
	),
)
var secondPageCryptoAssets = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💵 USDT"),
		tgbotapi.NewKeyboardButton("💳 DAI"),
		tgbotapi.NewKeyboardButton("🏵 TRON"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🐕 DOGE"),
		tgbotapi.NewKeyboardButton("🔘 LTC"),
		tgbotapi.NewKeyboardButton("🌚 XRP"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🦄 UNI"),
		tgbotapi.NewKeyboardButton("🥞 CAKE"),
		tgbotapi.NewKeyboardButton("♾ NEAR"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⏮"),
		tgbotapi.NewKeyboardButton("🧭 BAT"),
		tgbotapi.NewKeyboardButton("⏭"),
	),
)
var thirdPageCryptoAssets = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💰 BCH"),
		tgbotapi.NewKeyboardButton("⚙️ ETC"),
		tgbotapi.NewKeyboardButton("👨‍🎓 HBAR"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⌚️ BAND"),
		tgbotapi.NewKeyboardButton("🎚 ADA"),
		tgbotapi.NewKeyboardButton("📀 KAVA"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💾 FIL"),
		tgbotapi.NewKeyboardButton("🧸 SXP"),
		tgbotapi.NewKeyboardButton("empty"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⏮"),
		tgbotapi.NewKeyboardButton("empty"),
		tgbotapi.NewKeyboardButton("empty"),
	),
)

var chooseBinanceNetwork = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⏮"),
		tgbotapi.NewKeyboardButton("BEP20"),
		tgbotapi.NewKeyboardButton("BEP2"),
	),
)
var chooseUsdtNetwork = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⏮"),
		tgbotapi.NewKeyboardButton("ERC20"),
		tgbotapi.NewKeyboardButton("TRC20"),
	),
)

func main() {
	var pageNumber int = 0
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите криптовалюту и бот отправит в ответ номер кошелька:")

		switch update.Message.Text {
		case "/start":
			msg.ReplyMarkup = firstPageCryptoAssets
		case "⏭":
			if pageNumber == 0 {
				msg.ReplyMarkup = secondPageCryptoAssets
				pageNumber++
			} else if pageNumber == 1 {
				msg.ReplyMarkup = thirdPageCryptoAssets
				pageNumber++
			}
		case "⏮":
			if pageNumber == 1 {
				msg.ReplyMarkup = firstPageCryptoAssets
				pageNumber--
			} else if pageNumber == 2 {
				msg.ReplyMarkup = secondPageCryptoAssets
				pageNumber--
			} else {
				msg.ReplyMarkup = firstPageCryptoAssets
				pageNumber = 0
			}
		case "🌆 BNB":
			msg.Text = "Выберите сеть для перевода:"
			msg.ReplyMarkup = chooseBinanceNetwork
			pageNumber = 3
		case "🛡 TWT":
			msg.Text = "Выберите сеть для перевода:"
			msg.ReplyMarkup = chooseBinanceNetwork
			pageNumber = 4
		case "💵 USDT":
			msg.Text = "Выберите сеть для перевода:"
			msg.ReplyMarkup = chooseUsdtNetwork
			pageNumber = 5
		default:
			msg.Text = getWallet(update.Message.Text, pageNumber)
		}

		bot.Send(msg)
	}
}

func getWallet(asset string, number int) string {
	var publicKey string

	switch asset {
	case "🔱 BTC":
		publicKey = BITCOIN
	case "💠 ETH":
		publicKey = ETHEREUM
	case "BEP20":
		if number == 3 {
			publicKey = BNBBEP20
		} else if number == 4 {
			publicKey = TWTBEP20
		}
	case "BEP2":
		if number == 3 {
			publicKey = BNBBEP2
		} else if number == 4 {
			publicKey = TWTBEP2
		}
	case "🕴 XMR":
		publicKey = XMR
	case "📘 DASH":
		publicKey = DASH
	case "🔗 ZEC":
		publicKey = ZEC
	case "🧿 DOT":
		publicKey = POLKADOT
	case "🧬 SOL":
		publicKey = SOLANA
	case "⚛️ ATOM":
		publicKey = ATOM
	case "💎 TON":
		publicKey = TON
	case "ERC20":
		publicKey = USDTERC20
	case "TRC20":
		publicKey = USDTTRC20
	case "💳 DAI":
		publicKey = DAI
	case "🏵 TRON":
		publicKey = TRON
	case "🐕 DOGE":
		publicKey = DOGE
	case "🔘 LTC":
		publicKey = LITECOIN
	case "🌚 XRP":
		publicKey = RIPPLE
	case "🦄 UNI":
		publicKey = UNI
	case "🥞 CAKE":
		publicKey = CAKE
	case "♾ NEAR":
		publicKey = NEAR
	case "🧭 BAT":
		publicKey = BAT
	case "💰 BCH":
		publicKey = BCASH
	case "⚙️ ETC":
		publicKey = ETC
	case "👨‍🎓 HBAR":
		publicKey = HBAR
	case "⌚️ BAND":
		publicKey = BAND
	case "🎚 ADA":
		publicKey = ADA
	case "📀 KAVA":
		publicKey = KAVA
	case "💾 FIL":
		publicKey = FIL
	case "🧸 SXP":
		publicKey = SXP
	case "empty":
		publicKey = "Нажмите кнопку с названием криптовалюты или перейдите на другую страницу:"
	}

	return publicKey
}
