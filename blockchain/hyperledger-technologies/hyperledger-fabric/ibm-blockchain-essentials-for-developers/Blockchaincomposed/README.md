week 4


Blockchain Composed

Hyperledger composer - open source project. A suite of high level application abstractions for business network.


Benefits of Composer
1. Increases understanding.  (Bridges simply from business concepts to blockchain)
2. Saves time. (Develop blockchain applications more quickly and cheaply)
3. Reduces risk. (Well tested, efficient, design )
4. Increases flexibility (Higher level abstraction)


An Example of Business Network

Car Auction Market. 

Asset - Vehicle (Vin)
Transactions
Participants



Conceptual Components and Structure
1. Business network - defined by Models, Script Files, ACL's, and Metadata and packaged in a Business Network Archive.
2. Models - Assets, Transactions, Participants Schema. Contains the shape of data stored in transactions as a schema.
3. Script files - Business Logic.(smart contracts). Transaction processing functions. Business logic running as part of blockchain network on hyperledger fabric. At the runtime, when a transaction is submitted, it's going to find the script functions that are interested in the transaction, and run them. Script functions can have side effects on assets that are being managed in asset registries.
4. ACL's. Access control rules. Set of participants have control rules on the access of assets etc.

