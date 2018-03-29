FROM node:9.9

WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install

COPY . .

CMD ["node", "member.js"]
