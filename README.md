# OpenAI Telnet Gateway

A Docker container that runs interactive OpenAI client accessible via telnet. Why? So you can use it from vintage operating systems.

## Requirements

- [OpenAI API Key](https://platform.openai.com/docs/quickstart)
- Docker

## How to run

```
docker run \
    -e OPENAI_MODEL_NAME=gpt-4o \
    -e OPENAI_API_KEY=... \
     -p 4000:4000 \
    --rm \
    -d \
    --name skynet \
    tenox7/skynet
```

Sample `run.sh` is provided. Insert your own API Key.

Telnet to port 4000 or whatever you set it to.

