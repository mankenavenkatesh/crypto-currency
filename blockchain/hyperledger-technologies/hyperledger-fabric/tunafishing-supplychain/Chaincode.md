Setting the Stage
    Now that we have a general idea of how chaincode is coded, we will walk through a simple chaincode that creates assets on a ledger, based on our demonstrated scenario of creating records for tuna fish.

    Sometimes, code snippets can get lost in translation, especially if the context doesn’t make much sense. In hopes of avoiding this, we have adjusted our example chaincode to address our demonstration scenario. The chaincode we will be examining in this section will record a tuna catch by storing it to the ledger, as well as allow for queries and updates to tuna catch records.
    
Defining the Asset Attributes

    Here are the four example attributes of tuna fish that we will be recording on the ledger:

    Vessel (string)
    Location (string)
    Date and Time (datetime)
    Holder (string)
    
    We create a Tuna Structure that has four properties. Structure tags are used by the encoding/json library.

    type Tuna struct {
    Vessel string ‘json:"vessel"’
    Datetime string ‘json:"datetime"’
    Location string ‘json:"location"’
    Holder string ‘json:"holder"’

    }
    
Invoke Method (Part I)
    As described earlier, the Invoke method is the one which gets called when a transaction is proposed by a client application. Within this method, we have three different types of transactions -- recordTuna, queryTuna, and changeTunaHolder, which we will look at a little later.

    As a reminder, Sarah, the fisherman, will invoke the recordTuna when she catches each tuna.
    
    changeTunaHolder can be invoked by Miriam, the restaurateur, when she confirms receiving and passing on a particular tuna fish as it passes through the supply chain. queryTuna can be invoked by Miriam, the restaurateur, to view the state of a particular tuna.

    Regulators will invoke queryTuna and queryAllTuna based on their need to verify and check for sustainability of the supply chain.

Invoke Method (Part II)

    We’ll be getting into the different tuna chaincode methods in the following sections. But here is the Invoke method. As you can see, this method will look at the first parameter to determine which function should be called, and invoke the appropriate tuna chaincode method.

    func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
    // Retrieve the requested Smart Contract function and arguments
    function, args := APIstub.GetFunctionAndParameters()
    // Route to the appropriate handler function to interact with the ledger appropriately

    if function == "queryTuna" {
    return s.queryTuna(APIstub, args)
    } else if function == "initLedger" {
    return s.initLedger(APIstub)
    } else if function == "recordTuna" {
    return s.recordTuna(APIstub, args)
    } else if function == "queryAllTuna" {
    return s.queryAllTuna(APIstub)
    } else if function == "changeTunaHolder" {
    return s.changeTunaHolder(APIstub, args)
    }
    return shim.Error("Invalid Smart Contract function name.")
    }
    
Chaincode Methods - queryTuna
    The queryTuna method would be used by a fisherman, regulator, or restaurateur to view the record of one particular tuna. It takes one argument - the key for the tuna in question.

    func (s *SmartContract) queryTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting 1")
    }
    tunaAsBytes, _ := APIstub.GetState(args[0])
    if tunaAsBytes == nil {
    return shim.Error(“Could not locate tuna”)
    }
    return shim.Success(tunaAsBytes)
    }


Chaincode Methods - initLedger

The initLedger method will add test data to our network.

    func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
    tuna := []Tuna{
    Tuna{Vessel: "923F", Location: "67.0006, -70.5476", Timestamp: "1504054225", Holder: "Miriam"},
    Tuna{Vessel: "M83T", Location: "91.2395, -49.4594", Timestamp: "1504057825", Holder: "Dave"},
    Tuna{Vessel: "T012", Location: "58.0148, 59.01391", Timestamp: "1493517025", Holder: "Igor"},
    Tuna{Vessel: "P490", Location: "-45.0945, 0.7949", Timestamp: "1496105425", Holder: "Amalea"},
    Tuna{Vessel: "S439", Location: "-107.6043, 19.5003", Timestamp: "1493512301", Holder: "Rafa"},
    Tuna{Vessel: "J205", Location: "-155.2304, -15.8723", Timestamp: "1494117101", Holder: "Shen"},
    Tuna{Vessel: "S22L", Location: "103.8842, 22.1277", Timestamp: "1496104301", Holder: "Leila"},
    Tuna{Vessel: "EI89", Location: "-132.3207, -34.0983", Timestamp: "1485066691", Holder: "Yuan"},
    Tuna{Vessel: "129R", Location: "153.0054, 12.6429", Timestamp: "1485153091", Holder: "Carlo"},
    Tuna{Vessel: "49W4", Location: "51.9435, 8.2735", Timestamp: "1487745091", Holder: "Fatima"},
    }
    i := 0
    for i < len(tuna) {
    fmt.Println("i is ", i)
    tunaAsBytes, _ := json.Marshal(tuna[i])
    APIstub.PutState(strconv.Itoa(i+1), tunaAsBytes)
    fmt.Println("Added", tuna[i])
    i = i + 1
    }
    return shim.Success(nil)
    }


Chaincode Methods - recordTuna

The recordTuna method is the method a fisherman like Sarah would use to record each of her tuna catches. This method takes in five arguments (attributes to be saved in the ledger).

    func (s *SmartContract) recordTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    if len(args) != 5 {
    return shim.Error("Incorrect number of arguments. Expecting 5")
    }
    var tuna = Tuna{ Vessel: args[1], Location: args[2], Timestamp: args[3], Holder: args[4]}
    tunaAsBytes, _ := json.Marshal(tuna)
    err := APIstub.PutState(args[0], tunaAsBytes)
    if err != nil {
    return shim.Error(fmt.Sprintf("Failed to record tuna catch: %s", args[0]))
    }
    return shim.Success(nil)
    }


Chaincode Methods - queryAllTuna

The queryAllTuna method allows for assessing all the records; in this case, all the Tuna records added to the ledger. This method does not take any arguments. It will return a JSON string containing the results.

    func (s *SmartContract) queryAllTuna(APIstub shim.ChaincodeStubInterface) sc.Response {
    startKey := "0"
    endKey := "999"
    resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
    if err != nil {
    return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    // buffer is a JSON array containing QueryResults
    var buffer bytes.Buffer
    buffer.WriteString("[")
    bArrayMemberAlreadyWritten := false

    for resultsIterator.HasNext() {
    queryResponse, err := resultsIterator.Next()
    if err != nil {
    return shim.Error(err.Error())
    }

    // Add a comma before array members, suppress it for the first array member
    if bArrayMemberAlreadyWritten == true {
    buffer.WriteString(",")
    }
    buffer.WriteString("{\"Key\":")
    buffer.WriteString("\"")
    buffer.WriteString(queryResponse.Key)
    buffer.WriteString("\"")
    buffer.WriteString(", \"Record\":")
    // Record is a JSON object, so we write as-is
    buffer.WriteString(string(queryResponse.Value))
    buffer.WriteString("}")
    bArrayMemberAlreadyWritten = true
    }
    buffer.WriteString("]")
    fmt.Printf("- queryAllTuna:\n%s\n", buffer.String())
    return shim.Success(buffer.Bytes())
    }


Chaincode Methods - changeTunaHolder

As the tuna fish is passed to different parties in the supply chain, the data in the world state can be updated with who has possession. The changeTunaHolder method takes in 2 arguments, tuna id and new holder name.

    func (s *SmartContract) changeTunaHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    if len(args) != 2 {
    return shim.Error("Incorrect number of arguments. Expecting 2")
    }
    tunaAsBytes, _ := APIstub.GetState(args[0])
    if tunaAsBytes != nil {
    return shim.Error("Could not locate tuna")
    }
    tuna := Tuna{}
    json.Unmarshal(tunaAsBytes, &tuna)
    // Normally check that the specified argument is a valid holder of tuna but here we are skipping this check for this example. 
    tuna.Holder = args[1]
    tunaAsBytes, _ = json.Marshal(tuna)
    err := APIstub.PutState(args[0], tunaAsBytes)
    if err != nil {
    return shim.Error(fmt.Sprintf("Failed to change tuna holder: %s", args[0]))
    }
    return shim.Success(nil)
    }
    
Conclusion
We hope you now have a better idea of how chaincode is constructed and written, especially when applied to a simple example. To see all the code snippets, visit the educational GitHub repository: https://github.com/hyperledger/education/blob/master/LFS171x/fabric-material/chaincode/tuna-app/tuna-chaincode.go.


    

