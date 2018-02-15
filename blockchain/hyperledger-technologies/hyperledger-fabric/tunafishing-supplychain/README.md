We will Demonstrate end to end tuna fishing supply chain management using hyperledger fabric blockchain network.


With the Tuna 2020 Traceability Declaration in mind, our goal is to eliminate illegal, unreported, and unregulated fishing.
We will be using Hyperledger Fabric's framework to keep track of each part of this process.


Actors in the Network:-
    Sarah - is the fisherman who sustainably and legally catches tuna.
    Regulators - verify that the tuna has been legally and sustainably caught.
    Miriam - is a restaurant owner who will serve as the end user, in this situation.
    Carl - is another restaurant owner fisherman Sarah can sell tuna to.

Roles of Actors in the network

1. Catching Tuna and Recording (Seller)

    We will start with Sarah, our licensed tuna fisher, who makes a living selling her tuna to multiple restaurants. Sarah operates as a private business, in which her company frequently makes international deals. Through a client application, Sarah is able to gain entry to a Hyperledger Fabric blockchain network comprised of other fishermen, as well as regulators and restaurant owners. Sarah has the ability to add to and update information in the blockchain network's  ledger as tuna pass through the supply chain, while regulators and restaurants have read access to the ledger.

    After each catch, Sarah records information about each individual tuna, including: a unique ID number, the location and time of the catch, its weight, the vessel type, and who caught the fish. For the sake of simplicity, we will stick with these six data attributes. However, in an actual application, many more details would be recorded, from toxicology, to other physical characteristics.

    These details are saved in the world state as a key/value pair based on the specifications of a chaincode contract, allowing Sarah’s application to effectively create a transaction on the ledger. You can see the example below:

    $ var tuna = { id: ‘0001’, holder: ‘Sarah’, location: { latitude: '41.40238', longitude: '2.170328'}, when: '20170630123546', weight: ‘58lbs’, vessel : ‘9548E’ }


2. Restaurant owners purchasing tuna (Buyer)
    Miriam is a restaurant owner looking to source low cost, yet high quality tuna that have been responsibly caught. Whenever Miriam buys tuna, she is always uncertain whether she can trust that the tuna she is purchasing is legally and sustainably caught, given the prominence of illegal and unreported tuna fishing.

    At the same time, as a legitimate and experienced fisherman, Sarah strives to make a living selling her tuna at a reasonable price. She would also like autonomy over who she sells to and at what price.


3. The Sale
    Normally, Sarah sells her tuna to restaurateurs, such as Carl, for $80 per pound. However, Sarah agrees to give Miriam a special price of $50 per pound of tuna, rather than her usual rate. In a traditional public blockchain, once Sarah and Miriam have completed their transaction, the entire network is able to view the details of this agreement, especially the fact that Sarah gave Miriam a special price. As you can imagine, having other restaurateurs, such as Carl, aware of this deal is not economically advantageous for Sarah.
    
    To remedy this, Sarah wants the specifics of her deal to not be available to everyone on the network, but still have every actor in the network be able to view the details of the fish she is selling. Using Hyperledger Fabric's feature of channels, Sarah can privately agree on the terms with Miriam, such that only the two of them can see them, without anyone else knowing the specifics.

    Additionally, other fishermen, who are not part of Sarah and Miriam’s transaction, will not see this transaction on their ledger. This ensures that another fisherman cannot undercut the bid by having information about the prices that Sarah is charging different restaurateurs.
    
4. The Regulators

    Regulators will also gain entry to this Hyperledger Fabric blockchain network to confirm, verify, and view details from the ledger. Their application will allow these actors to query the ledger and see the details of each of Sarah’s catches to confirm that she is legally catching her fish. Regulators only need to have query access, and do not need to add entries to the ledger. With that being said, they may be able to adjust who can gain entry to the network and/or be able to remove fishermen from the network, if found to be partaking in illegal activities



Gaining Network Membership
1. Hyperledger Fabric is a permissioned network, meaning that only participants who have been approved can gain entry to the network. To handle network membership and identity, membership service providers (MSP) manage user IDs, and authenticate all the participants in the network. A Hyperledger Fabric blockchain network can be governed by one or more MSPs. This provides modularity of membership operations, and interoperability across different membership standards and architectures.

2. In our scenario, the regulator, the approved fishermen, and the approved restaurateurs should be the only ones allowed to join the network. To achieve this, a membership service provider (MSP) is defined to accommodate membership for all members of this supply chain. In configuring this MSP, certificates and membership identities are created. Policies are then defined to dictate the read/write policies of a channel, or the endorsement policies of a chaincode.

3. Our scenario has two separate chaincodes, which are run on three separate channels. The two chaincodes are: one for the price agreement between the fisherman and the restaurateur, and one for the transfer of tuna. The three channels are: one for the price agreement between Sarah and Miriam; one for the price agreement between Sarah and Carl; and one for the transfer of tuna. Each member of this network knows about each other and their identity. The channels provide privacy and confidentiality of transactions.

4. In Hyperledger Fabric, MSPs also allow for dynamic membership to add or remove members to maintain integrity and operation of the supply chain. For example, if Sarah was found to be catching her fish illegally, she can have her membership revoked, without compromising the rest of the network. This feature is critical, especially for enterprise applications, where business relationships change over time.



Summary of Demonstrated Scenario

Below is a summary of the tuna catch scenario presented in this section:

1. Sarah catches a tuna and uses the supply chain application’s user interface to record all the details about the catch to the ledger. Before it reaches the ledger, the transaction is passed to the endorsing peers on the network, where it is then endorsed. The endorsed transaction is sent to the ordering service, to be ordered into a block. This block is then sent to the committing peers in the network, where it is committed after being validated.

2. As the tuna is passed along the supply chain, regulators may use their own application to query the ledger for details about specific catches (excluding price, since they do not have access to the price-related chaincode).

3. Sarah may enter into an agreement with a restaurateur Carl, and agree on a price of $80 per pound. They use the blue channel for the chaincode contract stipulating $80/lb. The blue channel's ledger is updated with a block containing this transaction.

4. In a separate business agreement, Sarah and Miriam agree on a special price of $50 per pound. They use the red channel's chaincode contract stipulating $50/lb. The red channel's ledger is updated with a block containing this transaction.




References-
https://courses.edx.org/courses/course-v1:LinuxFoundationX+LFS171x+3T2017/courseware/f0db5224eb0e4bbb8cc1e93a6819012c/9d0da522dfb246c5bd3dfff09952ff53/?activate_block_id=block-v1%3ALinuxFoundationX%2BLFS171x%2B3T2017%2Btype%40sequential%2Bblock%409d0da522dfb246c5bd3dfff09952ff53

