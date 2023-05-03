## Developer Documentation

### Overview

This program is a simple Discord bot that utilizes the OpenAI GPT-3 API to generate responses to user messages containing questions. The bot connects to a Discord server and listens for messages in the channels it has access to. When a question is detected, it sends the message to the GPT-3 API to generate a response and sends the response back to the Discord channel.

### Prerequisites

1. [Go 1.17](https://golang.org/dl/) installed.
2. [Docker](https://www.docker.com/get-started) installed (optional, for containerized deployment).
3. A Discord bot token.
4. An OpenAI API key.

### Project Structure

- `main.go`: The main program file that contains the implementation of the Discord bot and GPT-3 API integration.

### Setting Up The Environment

1. Clone the repository.
2. Run `go mod init <module-name>` to initialize the Go module.
3. Run `go mod tidy` to download the required dependencies.

### Running The Program

1. Set the following environment variables:
   - `DISCORD_BOT_TOKEN`: Your Discord bot token.
   - `OPENAI_API_KEY`: Your OpenAI API key.

2. Run the program using `go run main.go`.

### Building The Docker Container

1. Build the Docker image using `docker build -t my-discord-bot .`.
2. Run the Docker container using `docker run -it --rm -e DISCORD_BOT_TOKEN=<your_bot_token> -e OPENAI_API_KEY=<your_openai_api_key> my-discord-bot`.

## User Documentation

### Overview

This Discord bot can generate responses to user messages containing questions. When you ask a question in a channel that the bot has access to, it will use the OpenAI GPT-3 API to generate a response and send it back to the channel.

### Usage

1. Invite the bot to your Discord server.
2. Ensure the bot has the necessary permissions to read and send messages in the desired channels.
3. Ask questions in the channels by typing your question and sending the message. The bot will automatically detect questions and respond accordingly.

### Notes

- The bot may not always generate relevant responses to every question. It relies on the GPT-3 API's output, which may occasionally produce irrelevant or default answers.
- The bot is designed to ignore non-question messages and messages sent by itself.
