# work
FROM golang:1.14
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program

## Our start command which kicks off
## our newly created binary executable
RUN go build -o ApiService .

CMD ["/app/ApiService"]

#WORK 
# FROM golang:1.14
# WORKDIR /app
# ADD . /app
# #COPY . ./app
# RUN go build
# EXPOSE 80
# CMD ["/app/hello"]
