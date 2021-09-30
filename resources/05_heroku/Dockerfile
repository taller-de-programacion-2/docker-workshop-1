FROM node:14

WORKDIR /app
RUN apt-get update
RUN apt-get -y install postgresql-client


COPY package* /app/

RUN npm ci

COPY index.js /app/

CMD npm start
