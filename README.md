# genpassword

Simple secure password generator written in Golang.

Never get compromised or be affected by password DB leaks again!

# Build / Install

Linux (Ubuntu and its derivatives)

## Download...

```bash
git clone https://github.com/bluemods/genpassword && cd genpassword
```

## Install...

```bash
go build && chmod +x ./genpassword && sudo cp ./genpassword /usr/local/bin && source ~/.bashrc
```

This builds and copies the program to /usr/local/bin where it can be run anywhere with the `genpassword` command on your terminal.

## Example usage

`genpassword` generates 32 random alphanumeric characters (may include ASCII symbols)
`genpassword -a -s 64` generates 64 random alphanumeric characters 
