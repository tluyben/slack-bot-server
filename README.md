# 🤖 Slack Bot Server in Go

This project implements a simple Slack bot server in Go that can send and receive messages, as well as forward messages to a webhook.

## 📋 Features

- 📡 Listen for incoming Slack events
- 💬 Send messages to a Slack channel
- 🔄 Forward Slack messages to a webhook (optional)
- 🚀 Easy to set up and run

## 🛠️ Prerequisites

- Go 1.16 or higher
- Slack Bot Token
- Slack Channel ID
- (Optional) Webhook URL

## 🚀 Getting Started

1. Clone this repository:

   ```
   git clone https://github.com/yourusername/slack-bot-server.git
   cd slack-bot-server
   ```

2. Install dependencies:

   ```
   make deps
   ```

3. Build the bot:

   ```
   make build
   ```

4. Run the bot:

   ```
   ./slack-bot-server -token=xoxb-your-token -channel=C0123456789
   ```

   To include a webhook:

   ```
   ./slack-bot-server -token=xoxb-your-token -channel=C0123456789 -webhook=https://your-webhook-url
   ```

## 🔧 Configuration

The bot accepts the following command-line flags:

- `-token`: Slack Bot Token (required)
- `-channel`: Slack Channel ID (required)
- `-port`: Port to listen on (default: 8080)
- `-webhook`: Webhook URL (optional)

## 📝 Usage

To send a message to Slack:

```
curl -X POST -d "Hello from the bot!" http://localhost:8080/send
```

## 🛠️ Makefile Commands

- `make deps`: Install dependencies
- `make build`: Build the bot
- `make run`: Run the bot (requires setting TOKEN and CHANNEL environment variables)
- `make clean`: Remove built binary

## 🔒 Security Note

Never commit your Slack token or sensitive URLs to version control. Use environment variables or a secure configuration management system in a production environment.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
