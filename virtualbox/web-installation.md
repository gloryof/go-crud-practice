
## firewalldの設定
```
# firewall-cmd --permanent --new-service=go-crud

# firewall-cmd --permanent --service=go-crud --set-short=go-crud
# firewall-cmd --permanent --service=go-crud --set-description=go-crud
# firewall-cmd --permanent --service=go-crud --add-port=80/tcp

# firewall-cmd --permanent --add-service=go-crud --zone=public

# firewall-cmd --reload
```

## sudoersの設定
```
# vi /etc/sudoers.d/app
```

設定内容。
```
app ALL=(ALL) NOPASSWD: /usr/bin/systemctl start go-crud
app ALL=(ALL) NOPASSWD: /usr/bin/systemctl stop go-crud
```