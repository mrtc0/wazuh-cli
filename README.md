# wazuh-cli

⚠️ WIP ⚠️


---

wazuh-cli is a command line interface for  Wazuh REST API.

## Example

```powershell
$ wazuh-cli agent list
ID                  STATUS              NAME                IP                  OS
000                 Active              wazuh-manager       127.0.0.1           Ubuntu
001                 Active              backend-api         10.142.0.3          Ubuntu
002                 Active              lxd-1               10.142.0.5          Ubuntu
003                 Active              frontend            10.142.0.2          Ubuntu
004                 Active              redis               10.142.0.4          Ubuntu
```

```powershell
$ wazuh-cli syscollector processes 001
User                         PID       STAT     Name                CMD
root                         1218      S        supervisord         /usr/bin/python
root                         1258      S        lxcfs               /usr/bin/lxcfs
root                         1275      S        accounts-daemon     /usr/lib/accountsservice/accounts-daemon
root                         1306      S        systemd-logind      /lib/systemd/systemd-logind
root                         1404      S        snapd               /usr/lib/snapd/snapd
root                         1420      S        sshguard            /usr/sbin/sshguard
root                         1423      S        mackerel-agent      /usr/bin/mackerel-agent
root                         1441      S        polkitd             /usr/lib/policykit-1/polkitd
root                         1535      S        mackerel-agent      /usr/bin/mackerel-agent
mysql                        1649      S        mysqld              /usr/sbin/mysqld
root                         1719      S        mdadm               /sbin/mdadm
root                         1723      S        google_network_     /usr/bin/python3
root                         1724      S        google_clock_sk     /usr/bin/python3
root                         1726      S        google_accounts     /usr/bin/python3
```

## Operations

### Agent

```
COMMANDS:
     list  show all agents
     get   show information an agents
```

### Syscollector

```
COMMANDS:
     hardware   show hardware information
     address    show network address information
     packages   show all packages
     ports      show all listen ports
     processes  show all processes
```

## Build

```
$ make build
```

## Usage

```
$ cat ~/.config/wazuh-cli/config.toml
APIURL = "https://host:55000/"

[BasicAuth]
Username = "foo"
Password = "bar"
```
