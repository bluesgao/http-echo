echo -e  "****开始执行自动化部署****"
APP_NAME="http-echo"

echo -e "---step0:应用名称${APP_NAME}"
echo -e "---step1:更新代码"
git pull origin main

echo -e "---step2:编译"
go build

echo -e "---step3:更改权限"
chmod -R 777 ${APP_NAME}

echo -e "---step4:杀掉进程"
pid1=`ps -ef|grep -E ${APP_NAME}|grep -v grep|awk '{print $2}'`
kill -9 $pid1

echo -e "---step5:运行"
nohup ./${APP_NAME} >/dev/null 2>&1 &

pid2=`ps -ef|grep -E ${APP_NAME}|grep -v grep|awk '{print $2}'`

echo -e "****部署成功,进程ID为:$pid2****"