## notes

copy weberser.service to /etc/systemd/system#

sudo systemctl daemon-reload

sudo systemctl enable webserver.service

sudo systemctl is-enabled webserver.service

### or nginix

sudo systemctl disable webserver.service
sudo systemctl enable nginx.service

sudo systemctl is-enabled webserver.service





### nginix on aws ami
sudo yum update
sudo yum install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo yum-config-manager --enable epel
sudo yum install certbot python3-certbot-nginx
certbot --version



https://user:auth@domains.google.com/nic/update?hostname=lucinda.dilworth.uk&myip=86.138.131.104


using new DDNS cleint
https://github.com/troglobit/inadyn/blob/master/README.md

### dev notes

- js main js has timers for page load


bfg --delete-files assets/work/paros/2495.mov  my-repo.git
assets/work/paros/2495.mov
