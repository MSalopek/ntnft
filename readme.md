# Demo App - ntnft
**ntnft** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

This app is intended for demonstration purposes.

The main goal is to explore and showcase how non-transferrable NFTs (aka Soul Bonded tokens) can be used by CosmosSDK apps/modules to grant users access to some module/app features.

In this example you can find the following modules:
* **x/ntnft** - demo implementation of non-transferrable NFTs
* **x/blog** - demo application that depends on `x/ntnft` for registering users and controlling access to app features
  * blog users must first request access to the blog module
  * once requested, access is granted via minting a non-trensferrable NFT and sending it to user
  * while the user holds the nt-NFT they can use the blog application to call `create-post`
  * the user's access can be revoked by simply removing the nt-NFT
    * due to how the app is structured, the `blog` module is allowed to remove any nt-NFTs that were minted to grant access to `blog` features
    * the user can remove their own nt-NFTif they with to no longer have access to app features
## Getting started

```
ignite chain serve -r
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

NOTE: open another terminal to use the `ntnftd`.

# Modules Overview
The app consists of 2 modules `x/ntnft` and `x/blog`. Blog is moddeled after ignite-cli blog from https://docs.ignite.com/.

## x/ntnft
Module tasked with creating token Classes, NtNFT tokens and managing token/class ownership.

The module public interface consists of the following methods:
```golang
// Exported methods implemented by NtNFT Keeper to be used by other modules.
type NtnftKeeperInterface interface {
	// MintToken creates new NtNFT token with classId and assigns it to createAddr.
	MintToken(ctx sdk.Context, classId, createAddr string) (types.NtNft, error)

	// SafeRemoveToken removes NtNFT token provided that the callerAddr is the token owner or class creator.
	SafeRemoveToken(ctx sdk.Context, tokenId, callerAddr string) error

	// SafeEditToken edits NtNFT token provided that the callerAddr is the token owner or class creator.
	// Only Uri, UriHash and Data can be edited on the NtNFT.
	SafeEditToken(ctx sdk.Context, editToken types.NtNft, callerAddr string) (types.NtNft, error)

	// NewClass creates new token class (basically a collection of non-transferrable NFTs).
	NewClass(ctx sdk.Context, class types.Class) types.Class

	// SafeEditClass removes updates Class fields: uri, uri_hash and data.
	// Tokens can only be edited if the callerAddr owns either the NtNFT or the Class.
	// Function should panic if either the Class.Creator or NtNFT.Owner is not set.
	// Function should error if called by callerAddr that does not match either Class.Creator or NtNFT.Owner.
	SafeEditClass(ctx sdk.Context, editClass types.Class, callerAddr string) (types.Class, error)

	// SetModuleAccountClass creates a Class with module account address as Class.Creator.
	SetModuleAccountClass(ctx sdk.Context, name, price, moduleName string) (types.Class, error)

	// GetAllModuleAccountClass gets all Classes where moduleName is set as Class.Creator.
	GetAllModuleAccountClass(ctx sdk.Context, moduleName string) (list []types.Class)

	// OwnerHasClass checks if ownerAddr has token with classId.
	// Should return false if:
	//    - ownerAddr is not registered
	//    - ownerAddr does not own token with classId
	// Returns true when ownerAddr has token with classId.
	OwnerHasClass(ctx sdk.Context, ownerAddr, classId string) bool
}
```

### Example 1: Create Class & mint nt-NFT tokens

This command will create a class named "first-class". The mint price (for minting 1 NtNFT of the Class) is 10token.
Fields `uri`, `uri-hash` represent additional `off-chain` data that is left to the Creators discression. They are ommited in this example.
Field `data` represents volatile `on-chain` data that is left to the Creators discression. It is ommited in this example.

#### Create class

`$ ntnftd tx ntnft create-class first-class "" "" "" 10token --from bob --chain-id=ntnft`

Check that class exists:

`$ ntnftd q ntnft list-class`

#### Mint an NT-NFT for Bob
Create a token belonging to Class "1" for Bob

`$ ntnftd tx ntnft mint "1" --from bob --chain-id=ntnft`

Check that NtNFT exists:
* this will show that Bob's address holds the ntNFT, same as above, but using different store and index.

`$ ntnftd q ntnft list-owner`

### Query cmds
Available query commands are listed below.

```bash
Usage:
  nt-nftd query ntnft [flags]
  nt-nftd query ntnft [command]

Available Commands:
  list-class  list all class
  list-nt-nft list all nt-nft
  list-owner  list all owner
  params      shows the parameters of the module
  show-class  shows a class
  show-nt-nft shows a nt-nft
  show-owner  shows a owner
```

## x/blog module
Blog is inspired by `x/blog` from the ignite cli docs.

In order to use the blog, users must first register their account.

Only registered users can create blog posts (call `CreatePost` successfully).


```bash
blog transactions subcommands

Usage:
  nt-nftd tx blog [flags]
  nt-nftd tx blog [command]

Available Commands:
  create-post    Broadcast message createPost
  init           Broadcast message init
  request-access Broadcast message requestAccess
```


### Init
`$ ntnftd tx blog init --from bob --chain-id=ntnft --yes`

Initializes the blog module.
* creates an Class owned by `blog` ModuleAccount
* can be called only once
* NOTE: for demo purposes, any address can call `init` (this should not be done in prod apps)

### Request access
`$ ntnftd tx blog request-access --from bob --chain-id=ntnft --yes`
* confirm: `$ ntnftd q ntnft list-nt-nft`

Users can request access to the blog module functionality.

The module requires an nt-NFT to be present in the users wallet.

The nt-NFT must have `ClassId` set to the ID/index of a Class controlled by the blog ModuleAccount.

On each operation, the module checks for the existence of the nt-NFT in the users' wallet.

Only registered adresses can create blog posts.

### Create blog Post
`$ ntnftd tx blog create-post first post --from bob --chain-id=ntnft --yes`

If Bob has the requried token, a post named `first` will be created:
* check: `$ ntnftd q blog posts`


### Create post as unregistered
`$ ntnftd tx blog create-post second post --from alice --chain-id=ntnft --yes`

Since alice is not a registered user, she cannot create posts.
```
...
raw_log: 'failed to execute message;
  message index: 0: address not registered: invalid
  request'
...
```
