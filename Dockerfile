FROM golang:alpine

ENV NKF_VERSION 2.1.4
ENV APP_PATH /go/src/gauss/go-mecab
ENV WORK_DIR /work
ENV MECAB_DIR /work/mecab/mecab
ENV IPADIC_DIR /work/mecab/mecab-ipadic
ENV NEOLOGD_DIR /work/mecab-ipadic-neologd
ENV CGO_LDFLAGS "-L/usr/local/lib -lmecab -lstdc++"
ENV CGO_CFLAGS "-I/usr/local/include"
ENV BUILD_DEPS 'build-base curl xz git'

WORKDIR ${WORK_DIR}
COPY . ${APP_PATH}

# Install dependencies
RUN apk add --update --no-cache --virtual .build-deps ${BUILD_DEPS}

# Install nkf
RUN wget http://jaist.dl.sourceforge.jp/nkf/64158/nkf-${NKF_VERSION}.tar.gz	\
 && tar zxf nkf-${NKF_VERSION}.tar.gz
WORKDIR ${WORK_DIR}/nkf-${NKF_VERSION}
RUN make \
 && make install

# Install MeCab
WORKDIR ${WORK_DIR}
RUN git clone https://github.com/taku910/mecab.git
WORKDIR ${MECAB_DIR}
RUN ./configure --enable-utf8-only --with-charset=utf8 \
 && make \
 && make install

# Install IPA dic
WORKDIR ${IPADIC_DIR}
RUN ./configure --with-charset=utf8 \
 && make install

# Install Neologd
WORKDIR ${WORK_DIR}
RUN git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git

# rebuild dictionary
WORKDIR ${IPADIC_DIR}
RUN nkf --overwrite -Ew ${IPADIC_DIR}/*.csv\
 && mv ${APP_PATH}/seed/*.csv ${IPADIC_DIR} \
 && xz -dkv ${NEOLOGD_DIR}/seed/mecab-user-dict-seed.*.csv.xz \
 && mv ${NEOLOGD_DIR}/seed/mecab-user-dict-seed.*.csv ${IPADIC_DIR} \
 && /usr/local/libexec/mecab/mecab-dict-index -f utf-8 -t utf-8 \
 && make install

# Install go app dependencies
WORKDIR ${APP_PATH}
RUN go get -u github.com/golang/dep/cmd/dep \
 && go get -u gopkg.in/godo.v2/cmd/godo \
 && go get -u github.com/swaggo/swag/cmd/swag \
 && dep ensure \
 && swag init

# Clean up
RUN rm -rf \
       ${WORK_DIR} \
       ~/.cache

CMD ["go", "run", "main.go"]
# CMD ["godo", "server", "--watch"]
