FROM golang:1.21.6
WORKDIR /app 
COPY . .

# DOWNLOAD AND INSTALL THE DEP?!
RUN go get -d -v ./...

#  BUILD THE GO APP
RUN go build -o api .

EXPOSE 8000

CMD ["./api"]