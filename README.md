### Задание:
Поднимите самбу(samba). Поднимите любой веб-сервер(nginx, apache, lighttpd, cherokee). пошарьте папку со страницами(или переназначьте веб сервер на пошаренную папку). Сделайте так чтобы вы могли коммитить в эту-же папку гитом. Делайте бекап кода веб страниц в 05-00 каждый день, если день не дождливый. Учитывайте что бекап кода может быть поврежден если в момент бекапа кто-то пишет по самбе.

### Последовательность шагов:

1. Первым шагом предполагается установка samba. В manjaro samba вшита изначально, иначе можно установить pacman -S samba
2. Вторым шагом я создаю группу smbgrp `groupadd smbgrp`, добавляю туда пользователя araxal `usermod araxal -aG smbgrp` и добавляю пользователя в smbpasswd с заданием ему пароля в cli `smbpasswd -a araxal`
3. Третьим шагом я создаю smb.conf в etc/samba с данным содержимым : 
```
[global]
workgroup = WORKGROUP
netbios name = manjaro
security = user

[secret]
comment = Git File Server Share
path = /var/lib/samba/usershare
valid users = @smbgrp
browsable =yes
writable = yes
guest ok = no
read only = no
```
4. После командой testparm проверяю, что конфиг файл задан правильно и перезапускаю самбу через systemctl restart smb
5. Далее на windows компьютере я ввожу имя сети (manjaro), ввожу логин и пароль и убеждаюсь что все работает.
6. В расшаренной папке инициализирую гит репозиторий git init ( так и не понял зачем он нам нужен )
7. Следующим шагом создаю сервер в 3 строчки на го [server](goserver), отдающий статический html по адресу localhost:8080 и кидаю в расшаренную папку.
8. Создаю папку backups, в которую есть доступ на запись только у araxal. Остальным на чтение
9. Создаю скрипт [backuper.sh](backuper.sh)  и новую задачу crontab ( файлик в папке /etc/cron.d с содержимым [run_backuper](run_backuper)

