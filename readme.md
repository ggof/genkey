# GENKEY

## What is it?
`genkey` is a small utility program to generate random bytes and print them to the console. I use it mainly to generate encryption keys, but you may use it to your heart's desire.

## How to use it?
First, make sure you have Go installed on your machine. I would also advise using the latest version of Go. Also make sure `$GOPATH/bin` (also known as `$GOBIN`) in your path.

- Pull the repository with `git clone git@github.com:ggof/genkey.git`
- run `go install`
- Enjoy `genkey` 
---
`genkey` requires exactly one argument: the number of bits you want to generate. The length you use has to be a multiple of 8, as this program only reads random as bytes.

