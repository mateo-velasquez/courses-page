FROM node:20-alpine
WORKDIR /app
COPY frontend-lms/package*.json ./
RUN npm install
COPY frontend-lms/ ./
EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
