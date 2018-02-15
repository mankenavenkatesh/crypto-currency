Getting Started (Part I)

In case you haven’t downloaded the education repository for this course, follow the below directions in your terminal window:

$ git clone https://github.com/hyperledger/education.git

$ cd education/LFS171x/fabric-material/tuna-app

Make sure you have Docker running on your machine before you run the next command. If you do not have Docker installed, return to Chapter 4, Technical Requirements.

Also, make sure that you have completed the Installing Hyperledger Fabric section in this chapter before moving on to this application section, as you will likely experience errors. 

First, remove any pre-existing containers, as it may conflict with commands in this tutorial:

$ docker rm -f $(docker ps -aq)

Then, let’s start the Hyperledger Fabric network with the following command:

$ ./startFabric.sh

 

Troubleshooting: If, after running the above you are getting an error similar to the following:

ERROR: failed to register layer: rename
/var/lib/docker/image/overlay2/layerdb/tmp/write-set-091347846 /var/lib/docker/image/overlay2/layerdb/sha256/9d3227c1793b7494e598caafd0a5013900e17dcdf1d7bdd31d39c82be04fcf28: file exists

try running the following command:

$ rm -rf ~/Library/Containers/com.docker.docker/Data/*



Getting Started (Part II)

Install the required libraries from the package.json file, register the Admin and User components of our network, and start the client application with the following commands:

$ npm install

$ node registerAdmin.js

$ node registerUser.js

$ node server.js

Load the client simply by opening localhost:8000 in any browser window of your choice, and you should see the user interface for our simple application at this URL (as in the screenshot below).


Troubleshooting: If you are getting an error similar to the one below while attempting to perform any of the functions on the application:

Error: [client-utils.js]: sendPeersProposal - Promise is rejected: Error: Connect Failed

error from query =  { Error: Connect Failed

   at /Desktop/prj/education/LFS171x/fabric-material/tuna-app/node_modules/grpc/src/node/src/client.js:554:15 code: 14, metadata: Metadata { _internal_repr: {} } }

try running the following commands:

$ cd ~

$ rm -rf .hfc-key-store/

Then, run the commands above starting with:

$ node registerAdmin.js





Query All Tuna Recorded

      // queryAllTuna - requires no arguments
      const request = {
          chaincodeId:’tuna-app’,
          txId: tx_id,
          fcn: 'queryAllTuna',
          args: ['']
          };
      return channel.queryByChaincode(request);
(Reference: The code comes from ..src/queryAllTuna.js)

Now, let’s query our database, where there should be some sample entries already, since our chaincode smart contract initiated the ledger with 10 previous catches. This function takes no arguments, as we see on line 6. Instead, it takes an empty array.

The query response you should see in the user interface is 10 pre-populated entries with the attributes for each catch.




Query a Specific Tuna Recorded

      // queryTuna - requires 1 argument
      const request = {
          chaincodeId:’tuna-app’,
          txId: tx_id,
          fcn: 'queryTuna',
          args: ['1']
          };
      return channel.queryByChaincode(request);
(Reference: The code comes from ..src/queryTuna.js)

Now, let’s query for a specific tuna catch. This function takes 1 argument, as you can see on line 6 above, an example would be ['1']. In this example, we are using the key to query for catches.

You should see the following query response detailing the attributes recorded for one particular catch.



Change Tuna Holder

      // changeTunaHolder - requires 2 argument
      var request = {
          chaincodeId:’tuna-app’,
          fcn: 'changeTunaHolder', 
          args: ['1', 'Alex'],
          chainId: 'mychannel',
          txId: tx_id
          };
      return channel.sendTransactionProposal(request);
(Reference: The code comes from ..src/changeHolder.js)

Now, let’s change the name of the person in possession of a given tuna. This function takes 2 arguments: the key for the particular catch, and the new holder, as we can see on line 5 in the example above. Ex: args: ['1', 'Alex'].

You may be able to see a similar success response in your terminal window:

The transaction has been committed on peer localhost:7053
 event promise all complete and testing complete

Successfully sent transaction to the orderer.
Successfully sent Proposal and received ProposalResponse: Status - 200, message - "OK", metadata - "", endorsement signature: 0D 9

This indicates we have sent a proposal from our application via the SDK, and the peer has been endorsed, committed, and the ledger has been updated.

Fabric application change tuna holder

You should see that the holder has indeed been changed by querying for key ['1'] again. Now, the holder attribute has been changed from Miriam to Alex, for example.

Fabric application change record




Record a Tuna Catch
      // recordTuna - requires 5 argument
      var request = {
          chaincodeId:’tuna-app’,
          fcn: 'recordTuna',   
          args: ['11', '239482392', '28.012, 150.225', '0923T', "Hansel"],
          chainId: 'mychannel',
          txId: tx_id
          };
      return channel.sendTransactionProposal(request);
(Reference: The code comes from ..src/recordTuna.js)

Lastly, we will practice recording a new tuna catch, and adding it to the ledger by invoking the recordTuna function. This function takes 5 arguments, itemizing each of the attributes of a new catch. You can see an example submission on line 5: args: ['11','239482392', '28.012, 150.225', '0923T', "Hansel"].

Fabric application create tuna record

Check and you should see that the holder has indeed been changed by querying all the tuna catches. Now, you should see an additional entry at the bottom of the table:

Fabric application query after creating tuna record




Finishing Up

Remove all Docker containers and images that we created in this tutorial with the following command in the tuna-app folder:

$ docker rm -f $(docker ps -aq)

$ docker rmi -f $(docker images -a -q)

