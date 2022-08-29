# ntnft
**ntnft** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve -r
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

NOTE: open another terminal to use the `ntnftd`.

## x/ntnft module
### Create class

This command will create a class named "first-class".

`$ ntnftd tx ntnft create-class first-class "" "" "" 10token --from bob --chain-id=ntnft`

Check that class exists:

`$ ntnftd q ntnft list-class`


### Mint an NT-NFT for Bob
Create a token from class for Bob

`$ ntnftd tx ntnft mint "1" --from bob --chain-id=ntnft`


### Query cmds
Available query commands are listef below.

```
$ ntnftd q ntnft list-class
$ ntnftd q ntnft list-ntnft
$ ntnftd q ntnft list-nt-nft
$ ntnftd q ntnft list-owner
```

## x/blog module
### Init
* ntnftd tx blog init --from bob --chain-id=ntnft --yes

### Request access
* ntnftd tx blog request-access --from bob --chain-id=ntnft --yes
* confirm: ntnftd q ntnft list-nt-nft

### Post as registered
* ntnftd tx blog create-post first post --from bob --chain-id=ntnft --yes

### Check post
* ntnftd q blog posts

### Create post as unregistered
* ntnftd tx blog create-post second post --from alice --chain-id=ntnft --yes
```
raw_log: 'failed to execute message; message index: 0: address not registered: invalid
  request'

```