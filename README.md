# 🤖 Slack Bot Server in Go

This project implements a simple Slack bot server in Go that can send and receive messages, as well as forward messages to a webhook.

## 📋 Features

- 📡 Listen for incoming Slack events
- 💬 Send messages to a Slack channel
- 🔄 Forward Slack messages to a webhook (optional)
- 🚀 Easy to set up and run
- 🖥️ Cross-compilation support for Linux AMD64
- 🐳 Docker support

## 🛠️ Prerequisites

- Go 1.16 or higher
- A Slack workspace where you have permissions to add apps
- Basic familiarity with Slack
- Docker (optional)

## 🚀 Setup Steps

### Create a Slack App

1. Go to https://api.slack.com/apps
2. Click "Create New App" and choose "From scratch"
3. Name your app and select your workspace

### Set Up Bot User

1. In the left sidebar, click "Bot Users"
2. Click "Add a Bot User"
3. Choose a display name and username for your bot

### Configure Permissions

1. In the left sidebar, click "OAuth & Permissions"
2. Scroll to "Scopes" and add these Bot Token Scopes:
   - app_mentions:read
   - chat:write

### Enable Events

1. In the left sidebar, click "Event Subscriptions"
2. Toggle "Enable Events" to On
3. Under "Subscribe to bot events", add app_mention

### Install App to Workspace

1. In the left sidebar, click "Install App"
2. Click "Install App to Workspace"
3. Review and allow the permissions

### Get Bot Token

1. After installation, go back to "OAuth & Permissions"
2. Copy the "Bot User OAuth Token" (starts with xoxb-)

### Invite Bot to Channel

1. In Slack, go to the channel where you want to use the bot
2. Type /invite @your_bot_name

## 🚀 Getting Started

### Using Docker

1. Clone this repository:

   ```
   git clone https://github.com/tluyben/slack-bot-server.git
   cd slack-bot-server
   ```

2. Build and run using Docker Compose:

   ```
   docker-compose up --build
   ```

   Make sure to set the `SLACK_CHANNEL` and `SLACK_TOKEN` environment variables before running.

### Without Docker

1. Clone this repository:

   ```
   git clone https://github.com/tluyben/slack-bot-server.git
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

4. Cross-compile for Linux AMD64:

   ```
   make build-linux
   ```

5. Run the bot:

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
- `make build`: Build the bot for the current system
- `make build-linux`: Cross-compile the bot for Linux AMD64
- `make run`: Run the bot (requires setting TOKEN and CHANNEL environment variables)
- `make clean`: Remove built binaries

## 🔒 Security Note

Never commit your Slack token or sensitive URLs to version control. Use environment variables or a secure configuration management system in a production environment.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.