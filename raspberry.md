### notes about pi setups

### useful commands

systemctl --type=service --state=active

### Cloudfared tunnel

my hosts file
 - dilly01 2GB RPi kiosk 10.1.1.53
  - dilly02 8 8GB RPi kiosk 10.1.1.52
  - dilly03 raspbery pi 3 10.1.1.54


for RDP from mac run

cloudflared access rdp --hostname rdp.dilworth.uk --url rdp://localhost:3389

i have the plist file in sys dir for this

plst
these are local in ~/Lbrary/LaunchAgents

launchctl load cloudflared-rdp.plist
launchctl start cloudflared-rdp.plist


ssh config setup on my mac for pis

Host dilly01
  ProxyCommand /opt/homebrew/bin/cloudflared access ssh --hostname %h
Host dilly02
  ProxyCommand /opt/homebrew/bin/cloudflared access ssh --hostname %h
Host dilly03
  ProxyCommand /opt/homebrew/bin/cloudflared access ssh --hostname %h


### dilly network setups



  - 4G modem 192.168.8.1
  - network 10.1.1.0
  - netmask 255.255.255.0
  - gw 10.1.1.1
  - WIFI SSDN pickle-net
  - WIFI PASS dillpickle
  - dilly01 2GB RPi kiosk 10.1.1.53
  - dilly02 8 8GB RPi kiosk 10.1.1.52
  - dilly03 raspbery pi 3 10.1.1.54
  - dilly100 win 11 pro workstatiom 10.1.1.100
  - dilly101 Datapath Fx4 10.1.1.101
  - pickel01
  - pickle02 - exhit old
  - pickle03 - 2nd demo with scratch
  - pickle04 - tek one
  - pickle05 - newversion of optoma
  - pickle06 - latest version - gucci one
  - Mike Mac Air 10.1.1.200
  - DHCP range for one offs 10.1.1.150 - 10.1.1.170




https://pimylifeup.com/raspberry-pi-cloudflare-tunnel/

TP-LINK wifi
Firmware Version:	
3.16.9 Build 20170628 Rel.37804n
Hardware Version:	
TL-WR902AC v1 00000000
TL-WR902AC v1 00000000



E3372-325
Current version 3.0.2.61(H057SP7C983)
Web UI version WEBUI 3.0.2.61(W13SP4C7110)

Hardware CL5E3372M
Serial number W6S7S22C27502932



systemctl list-units --type target

sudo systemctl disable bluetooth.target
sudo systemctl disable cryptsetup.target
sudo systemctl disable nfs-client.target
sudo systemctl disable sound.target


#### useful vlc
https://superuser.com/questions/521590/osx-using-cli-version-of-vlc

#### this works
vlc d:\a.mp4 :vout-filter=transform --transform-type=270 --video-filter "transform{true}"

vlc --loop --no-osd --no-audio /Users/mike/Projects/tmp/videos :vout-filter=transform --transform-type=270 --video-filter "transform{true}"

/usr/bin/cvlc --fullscreen --loop --no-osd --no-audio --transform-type{90} /home/mike/Videos/90_digitaldisplaytest.mp4

/usr/bin/cvlc --fullscreen --loop --no-osd --no-audio /home/mike/Videos/90_digitaldisplaytest.mp4 :vout-filter=transform --transform-type=90 --video-filter "transform{true}"

usr/bin/cvlc --fullscreen --loop --no-osd /home/mike/Videos :vout-filter=transform --transform-type=270 --video-filter "transform{true}"

kiosk.service in /etc/systemd/system/
the sudo systemctl enable kiosk

[Unit]
Description=lubyKiosk

[Service]
User=mike
ExecStart=/usr/bin/cvlc --fullscreen --loop --no-osd /home/mike/Videos
WorkingDirectory=/home/mike/Videos

[Install]
WantedBy=multi-user.target


### Server

[Unit]
Description=tunnel

[Service]
User=ubuntu
ExecStart=/home/ubuntu/frp/frps -c /home/ubuntu/frp/frps.toml
WorkingDirectory=/home/ubuntu/frp

[Install]
WantedBy=multi-user.target


### Client
sudo vi /etc/systemd/system/tunnel.service

[Unit]
Description=tunnel

[Service]
User=mike
ExecStart=/home/mike/frp/frpc -c /home/mike/frp/frpc.toml
WorkingDirectory=/home/mike/frp

[Install]
WantedBy=multi-user.target


sudo systemctl enable tunnel.service



ssh -oPort=6000 mike@www.lucindadilworth.com

Host dilly02
        HostName www.lucindadilworth.com
        IdentityFile ~/.ssh/id_rsa
        ProxyCommand ssh -oPort=6000 --hostname %h
        User mike