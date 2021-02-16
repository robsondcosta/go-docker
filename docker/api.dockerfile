# Inicia com a imagem do golang
FROM golang:alpine as builder 

# Instalação do git, o git é necessário para buscar os arquivos de dependências.
RUN apk update && apk add --no-cache git

# Informa o diretório que será utilizado no container
WORKDIR /app

# Copia os arquivos go.mod e go.sum 
COPY go.mod go.sum ./

# Baixa todas as dependências. As dependências serão armazenadas em cache se os arquivos go.mod e go.sum não forem alterados
RUN go mod download 

# Copia os arquivos do diretório atual para o diretório especificado do container
COPY . .

# Copia o aplicativo do go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Inicia um novo estágio
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia o arquivo binário pré-construído do estágio anterior. Observe que também copiamos o arquivo .env
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 

# Precisa tanto do dockerize como do wait para que a aplicação espere a execução do banco, excencial para utilização de migrations
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

# A porta de ecesso a aplicação
EXPOSE 3000

# Comando para execução da aplicação 😢
CMD ["./main"]