set -x

mkdir -p $HOME/gopath-staging
cd $HOME/gopath-staging
git clone --recursive https://github.com/storj/storj-vendor.git .
./setup.sh
mkdir -p src/storj.io
mv $HOME/gopath/src/github.com/storj/storj src/storj.io
rm -rf $HOME/gopath
mv $HOME/gopath{-staging,}
export TRAVIS_BUILD_DIR=$HOME/gopath/src/storj.io/storj
cd $TRAVIS_BUILD_DIR

go install -v storj.io/storj/cmd/captplanet
#apt-get install -y python3.4 #sudo
#apt-get install --upgrade -y python-pip #sudo
#pip install --user virtualenv # sudo
virtualenv my_py3 --python=/usr/bin/python3.4
source my_py3/bin/activate
pip install --upgrade awscli

set +x
