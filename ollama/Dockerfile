FROM ollama/ollama:latest

COPY ./Modelfile /Modelfile
COPY ./start.sh /start.sh

RUN apt-get update && apt-get install -y curl

RUN chmod +x /start.sh

EXPOSE 11434

ENTRYPOINT ["/start.sh"]
