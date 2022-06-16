.PHONY: all build run clean help go_build

# 定义变量
PBDIR=".:./pbfiles"
PBOUTPUTDIR="./pbservices/"
PBFILE?="hello"

outpb:
	protoc --proto_path=${PBDIR} --go_out=plugins=grpc:${PBOUTPUTDIR} ${PBFILE}.proto

help:
	@echo "make help - 帮助"
	@echo "make outpb - 生成pb代码,PBFILE=user make outpb"
