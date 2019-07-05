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

```bash
echo "deb https://dl.bintray.com/foilen/debian stable main" | sudo tee /etc/apt/sources.list.d/foilen.list
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61
sudo apt update
sudo apt install sendmail-to-msmtp
```
