#!/bin/bash
if [[ `curl -s wttr.in/ННовгород | head -n 3 | grep rain` ]]; then
	systemctl stop smb
    tar -cf /var/lib/samba/usershare/backups/$(date +%Y-%m-%d).tar.gz /var/lib/samba/usershare/goserver
    systemctl start smb     
fi
