# About

This is a bridge from sendmail to msmtp. The goal is to support all the different ways to provide arguments and translate them to msmtp.

# Configuration of "from"

The "from" email can be defined in multiple places. The latest defined will be used.

- *TODO* In the */etc/sendmail-to-msmtp.json* configuration file
- *TODO* In the "-r" argument
- *TODO* In the "-f" argument
- *TODO* In the "From: " header

# Local Usage

## Compile

`./gradlew goClean goBuild` 

The file is then in `.gogradle/sendmail-to-msmtp-linux-amd64`

## Execute

To execute:
`./.gogradle/sendmail-to-msmtp-linux-amd64`

# Create release

`./create-public-release.sh`

# Use with debian

```bash
echo "deb https://dl.bintray.com/foilen/debian stable main" | sudo tee /etc/apt/sources.list.d/foilen.list
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61
sudo apt update
sudo apt install sendmail-to-msmtp
```
