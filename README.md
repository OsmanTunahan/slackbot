# SlackBot

This repository contains a Slack bot written in Golang. The Slack bot uses socketmode to handle slash commands for sending emails, SMS, and generating presigned URLs. It returns a Slack attachment to the channel including the status code and body, which contains either a success message, error message, or presigned URL.

## Features

- **Email Sending**: Send emails using the `/emails` slash command.
- **SMS Sending**: Send SMS messages using the `/sms` slash command.
- **Presigned URLs**: Generate presigned URLs for file uploads/downloads using the `/url` slash command.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Development](#development)
- [License](#license)

## Requirements

### Slack Bot

* You need to create a Slack app and enable SocketMode. You can learn more at [Slack API SocketMode](https://api.slack.com/apis/connections/socket).
* I've attached a `manifest.yml` file if you'd like to set up the app with manifests. The manifest file contains the scopes of the bot you will be adding to allow the bot to read and write to the channel as well as accept slash commands.
* You need to add your `AUTH_TOKEN`, `APP_TOKEN`, `CHANNEL_ID`, and `URL` to the `.env` file to connect to the Slack API as well as your channel that you will be adding the Slack bot to.

### Setting Up Slack App (If not using manifest.yml)

- Activate Socket Mode
- Activate Interactivity & Shortcuts
- Activate Incoming Webhooks
- Add Slash Commands
    - `/emails`
        - usage hint: `[Your email] [Receiver's email] [Subject] [Message]`
    - `/sms`
        - usage hint: `[Twilio number] [Receiver's number] [Message]`
    - `/url`
        - usage hint: `[Filename] [Request] [Duration]`
- Bot Token Scopes
    - `app_mentions:read` : view messages that directly mention your app
    - `chat:write` : send messages via your app
    - `commands` : using slash commands
    - `incoming-webhook` : post attachments & messages to channels

## Installation

To get started, clone the repository and build the project.

```sh
git clone https://github.com/OsmanTunahan/slackbot.git
cd slackbot
go build
```

## Configuration

Set the following environment variables to configure the application:

- `AUTH_TOKEN`: Slack bot user OAuth token.
- `APP_TOKEN`: Slack app-level token.
- `CHANNEL_ID`: ID of the Slack channel where responses will be sent.
- `URL`: Base URL for your backend services.

You can set these variables in your shell or use a `.env` file with a tool like [godotenv](https://github.com/joho/godotenv).

```sh
export AUTH_TOKEN="auth-token"
export APP_TOKEN="app-token"
export CHANNEL_ID="channel-id"
export URL="https://blabla.com/endpoint"
```

## Usage

Run the application:

```sh
go build
./slackbot
```

### Slash Commands

- `/emails from@example.com to@example.com subject message`: Send an email using the SendGrid API.
- `/sms from_number to_number message`: Send an SMS using the Twilio API.
- `/url filename GET|PUT duration`: Generate a presigned URL for file upload/download.

### Using the Slash Commands in Slack

#### To send an SMS using the Twilio API function:

```sh
/sms [Twilio number] [Receiver's number] [Message]
```

#### To send an email using the SendGrid API function:

```sh
/emails [Your email] [Receiver's email] [Subject] [Message]
```

#### To get a presigned URL for file upload/download:

```sh
/url [Filename] [Request] [Duration]
```

### To Upload or Download a File Using `curl` in Terminal:

```sh
curl -X PUT -d 'The contents of the file.' "{url}"
```

## Development

To contribute to this project, follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Make your changes.
4. Test thoroughly.
5. Submit a pull request.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=OsmanTunahan/slackbot&type=Date)](https://star-history.com/#OsmanTunahan/slackbot&Date)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.