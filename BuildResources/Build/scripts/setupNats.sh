# @Author: Ximidar
# @Date:   2019-02-15 19:29:33
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-17 15:26:59

AMD64=https://github.com/nats-io/gnatsd/releases/download/v1.4.1/gnatsd-v1.4.1-linux-amd64.zip
ARM=https://github.com/nats-io/gnatsd/releases/download/v1.4.1/gnatsd-v1.4.1-linux-arm6.zip
ARM64=https://github.com/nats-io/gnatsd/releases/download/v1.4.1/gnatsd-v1.4.1-linux-arm64.zip

DL=$HOME/downloads/
mkdir -p $DL

# wget all files
wget $AMD64 -O $DL/AMD64.zip
wget $ARM -O $DL/ARM.zip
wget $ARM64 -O $DL/ARM64.zip

unzip $DL/AMD64.zip -d /usr/local/ && mv /usr/local/*amd64/ /usr/local/natsAMD64
unzip $DL/ARM.zip -d /usr/local/ && mv /usr/local/*arm6/ /usr/local/natsARM6
unzip $DL/ARM64.zip -d /usr/local/ && mv /usr/local/*arm64/ /usr/local/natsARM64