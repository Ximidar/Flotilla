# @Author: Ximidar
# @Date:   2019-02-15 19:29:33
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-27 17:57:49

move(){
	echo "Moving From $1 to $2"
	if cd $2 ; then
		mv -v $1 $2
	else
		echo "Target $2 doesn't exist. Making it"
		mkdir -p $2
		mv -v $1 $2
	fi

}

AMD64=https://github.com/nats-io/gnatsd/releases/download/v1.4.1/gnatsd-v1.4.1-linux-amd64.zip
ARM=https://github.com/nats-io/gnatsd/releases/download/v1.4.1/gnatsd-v1.4.1-linux-arm6.zip
ARM64=https://github.com/nats-io/gnatsd/releases/download/v1.4.1/gnatsd-v1.4.1-linux-arm64.zip

DL=$HOME/downloads
TEMP_LOC=$HOME/tmp
mkdir -p $DL
mkdir -p $TEMP_LOC
mkdir -p $NATS_LOC/AMD64/
mkdir -p $NATS_LOC/ARM6/
mkdir -p $NATS_LOC/ARM64/

# wget all files
wget $AMD64 -O $DL/AMD64.zip
wget $ARM -O $DL/ARM.zip
wget $ARM64 -O $DL/ARM64.zip


unzip $DL/AMD64.zip -d $TEMP_LOC && move $TEMP_LOC/*amd64/gnatsd $NATS_LOC/AMD64/
unzip $DL/ARM.zip -d $TEMP_LOC && move $TEMP_LOC/*arm6/gnatsd $NATS_LOC/ARM6/
unzip $DL/ARM64.zip -d $TEMP_LOC && move $TEMP_LOC/*arm64/gnatsd $NATS_LOC/ARM64/


# Cleanup
rm -rf $TEMP_LOC
rm -rf $DL

# Nats servers are now located at the NATS_LOC