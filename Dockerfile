FROM golang:1.12.4

WORKDIR $GOPATH/src/CookingBoy

COPY . .

RUN apt-get update && apt-get install -y --no-install-recommends python3.4

RUN apt-get update && apt-get install -y --no-install-recommends python3-pip

RUN apt-get update && apt-get install -y --no-install-recommends python-setuptools

RUN apt-get update && apt-get install -y --no-install-recommends ffmpeg

RUN apt-get update && apt-get install -y --no-install-recommends curl

RUN pip3 install imgurpython

RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl && chmod a+rx /usr/local/bin/youtube-dl

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["CookingBoy"]