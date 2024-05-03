![shellHub's FreeBSD Port logo](./shellhub/logo.png)

FreeBSD Port for ShellHub Agent.

## Build

Clone the FreeBSD Ports Collection repository to your local system:

```sh
git clone https://github.com/shellhub-io/ports.git
```

Change into the directory containing the ShellHub Port:

```sh
cd ports
```

Build and install ShellHub using the `make` command:

```sh
make install clean
```

Configure the required variables in `/etc/rc.conf` to start the agent:

```conf
# ShellHub Agent configuration.

shellhub_enable="YES"
shellhub_server_address="https://cloud.shellhub.io/"
shellhub_tenant_id="00000000-0000-4000-0000-000000000000"
```

> [!NOTE]
> The `shellhub_server_address` and `shellhub_tenant_id` are placeholder values and should be changed to your
> desirable instance and the equivalent Tenant ID.

Start the service:

```sh
service shellhub start
```

For further configuration and usage instructions, refer to the [documentation](https://docs.shellhub.io).

## Contribute

For information about contributing to FreeBSD ports, please see the [Porter's
Handbook](https://docs.freebsd.org/en/books/porters-handbook/).
