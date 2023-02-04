FROM node:alpine
WORKDIR /client
COPY package.json .
COPY package-lock.json .
RUN npm i && npm i typescript -g
COPY . .
CMD ["npm", "run", "build"]