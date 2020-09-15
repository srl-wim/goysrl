## generaate code

```bash 
generator -path=yang/ietf/ -generate_fakeroot -fakeroot_name device -package_name goysrl yang/srl/*
```

## yang prepare

cp ~/CodeProjects/yang/srlinux/20_06_1/yang/modules/models/ietf/*.yang ietf/
cp ~/CodeProjects/yang/srlinux/20_06_1/yang/modules/models/iana/*.yang ietf/
cp ~/CodeProjects/yang/srlinux/20_06_1/yang/modules/models/srl_nokia/models/*/*.yang srl/


remove default in gnmi/snmp/json-rpc models
remove srl-ext import statement and lines/
remove srl_nokia-ext mport statement and lines/
remove tools yang models
