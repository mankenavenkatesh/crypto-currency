What is truffle

1. https://github.com/trufflesuite/truffle
2. Truffle is a development environment, testing framework and asset pipeline for Ethereum, aiming to make life as an Ethereum developer easier.


How to install?

1. npm install -g truffle

Initialize project with truffle
1. truffle init


Compiling with truffle
1. truffle compile

Deploying Smart Contract
1. truffle migrations
2. Migrations are scripts which automate series of steps that are needed to deploy contracts, set them up in the series u wanted etc.
3. The migration scripts can be edited to change the order of smart contract deployment.
4. Recently truffle has introducted migration smart contract. which will be deployed on blockchain and it contains the data about which and how the smart contracts are deployed.




Interacting with Contract
1. Once migration is done, the contract address will be provided on terminal.
2. Run truffle console.
3. Run SmartContractObjectName(eg.HelloWorld).deployed()
4. Above command will provide details about the smart contract object. Everything required to interact with smart contract will be available here.
5. This object can be assigned to a variable and do some operations on it. 


Issues
https://github.com/trufflesuite/truffle/issues/477
Solution
npm install -g truffle-expect truffle-config web3

