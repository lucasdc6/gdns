# gdns

gdns is a minimal implementation of a DNS server used for development.

In a few steps, you can configure and run a simple DNS server.

```yaml
# config.yml
- name: asd.com
  type: A
```

```bash
$ gdns -f ./config.yml -p 53
2019/12/31 16:40:39 Server started at 127.0.0.1:53
```

## Development

### Generate a DNS Message

```bash
dig google.com -p 3000 @localhost +tries=1 +time=2
```

### Capture DNS Messages

```bash
sudo tcpdump 'udp port 3000' -i lo -X
```