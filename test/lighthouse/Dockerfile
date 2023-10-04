FROM cypress/included:13.2.0

WORKDIR /app
RUN npm install -g lighthouse-ci@1.13.1

COPY cmd.sh config.json ./
RUN chmod +x cmd.sh

ENTRYPOINT []
CMD [ "./cmd.sh" ]