#!/bin/bash

# if [[ "" ==  "$1" ]]; then
#  echo "无运行类型参数，exit"
#  exit
# fi

#查询pid并kill
PID=`lsof -i:46200 | grep main | awk '{print $2}'`
if [[ "" !=  "$PID" ]]; then
  echo "killing $PID"
  kill -9 $PID
fi

#当前文件目录
CDIR="`pwd`"/"`dirname $0`"
echo $CDIR

#判断是否需要重新编译
if [[ "build" ==  "$1" ]]; then
  #删除main文件
  if [ -e $CDIR/main ]; then
    rm $CDIR/main
  fi

  #重新编译main文件
  go build $CDIR/main.go
fi

#运行main
if [ -e $CDIR/main ]; then
     nohup $CDIR/main > main.log 2 >&1 &
     tail -f main.log
else
     echo "main文件不存在,exit"
fi