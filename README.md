# GPT-TeleChat-Bot: A Telegram Bot with ChatGPT and Golang\n\nThis is a Telegram bot integrating ChatGPT and Golang. It leverages OpenAI's GPT-3 language model for real-time responses to user messages.\n\n## Prerequisites\nFor building and running this application, ensure Go is installed on your system.\n\n## Features\n- Real-time, human-like responses to user messages with ChatGPT API\n- User messages persisted with SQLite\n- Telegram support\n- Built using Go for optimized performance\n\nPrior to using the bot, create a Telegram bot using the [BotFather framework](https://t.me/botfather). After obtaining an API token, also get an [API key from OpenAI](https://platform.openai.com/account/api-keys)\n\nRefer the example `.env` file below:\n\n```.env\nTELEGRAM_API_KEY=""\nOPENAI_TOKEN=""\nRETAIN_HISTORY="false"\n```\nWhen `RETAIN_HISTORY="true"`, the bot submits previous conversations with the current text (see details [here](https://platform.openai.com/docs/guides/chat/introduction)). If set to false, only the prompt + the current user's text is sent, reducing the tokens sent per request.\n\nCreate a `prompt.txt` or rename the example file:\n\n```sh\n$ mv prompt.example.txt prompt.txt\n```\nThe prompt customizes the bot's reaction to messages.\n\n## Example chat conversation\n![Example Image](./screenshots/scrnsht1.png)\n\n## Installation\nCloning this repository to your local machine:\n\n```sh\ngit clone github.com/UE5d/GPT-TeleChat-Bot.git\n```\nThen, navigate to the project directory:\n\n```sh\ncd GPT-TeleChat-Bot\n```\nFinally, compile the project:\n\n```sh\ngo build -o /opt/GPT-TeleChat-Bot\n```\n\n## Contributing\nContributions are welcome! Follow these steps:\n\n- Fork this repository.\n- Create a new branch (git checkout -b <branch-name>).\n- Commit your changes (git commit -am "<