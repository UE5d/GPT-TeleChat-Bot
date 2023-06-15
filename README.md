# GPT-TeleChat-Bot: A Telegram Bot with ChatGPT and Golang\n\nThis is a Telegram bot integrating ChatGPT and Golang. It leverages OpenAI's GPT-3 language model for real-time responses to user messages.\n\n## Prerequisites\nFor building and running this application, ensure Go is installed on your system.\n\n## Features\n- Real-time, human-like responses to user messages with ChatGPT API\n- User messages persisted with SQLite\n- Telegram support\n- Built using Go for optimized performance\n\nPrior to using the bot, create a Telegram bot using the [BotFather framework](https://t.me/botfather). After obtaining an API token, also get an [API key from OpenAI](https://platform.openai.com/account/api-keys)\n\nRefer the example `.env` file below:\n\n```.env\nTELEGRAM_API_KEY=""\nOPENAI_TOKEN=""\nRETAIN_HISTORY="false"\n```\nWhen `RETAIN_HISTORY="true"`, the bot submits previous conversations with the current text (see details [here](https://platform.openai.com/docs/guides/chat/introduction)). If set to false, only the prompt + the current user's text is sent, reducing the tokens sent per request.\n\nCreate a `prompt.txt` or rename the example file:\n\n```sh\n$ m