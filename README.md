# Hyperledger Fabric Setup Guide
## Prerequisites

1. Docker installed on your system
2. Docker Compose installed on your system
3. Go (Golang) installed on your system

### Step 1: Install Docker
```
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
$(. /etc/os-release && echo $VERSION_CODENAME) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo groupadd docker
sudo usermod -aG docker ${USER}
sudo chmod 666 /var/run/docker.sock
```
### Step 2: Install Docker Compose
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
cd ..
```
### Step 3: Download Hyperledger Fabric Samples
```
curl -sSL https://bit.ly/2ysbOFE | bash -s
mv fabric-samples/ fabric
```
### Step 4: Install Go (Golang)
```
wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz
```
### Step 5: Set up Go Environment
Edit the .bashrc or .profile file using your preferred text editor (e.g., nano, vi, vim) and add the following line at the end:
```
export PATH=$PATH:/usr/local/go/bin
```
then reload .bashrc file
```
source ~/.bashrc
```
check go version 
```
go version
```
### Step 6: Start Hyperledger Fabric Test Network
```
cd fabric/
cd test-network
./network.sh up createChannel -c mychannel -ca
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go
```
*** I believed that I should change this folder "../asset-transfer-basic/chaincode-go/" with my own project folder.
