FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/app
RUN go get github.com/dgrijalva/jwt-go \
	github.com/gorilla/mux \
	github.com/dgrijalva/jwt-go/request \
	github.com/jinzhu/inflection \
	github.com/go-sql-driver/mysql \
	github.com/jinzhu/gorm \
	github.com/jinzhu/gorm/dialects/mysql \
	github.com/joho/godotenv \
	github.com/go-playground/locales/currency \
	github.com/go-playground/locales \
	github.com/go-playground/universal-translator \
	gopkg.in/go-playground/validator.v9

COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]
