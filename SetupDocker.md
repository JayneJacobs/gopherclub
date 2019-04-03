

# Set up Docker on Cloud server

apt-get install dmsetup && dmsetup && dmsetup mknotes

### Add gpg key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

### Verify fingerprint of gpg key

```sh
apt-key fingerprint 0EBFCD88


root@gopherclub:~# apt-key fingerprint 0EBFCD88
pub   4096R/0EBFCD88 2017-02-22
      Key fingerprint = 9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid                  Docker Release (CE deb) <docker@docker.com>
sub   4096R/F273FCD8 2017-02-22

add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu xenial stable"
apt-get update
apt-get install docker-ce
```

//add user to docker group

```sh
usermod -aG docker gopherclub


root@gopherclub:~# docker --version
Docker version 18.09.4, build d14af54

docker run hello-world

curl -L https://github.com/docker/compose/releases/download/1.24.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
```

