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

`$ ntnftd tx ntnft create-class name=first-class uri="" uri-hash="" data="" price=1 --from bob --chain-id=ntnft`

Check that class exists:

`$ ntnftd q ntnft list-class`


### Mint an NT-NFT for Bob
Create a token from class for Bob
```
$ ntnftd tx ntnft mint "0" --from bob --chain-id=ntnft

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

Now poor Bob cannot get rid of the NT-NFT :(.


### Check query cmds
Check that the NT-NFT exists and is properly handled (assigned to bob)

```
$ ntnftd q ntnft list-class
$ ntnftd q ntnft list-ntnft
$ ntnftd q ntnft list-nt-nft
$ ntnftd q ntnft list-owner
```
