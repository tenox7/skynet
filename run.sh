docker rm -f skynet
docker run -e OPENAI_MODEL_NAME=gpt-4o -e OPENAI_API_KEY=YOURKEY_HERE -p 4000:4000 --rm -d --name skynet  tenox7/skynet
