FROM node:14

WORKDIR /app


COPY package* /app/

RUN npm ci

COPY * /app/

EXPOSE 2000

CMD npm start
