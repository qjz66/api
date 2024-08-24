#!/bin/bash

# 参数：主机名、端口、命令
HOST="$1"
PORT="$2"
shift 2
CMD="$@"

# 检查参数
if [ -z "$HOST" ] || [ -z "$PORT" ] || [ -z "$CMD" ]; then
  echo "Usage: $0 host port -- command"
  exit 1
fi

# 等待服务准备好
echo "Waiting for $HOST:$PORT to be available..."
while ! nc -z $HOST $PORT; do
  sleep 1
done

echo "$HOST:$PORT is up - executing command"
exec $CMD
