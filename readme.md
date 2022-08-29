# ntnft
**ntnft** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve -r
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

NOTE: open another terminal to use the `ntnftd`.

## Check the demo app tx commands
### Create class

This command will create a class named "first-class".

`$ ntnftd tx ntnft create-class first-class "" "" "" 10token --from bob --chain-id=ntnft`

Check that class exists:

`$ ntnftd q ntnft list-class`


### Mint an NT-NFT for Bob
Create a token from class for Bob
```
$ ntnftd tx ntnft mint "1" --from bob --chain-id=ntnft

{
	"body": {
		"messages": [
			{
				"@type": "/ntnft.ntnft.MsgMint",
				"creator": "cosmos132m645zmy9a4epsepcu5f6m0gvtxm9p5cl0p43",
				"class_id": "0"
			}
		],
		"memo": "",
		"timeout_height": "0",
		"extension_options": [],
		"non_critical_extension_options": []
	},
	"auth_info": {
		"signer_infos": [],
		"fee": {
			"amount": [],
			"gas_limit": "200000",
			"payer": "",
			"granter": ""
		}
	},
	"signatures": []
}
```

### Query cmds
Available query commands are listef below.

```
$ ntnftd q ntnft list-class
$ ntnftd q ntnft list-ntnft
$ ntnftd q ntnft list-nt-nft
$ ntnftd q ntnft list-owner
```
