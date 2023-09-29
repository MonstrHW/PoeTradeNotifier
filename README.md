# Path of Exile Trade Notifier

A Command-line tool that listening log file of a game named as "Path of Exile" for recognizing buy messages and sending them to your Telegram/Discord PM through Bot.

## Command-line arguments

- n - notifier type (string)
- t - tg/discord bot token (string)
- p - path to Client.txt (string)
- c - tg chat id (int)
- u - discord user id (string)
- a - send notifications only when AFK
- d - notify about abnormal disconnect
- s - start tool only for send current tg chat id
- v - version
- h - help

## Install

Download archive from [Releases](https://github.com/MonstrHW/PoeTradeNotifier/releases) page and extract it in any place.

## Telegram

1. Contact the [BotFather](https://telegram.me/BotFather) for creating Bot.
2. Type /newbot and follow the instructions.
3. After Bot created, edit start script using example from comment. Use one with telegram type as base.
4. Put token to -t key and add -s key at end of script line, then start tool using script.
5. Find your created Bot in Telegram and type /start in chat with him, Bot should send to you current chat id.
6. Now remove -s key and put chat id to -c key, also edit options you need like path to Client.txt.

## Discord

1. Enable developer mode in discord settings: Settings > Advanced.
2. Go to [Discord API](https://discord.com/developers/applications).
3. Create App: Applications > New Application (Bot token located here).
4. Then you should invite Bot to your server: Your App > OAuth2 > URL Generator > Click "bot" option > Copy Link > Past in browser.
5. Get your Discord user id: Settings > My Account > "dots" after your name > Copy User ID.
6. Edit start script using example from comment. Use one with discord type as base.
7. Put token to -t key and user id to -u key, also edit options you need like path to Client.txt.

> You should be in one server with Bot.

## Start

Start tool for any time using script.
