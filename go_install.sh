echo "===========download===="
#wget https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz


echo "==========tar======"
#tar -zxf go1.17.7.linux-amd64.tar.gz -C /usr/local

echo "==========profile======"
#golang env config
echo "
##go env
export GO111MODULE=on
export GOROOT=/usr/local/go
export PATH=$PATH:/usr/local/go/bin
export GOPROXY=https://goproxy.cn,direct

" >> /etc/bashrc


source /etc/bashrc

echo "==========verify========="
go version
