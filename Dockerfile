# Use uma imagem do Go para desenvolvimento e execução
FROM golang:1.23.2-alpine

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie o código-fonte para o container
COPY . .

# Baixe as dependências
RUN go mod download

# Use `go run` para executar o projeto
ENTRYPOINT ["go", "run", "./cmd"]
EXPOSE 8080