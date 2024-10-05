FROM ubuntu:latest

# Update and install ca-certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

CMD /app/slack-bot-server -channel "$SLACK_CHANNEL" -token "$SLACK_TOKEN"