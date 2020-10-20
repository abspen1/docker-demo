FROM node:12.18.3

WORKDIR /app

COPY package.json package.json
COPY package-lock.json package-lock.json
COPY app.js app.js
COPY .env .env

RUN npm install

CMD [ "node", "app.js" ]