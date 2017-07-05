Tools required to build Smart Contracts on Ethereum


SmartContract Online Compiler
1. https://ethereum.github.io/browser-solidity/#version=soljson-v0.4.12+commit.194ff033.js



Every Operation of smart contract will cost a gas.
Any state change will cost more gas.


Testing Smart Contracts

We can test the smart contracts by setting test network using TESTRPC.
Deploying smart contract on testrpc.
using interface provided, execute operations on smart contract.

or
Ethereum Provides Live Test Network where people around world are interacting with.On a live network,
Every transaction has to be put in a block.
Every block has to be mined. 
Then smart contracts has to be deployed. This is time taking.


What is testrpc
https://github.com/ethereumjs/testrpc
1. testrpc is a simulated blockchain that runs in memory on local computer.
2. Deploying in RAM is much faster than live network.  

Install and start testrpc
1. npm install -g ethereumjs-testrpc
2. run testrpc
 
How authentication works with testrpc
1. when any smart contract is deployed, testrpc considers account0 for authentication.
2. the public and private keys for account0 can be seen on terminal when testrpc is executed.

Deploying contract
option 1
1. you can deploy from the browser. create solidity contract online from below link.
2. It gets automatically compiled by solidity compiler.
3. Click on create. the contract will be deployed by the compiler. Done by browser.
2. https://ethereum.github.io/browser-solidity/#version=soljson-v0.4.12+commit.194ff033.js

option 2
1. trufle can be used to deploy contract on local network.
2. trufle takes the smart contract, compiles it. take the byte code. deploying the byte code.
3. all the annoying stuff is taken care by trufle. only smart contract has to be provided.
4. More on truffle in another folder.


How to interact with the deployed contract?
1. once the contract is deployed, an address is generated for the contract.
2. In above example, the contract address is displayed with the name. copy the address.
3. using call function on Smart Contract Object, we can get information about the state of smart contract.

References
1. https://www.youtube.com/watch?v=8jI1TuEaTro
2. https://www.ethereum.org/foundation



