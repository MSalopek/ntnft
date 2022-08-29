// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgMint } from "./types/ntnft/tx";
import { MsgEditClass } from "./types/ntnft/tx";
import { MsgEditToken } from "./types/ntnft/tx";
import { MsgCreateClass } from "./types/ntnft/tx";
import { MsgRemoveToken } from "./types/ntnft/tx";
import { MsgCreateModuleAccountClass } from "./types/ntnft/tx";


const types = [
  ["/ntnft.ntnft.MsgMint", MsgMint],
  ["/ntnft.ntnft.MsgEditClass", MsgEditClass],
  ["/ntnft.ntnft.MsgEditToken", MsgEditToken],
  ["/ntnft.ntnft.MsgCreateClass", MsgCreateClass],
  ["/ntnft.ntnft.MsgRemoveToken", MsgRemoveToken],
  ["/ntnft.ntnft.MsgCreateModuleAccountClass", MsgCreateModuleAccountClass],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgMint: (data: MsgMint): EncodeObject => ({ typeUrl: "/ntnft.ntnft.MsgMint", value: MsgMint.fromPartial( data ) }),
    msgEditClass: (data: MsgEditClass): EncodeObject => ({ typeUrl: "/ntnft.ntnft.MsgEditClass", value: MsgEditClass.fromPartial( data ) }),
    msgEditToken: (data: MsgEditToken): EncodeObject => ({ typeUrl: "/ntnft.ntnft.MsgEditToken", value: MsgEditToken.fromPartial( data ) }),
    msgCreateClass: (data: MsgCreateClass): EncodeObject => ({ typeUrl: "/ntnft.ntnft.MsgCreateClass", value: MsgCreateClass.fromPartial( data ) }),
    msgRemoveToken: (data: MsgRemoveToken): EncodeObject => ({ typeUrl: "/ntnft.ntnft.MsgRemoveToken", value: MsgRemoveToken.fromPartial( data ) }),
    msgCreateModuleAccountClass: (data: MsgCreateModuleAccountClass): EncodeObject => ({ typeUrl: "/ntnft.ntnft.MsgCreateModuleAccountClass", value: MsgCreateModuleAccountClass.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
