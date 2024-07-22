<img width="1200" alt="Labs" src="https://user-images.githubusercontent.com/99700157/213291931-5a822628-5b8a-4768-980d-65f324985d32.png">

<p>
 <h3 align="center">Chainstack is the leading suite of services connecting developers with Web3 infrastructure</h3>
</p>

<p align="center">
  <a target="_blank" href="https://chainstack.com/build-better-with-ethereum/"><img src="https://github.com/soos3d/blockchain-badges/blob/main/protocols_badges/Ethereum.svg" /></a>&nbsp;  
  <a target="_blank" href="https://chainstack.com/build-better-with-bnb-smart-chain/"><img src="https://github.com/soos3d/blockchain-badges/blob/main/protocols_badges/BNB.svg" /></a>&nbsp;
  <a target="_blank" href="https://chainstack.com/build-better-with-polygon/"><img src="https://github.com/soos3d/blockchain-badges/blob/main/protocols_badges/Polygon.svg" /></a>&nbsp;
  <a target="_blank" href="https://chainstack.com/build-better-with-avalanche/"><img src="https://github.com/soos3d/blockchain-badges/blob/main/protocols_badges/Avalanche.svg" /></a>&nbsp;
  <a target="_blank" href="https://chainstack.com/build-better-with-solana/"><img src="https://github.com/soos3d/blockchain-badges/blob/main/protocols_badges/Solana.svg" /></a>&nbsp;
</p>

<p align="center">
  <a target="_blank" href="https://chainstack.com/protocols/">Supported protocols</a> •
  <a target="_blank" href="https://chainstack.com/blog/">Chainstack blog</a> •
  <a target="_blank" href="https://docs.chainstack.com/">Chainstack docs</a> •
  <a target="_blank" href="https://docs.chainstack.com/api/">Blockchain API reference</a> •
  <a target="_blank" href="https://console.chainstack.com/user/account/create">Start for free</a>
</p>

# Using Chainlink data feeds with Foundry

This project shows how to use Foundry to compile and deploy Smart contracts. 

See the full tutorial on the Chainstack blog:
* [Using Chainlink data feeds with Foundry](https://chainstack.com/using-chainlink-data-feeds-with-foundry/)

## Project details

Foundry is one of the latest smart contract development tools currently in the market, and it allows users to compile contracts, write tests, deploy contracts, and much more through its command line interface. This project is already set up for you to deploy smart contracts and interact with Chainlink’s data feeds.

Learn more about Foundry: Foundry: [A fast Solidity smart contract development toolkit](https://chainstack.com/foundry-a-fast-solidity-contract-development-toolkit/)

## Quick start

Clone this repository.

### Install Foundry

```sh
curl -L https://foundry.paradigm.xyz | bash
```

Then:

```sh
foundryup
```

### Compile smart contracts

```sh
forge build
```

### Create a .env file

In the root directory:

```sh
touch .env
```

This will create a new `.env` file, paste and edit the following in it:

```env
GOERLI_RPC_URL=CHAINSTACK_GOERLI_NODE_URL
PRIVATE_KEY=YOU_PRIVATE_KEY
ETHERSCAN_KEY=ETHERSCAN_API_KEY(to verify the smart contract)
```

### Deploy the smart contract

Run the following command:

```sh
forge script script/priceFeedsScript.s.sol:ChainlinkScript --rpc-url $GOERLI_RPC_URL  --private-key $PRIVATE_KEY --broadcast --verify --etherscan-api-key $ETHERSCAN_KEY -vvvv
```

## Prerequisites

* Linux or MacOS
* A Goerli Chainstack endpoint

Deploy a Goerli node:
1. [Sign up with Chainstack](https://console.chainstack.com/user/account/create).  
1. [Deploy a node](https://docs.chainstack.com/platform/join-a-public-network).  
1. [View node access and credentials](https://docs.chainstack.com/platform/view-node-access-and-credentials). 

## Dependencies

* [Foundry toolkit](https://github.com/foundry-rs/foundry)

## Install

Clone this repository.

Install Foundry

```sh
curl -L https://foundry.paradigm.xyz | bash
```

Then:

```sh
foundryup
```

## Example

Here is a link to a verified smart contract deployed to the Goerli testnet, and then verified using Foundry's command line.
* [0xda96bbe0b02e64d374bd98355a91405653b081e2](https://goerli.etherscan.io/address/0xda96bbe0b02e64d374bd98355a91405653b081e2)
