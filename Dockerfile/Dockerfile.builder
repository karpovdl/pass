FROM karpovdl/golang:1.14.4-alpine3.12

ARG app_df
ARG app_version

LABEL author="Denis Karpov" \
      site="github.com/karpovdl" \
      email="karpovdl@hotmail.com" \
      version=${app_version} \
      release-date="2020-06-18" \
      alpine="3.12" \
      golang="1.14.4"

ENV TZ="Europe/Moscow" \
### APP SRC
    APP_NAME="github.com/karpovdl/pass" \
    APP_NAME_BRANCH="master" \
    APP_VERSION=${app_version}

COPY ${app_df} $GOPATH/bin/Dockerfile.run

CMD export WORK_DIR=$GOPATH/src \
### FIRST INIT
 && rm -rf /var/lib/apt/lists/* \
 && update-ca-certificates \
### GET APP
 && go get -d $APP_NAME \
 && cd $GOPATH/src/$APP_NAME \
 && git checkout $APP_NAME_BRANCH > null \
 && cd $GOPATH/src \
### BUILD APP
 && go install -ldflags '-s -w' -v $APP_NAME \
 && cd $GOPATH/bin && tar -C ./ -cf - .
