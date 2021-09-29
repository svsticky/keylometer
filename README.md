# Keylometer
> how to not update ssh keys and instead update a config

Long story short, keylometer provides an AuthorizedKeysCommand for `sshd`,
allowing us to not have to update our ssh keys ever. Just add a user to 
the config ,and they can use our server.

The `keylometer.yml` file is an example of the config file. It is tried first
from `/etc/keylometer.yml`, and then in the current folder. 

To make it work, build the program using `go build`, then copy it to a folder
that everyone can access. Then add the following to `/etc/ssh/sshd_config`:
```
AuthorizedKeysCommand /path/to/keylometer %u
AuthorizedKeysCommandUser nobody
```
The `%u` is important so Keylometer can figure out what keys to use for the
current user. 