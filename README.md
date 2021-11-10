# Fantom NFT Sales

So first of all, I'm really sorry for this code haha this is my first project in Go, so I know 
there's a lot of improvements to do.

## About
This is a bot to tweet about the NFT sales in Fantom Network.
Currently, only works with at least one sale after mint.

## Development
```bash
# Twitter credentials
cp export_vars.sh-example export_vars.sh
# Add your credentials
source export_vars.sh

# Download dependencies
go add .

# Run project
go run .
```