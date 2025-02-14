# Usa un'immagine di Node.js come base
FROM node:18

# Imposta la directory di lavoro
WORKDIR /app

# Copia package.json e installa le dipendenze
COPY package.json package-lock.json ./
RUN npm install

# Copia tutto il codice dell'app
COPY . .

# Esponi la porta 3000 per le API
EXPOSE 3000

# Avvia l'applicazione
CMD ["node", "server.js"]

