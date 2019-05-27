# Blimpy Bot

A telegram (https://telegram.org) bot service for Zeit Now (https://zeit.co/now).

# Deployment

On Mac:
```
$ git clone https://github.com/voutasaurus/blimpybot
$ cd blimpybot
$ export blimpysuffix=$(base64 < /dev/urandom | head -c 10)
$ sed -i '' -e "s/blimpybot/blimpybot-$blimpysuffix/g" now.json
$ npm install -g now
$ now # this will prompt you to login/create an Zeit Now account
$ export $BOT_TOKEN=<Get this from Telegram chat with @botfather>
$ now secret add bot-token $BOT_TOKEN
$ now --target production
$ curl -F "url=https://blimpybot-$blimpysuffix.now.sh" https://api.telegram.org/bot$BOT_TOKEN/setWebhook
```

You should be able to chat with your bot now.

![Woop-Woop-Woop](https://raw.githubusercontent.com/voutasaurus/blimpybot/master/screenshot.png)

# Modifications

This app uses environment variables defined in `now.json` prefixed with
"BLIMPY_" as a key value map to define responses of the bot.

You can change the env section in now.json to add triggers and preprogrammed
responses.

# TODO

Right now all triggers are exact matches (case insensitive). The program could
be modified to match on substrings or filter out punctuation.

