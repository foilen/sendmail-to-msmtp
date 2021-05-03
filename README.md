# About

This is a bridge from sendmail to msmtp. The goal is to support all the different ways to provide arguments and translate them to msmtp.

# Configuration of "from"

The "from" email can be defined in multiple places. The latest defined will be used.

- In the */etc/sendmail-to-msmtp.json* configuration file
- In the "-r" argument
- In the "-f" argument
- In the "From: " header

# Local Usage

## Compile

`./create-local-release.sh`

The file is then in `build/bin/sendmail`

## Execute

To execute:
`./build/bin/sendmail`

# Create release

`./create-public-release.sh`

That will show the latest created version. Then, you can choose one and execute:
`./create-public-release.sh X.X.X`

# Use with debian

Get the version you want from https://deploy.foilen.com/sendmail-to-msmtp/ .

```bash
wget https://deploy.foilen.com/sendmail-to-msmtp/sendmail-to-msmtp_X.X.X_amd64.deb
sudo dpkg -i sendmail-to-msmtp_X.X.X_amd64.deb
```

# Configuration file

If you want to set a default email address for the "from" when no other is specified, you can create the */etc/sendmail-to-msmtp.json* configuration file with the following content:
```
{
  "defaultFrom" : "default@foilen-lab.com"
}
```

For debugging purpose, you can dump all the raw emails in a directory. Please note that there is no cleanup of these files; you need to clean them up yourself.
```
{
  "defaultFrom" : "default@foilen-lab.com",
  "emailDumpDirectory" : "/tmp/rawemails"
}
```
