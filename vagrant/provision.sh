#!/bin/bash

# Define variables.
BASH_PROFILE=/home/vagrant/.bash_profile
GO_HOME=/host/go
GO_FILE_NAME=go1.3beta1.linux-amd64.tar.gz
JAVA_FILE_NAME=jre-7u55-linux-x64.gz
JAVA_DIRECTORY_NAME=jre1.7.0_55
ELASTICSEARCH_DIRECTORY_NAME=elasticsearch-1.1.1
ELASTICSEARCH_FILE_NAME=${ELASTICSEARCH_DIRECTORY_NAME}.tar.gz

# Create .bash_profile
touch $BASH_PROFILE
chown vagrant:vagrant $BASH_PROFILE

apt-get update

# Install curl
apt-get install -y curl

# Install Git
apt-get install -y git

# Install Go 1.3beta1
curl -o /usr/local/$GO_FILE_NAME https://storage.googleapis.com/golang/$GO_FILE_NAME
tar -C /usr/local -xzf /usr/local/$GO_FILE_NAME
rm /usr/local/$GO_FILE_NAME
echo "export GOROOT=/usr/local/go" >> $BASH_PROFILE
echo "export GOPATH=$GO_HOME" >> $BASH_PROFILE
echo "export PATH=\$PATH:\$GOPATH/bin\$GOROOT:\$GOROOT/bin" >> $BASH_PROFILE
. $BASH_PROFILE

# Install Java 7 Update 55
curl -o /usr/local/lib/$JAVA_FILE_NAME https://s3-ap-northeast-1.amazonaws.com/yosssi/java/$JAVA_FILE_NAME
tar -C /usr/local/lib -xzf /usr/local/lib/$JAVA_FILE_NAME
rm /usr/local/lib/$JAVA_FILE_NAME
echo "export JAVA_HOME=/usr/local/lib/$JAVA_DIRECTORY_NAME" >> $BASH_PROFILE
echo "export PATH=\$PATH:\$JAVA_HOME/bin" >> $BASH_PROFILE
. $BASH_PROFILE

# Install Elasticsearch 1.1.1
curl -o /usr/local/lib/$ELASTICSEARCH_FILE_NAME https://download.elasticsearch.org/elasticsearch/elasticsearch/$ELASTICSEARCH_FILE_NAME
tar -C /usr/local/lib -xzf /usr/local/lib/$ELASTICSEARCH_FILE_NAME
rm /usr/local/lib/$ELASTICSEARCH_FILE_NAME
echo "export ELASTICSEARCH_HOME=/usr/local/lib/$ELASTICSEARCH_DIRECTORY_NAME" >> $BASH_PROFILE
echo "export PATH=\$PATH:\$ELASTICSEARCH_HOME/bin" >> $BASH_PROFILE
. $BASH_PROFILE
chmod 777 $ELASTICSEARCH_HOME

# Install Elasticsearch plugins
## Install Marvel
$ELASTICSEARCH_HOME/bin/plugin -i elasticsearch/marvel/latest
## Install elasticsearch-inquisitor
$ELASTICSEARCH_HOME/bin/plugin -i polyfractal/elasticsearch-inquisitor
## Install Japanese (kuromoji) Analysis
$ELASTICSEARCH_HOME/bin/plugin -i elasticsearch/elasticsearch-analysis-kuromoji/2.1.0
## Install elasticsearch-head
$ELASTICSEARCH_HOME/bin/plugin -i mobz/elasticsearch-head
