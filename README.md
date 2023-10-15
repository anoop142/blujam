# blujam
A bluetooth stress testing tool  using l2ping.

### Based on [BLUETOOTH-DOS-ATTACK-SCRIPT](https://github.com/crypt0b0y/BLUETOOTH-DOS-ATTACK-SCRIPT)

Requires bluez l2ping



## Usage
```shell
sudo ./blujam  -d <target addr> [-t thread count] [-p packet size]
```

## Example
```
sudo ./blujam -d A4:81:HC:03:60:47 -t 5
```
## Params
```
  -d string
    	device address
  -p int
    	packet size (default 600)
  -t int
    	threads (default 10)

```
