
Chaincode

    In Hyperledger Fabric, chaincode is the 'smart contract' that runs on the peers and creates transactions. More broadly, it enables users to create transactions in the Hyperledger Fabric network's shared ledger and update the world state of the assets.

    Chaincode is programmable code, written in Go, and instantiated on a channel. Developers use chaincode to develop business contracts, asset definitions, and collectively-managed decentralized applications. The chaincode manages the ledger state through transactions invoked by applications. Assets are created and updated by a specific chaincode, and cannot be accessed by another chaincode.

    Applications interact with the blockchain ledger through the chaincode. Therefore, the chaincode needs to be installed on every peer that will endorse a transaction and instantiated on the channel.

    There are two ways to develop smart contracts with Hyperledger Fabric:

    Code individual contracts into standalone instances of chaincode
    (More efficient way) Use chaincode to create decentralized applications that manage the lifecycle of one or multiple types of business contracts, and let the end users instantiate instances of contracts within these applications.


Chaincode Key APIs
An important interface that you can use when writing your chaincode is defined by Hyperledger Fabric - ChaincodeStub and ChaincodeStubInterface. The ChaincodeStub provides functions that allow you to interact with the underlying ledger to query, update, and delete assets. The key APIs for chaincode include:

    func (stub *ChaincodeStub) GetState(key string) ([]byte, error)
    Returns the value of the specified key from the ledger. Note that GetState doesn't read data from the Write set, which has not been committed to the ledger. In other words, GetState doesn't consider data modified by PutState that has not been committed. If the key does not exist in the state database, (nil, nil) is returned.

    func (stub *ChaincodeStub) PutState(key string, value []byte) error
    Puts the specified key and value into the transaction's Write set as a data-write proposal. PutState doesn't affect the ledger until the transaction is validated and successfully committed.
    
    func (stub *ChaincodeStub) DelState(key string) error
    Records the specified key to be deleted in the Write set of the transaction proposal. The key and its value will be deleted from the ledger when the transaction is validated and successfully committed.




Overview of a Chaincode Program
    When creating a chaincode, there are two methods that you will need to implement:

    Init
    Called when a chaincode receives an instantiate or upgrade transaction. This is where you will initialize any application state.

    Invoke
    Called when the invoke transaction is received to process any transaction proposals.
    As a developer, you must create both an Init and an Invoke method within your chaincode. The chaincode must be installed using the peer chaincode install command, and instantiated using the peer chaincode instantiate command before the chaincode can be invoked. Then, transactions can be created using the peer chaincode invoke or peer chaincode query commands.


Sample Chaincode Decomposed - Dependencies

    Let’s now walk through a sample chaincode written in Go, piece by piece:
    package main

    import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
    )

    The import statement lists a few dependencies that you will need for your chaincode to build successfully.

    fmt - contains Println for debugging/logging
    github.com/hyperledger/fabric/core/chaincode/shim - contains the definition for the chaincode interface and the chaincode stub, which you will need to interact with the ledger, as we described in the Chaincode Key APIs section
    github.com/hyperledger/fabric/protos/peer - contains the peer protobuf package.
    
Sample Chaincode Decomposed - Struct

    type SampleChaincode struct {

    }

    This might not look like much, but this is the statement that begins the definition of an object/class in Go. SampleChaincode implements a simple chaincode to manage an asset.


Sample Chaincode Decomposed - Init Method

    Next, we’ll implement the Init method. Init is called during the chaincode instantiation to initialize data required by the application. In our sample, we will create the initial key/value pair for an asset, as specified on the command line:

    func (t *SampleChaincode) Init(stub shim.ChainCodeStubInterface) peer.Response {

    // Get the args from the transaction proposal

       args := stub.GetStringArgs()

    if len(args) != 2 {

    return shim.Error("Incorrect arguments. Expecting a key and a value")

    }

    // We store the key and the value on the ledger
    err := stub.PutState(args[0], []byte(args[1]))

    if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))

    }
    return shim.Success(nil)

    }

    The Init implementation accepts two parameters as inputs, and proposes to write a key/value pair to the ledger by using the stub.PutState function. GetStringArgs retrieves and checks the validity of arguments which we expect to be a key/value pair. Therefore, we check to ensure that there are two arguments specified. If not, we return an error from the Init method, to indicate that something went wrong. Once we have verified the correct number of arguments, we can store the initial state in the ledger. In order to accomplish this, we call the stub.PutState function, specifying the first argument as the key, and the second argument as the value for that key. If no errors are returned, we will return success from the Init method.


Sample Chaincode Decomposed - Invoke Method
    Now, we’ll explore the Invoke method, which gets called when a transaction is proposed by a client application. In our sample, we will either get the value for a given asset, or propose to update the value for a specific asset.

    func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

    // Extract the function and args from the transaction proposal

    fn, args := stub.GetFunctionAndParameters()
    var result string
    var err error

    if fn == "set" {
    result, err = set(stub, args)
    } else { // assume 'get' even if fn is nil
    result, err = get(stub, args)
    }

    if err != nil { //Failed to get function and/or arguments from transaction proposal
    return shim.Error(err.Error())
    }

    // Return the result as success payload

    return shim.Success([]byte(result))
    }

    There are two basic actions a client can invoke: get and set.

    The get method will be used to query and return the value of an existing asset.
    The set method will be used to create a new asset or update the value of an existing asset.
    To start, we’ll call GetFunctionandParameters to isolate the function name and parameter variables. Each transaction is either a set or a get. Let's first look at how the set method is implemented:

    func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 2 {
    return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
    }

    err := stub.PutState(args[0], []byte(args[1]))

    if err != nil {
    return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }

    return args[1], nil
    }

    The set method will create or modify an asset identified by a key with the specified value. The set method will modify the world state to include the key/value pair specified. If the key exists, it will override the value with the new one, using the PutState method; otherwise, a new asset will be created with the specified value.

    Next, let's look at how the get method is implemented:

    func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {

    if len(args) != 1 {
    return "", fmt.Errorf("Incorrect arguments. Expecting a key")
    }
    value, err := stub.GetState(args[0])
    if err != nil {
    return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if value == nil {
    return "", fmt.Errorf("Asset not found: %s", args[0])
    }
    return string(value), nil
    }

    The get method will attempt to retrieve the value for the specified key. If the application does not pass in a single key, an error will be returned; otherwise, the GetState method will be used to query the world state for the specified key. If the key has not yet been added to the ledger (and world state), then an error will be returned; otherwise, the value that was set for the specified key is returned from the method.



Sample Chaincode Decomposed - Main Function
    The last piece of code in this sample is the main function, which will call the Start function. The main function starts the chaincode in the container during instantiation.

    func main() {
    err := shim.Start(new(SampleChaincode))
    if err != nil {
    fmt.Println("Could not start SampleChaincode")
    } else {
    fmt.Println("SampleChaincode successfully started")
    }
    }

