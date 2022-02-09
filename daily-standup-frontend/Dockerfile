FROM node:alpine

LABEL Jubril Aminu  <jubrilaminu304@gmail.com>

WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm i
CMD ["npm", "run", "start"]
