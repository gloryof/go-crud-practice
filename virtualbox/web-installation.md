
## firewalldの設定
```
# firewall-cmd --permanent --new-service=go-crud

# firewall-cmd --permanent --service=go-crud --set-short=go-crud
# firewall-cmd --permanent --service=go-crud --set-description=go-crud
# firewall-cmd --permanent --service=go-crud --add-port=8000/tcp

# firewall-cmd --permanent --add-service=go-crud --zone=public

# firewall-cmd --reload
```