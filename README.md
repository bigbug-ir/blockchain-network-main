# blockchain-network-main

## installation Go
To install Go on Ubuntu, use the apt-get command:

```bash 
    sudo apt-get update && sudo apt-get -y install golang-go 
```
Output:
```bash
$ sudo apt-get update && sudo apt-get -y install golang-go
Hit:1 http://nova.clouds.archive.ubuntu.com/ubuntu jammy InRelease
Hit:2 http://nova.clouds.archive.ubuntu.com/ubuntu jammy-updates InRelease
Hit:3 http://nova.clouds.archive.ubuntu.com/ubuntu jammy-backports InRelease
Hit:4 http://security.ubuntu.com/ubuntu jammy-security InRelease
Ign:5 https://pkg.jenkins.io/debian-stable binary/ InRelease
Hit:6 https://pkg.jenkins.io/debian-stable binary/ Release
Reading package lists... Done
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
The following additional packages will be installed:
  bzip2 cpp cpp-11 g++ g++-11 gcc gcc-11 gcc-11-base golang-1.18-go golang-1.18-src golang-src libasan6 libatomic1 libc-dev-bin libc-devtools libc6-dev libcc1-0
  libcrypt-dev libdpkg-perl libfile-fcntllock-perl libgcc-11-dev libgd3 libgomp1 libisl23 libitm1 liblsan0 libmpc3 libnsl-dev libquadmath0 libstdc++-11-dev libtirpc-dev
  libtsan0 libubsan1 linux-libc-dev manpages-dev pkg-config rpcsvc-proto
Suggested packages:
  bzip2-doc cpp-doc gcc-11-locales g++-multilib g++-11-multilib gcc-11-doc gcc-multilib autoconf automake libtool flex bison gdb gcc-doc gcc-11-multilib bzr | brz
  mercurial subversion glibc-doc debian-keyring bzr libgd-tools libstdc++-11-doc dpkg-dev
The following NEW packages will be installed:
  bzip2 cpp cpp-11 g++ g++-11 gcc gcc-11 gcc-11-base golang-1.18-go golang-1.18-src golang-go golang-src libasan6 libatomic1 libc-dev-bin libc-devtools libc6-dev libcc1-0
  libcrypt-dev libdpkg-perl libfile-fcntllock-perl libgcc-11-dev libgd3 libgomp1 libisl23 libitm1 liblsan0 libmpc3 libnsl-dev libquadmath0 libstdc++-11-dev libtirpc-dev
  libtsan0 libubsan1 linux-libc-dev manpages-dev pkg-config rpcsvc-proto
0 upgraded, 38 newly installed, 0 to remove and 14 not upgraded.
Need to get 143 MB of archives.
After this operation, 634 MB of additional disk space will be used.
...
Scanning processes...                                                                                                                                                       
Scanning candidates...                                                                                                                                                      
Scanning linux images...                                                                                                                                                    

Restarting services...
Service restarts being deferred:
 systemctl restart systemd-logind.service

No containers need to be restarted.

No user sessions are running outdated binaries.

No VM guests are running outdated hypervisor (qemu) binaries on this host.
```
```bash
    go version 
```
Output:
```bash
$ go version
    go version go1.18.1 linux/amd64
```

## clone the repository
```bash
    git clone https://github.com/bigbug-ir/blockchain-network-main.git
```


## Run 
Init modules:
```bash
    go mod init main
```
Output:
```bash
    go: creating new go.mod: module main
```
choose directory /cmd/web:
```bash
    cd ./cmd/web/
```
run:
```bash
    go run .
```
Output:
```bash
 go run .
Hash of the block 52fda72036bee99bf540e4cfda04d1cc6f31af78d3ca4721891b30f1011bfe90
Hash of the previous Block: 
All the transactions: Genesis Block
Hash of the block d8c0134d6311eec995b514befc0c4950987fcc6add4d6fb7875853338146fe9c
Hash of the previous Block: 52fda72036bee99bf540e4cfda04d1cc6f31af78d3ca4721891b30f1011bfe90
All the transactions: first transaction
Hash of the block cd7d4028d0322d3ce8ccf206f813daaa6328a204d953a4087b5ce801a9fb6480
Hash of the previous Block: d8c0134d6311eec995b514befc0c4950987fcc6add4d6fb7875853338146fe9c
All the transactions: Second transaction
```