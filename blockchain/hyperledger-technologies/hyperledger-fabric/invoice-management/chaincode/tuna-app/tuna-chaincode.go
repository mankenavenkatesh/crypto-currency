// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Invoice structure
Structure tags are used by encoding/json library
*/
type Invoice struct {
	Seller string `json:"seller"`
    Buyer  string `json:"buyer"`
	Timestamp string `json:"timestamp"`
	Taxableamount string `json:"taxableamount"`
	Totaltax  string `json:"totaltax"`
    Invoicetotal  string `json:"invoicetotal"`
    Itemdesc  string `json:"itemdesc"`
    Status string `json:"status"`
}

/*
 * The Init method *
 called when the Smart Contract "invoice-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "tuna-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryInvoice" {
		return s.queryInvoice(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordInvoice" {
		return s.recordInvoice(APIstub, args)
	} else if function == "queryAllInvoice" {
		return s.queryAllInvoice(APIstub)
	} else if function == "changeInvoiceStatus" {
		return s.changeInvoiceStatus(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryInvoice method *
Used to view the records of one particular Invoice
It takes one argument -- the key for the tuna in question
 */
func (s *SmartContract) queryInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	invoiceAsBytes, _ := APIstub.GetState(args[0])
	if invoiceAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(invoiceAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 tuna catches)to our network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	invoice := []Invoice{
		Invoice{Seller: "Venkatesh",  Buyer: "Sharath", Timestamp: "20170630123546", Taxableamount:"4700rs", Totaltax: "270rs", Invoicetotal: "4970rs", Itemdesc:"Dry fruits pack" , Status:"Completed"},
		Invoice{Seller: "Divakar",  Buyer: "Shashank", Timestamp: "20160630123546", Taxableamount:"500rs", Totaltax: "30rs", Invoicetotal: "530rs", Itemdesc:"Vegetables",  Status:"Completed" },
	}
    
	i := 0
	for i < len(invoice) {
		fmt.Println("i is ", i)
		invoiceAsBytes, _ := json.Marshal(invoice[i])
		APIstub.PutState(strconv.Itoa(i+1), invoiceAsBytes)
		fmt.Println("Added", invoice[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordTuna method *
Sellers like Divakar would use to record each of her invoices generated.
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) recordInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}    
    
	var invoice = Invoice{Seller: args[1],  Buyer: args[2], Timestamp: args[3], Taxableamount:args[4], Totaltax: args[5], Invoicetotal: args[6], Itemdesc:args[7],Status:args[8] }

	invoiceAsBytes, _ := json.Marshal(invoice)
	err := APIstub.PutState(args[0], invoiceAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record invoice : %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllTuna method *
allows for assessing all the records added to the ledger(all tuna catches)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) queryAllInvoice(APIstub shim.ChaincodeStubInterface) sc.Response {

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
		// Add comma before array members,suppress it for the first array member
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

/*
 * The changeInvoiceStatus method *
The data in the world state can be updated with who has possession. 
This function takes in 2 arguments, Invoice id and new Invoice status
 */
func (s *SmartContract) changeInvoiceStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	invoiceAsBytes, _ := APIstub.GetState(args[0])
	if invoiceAsBytes == nil {
		return shim.Error("Could not locate invoice")
	}
	invoice := Invoice{}

	json.Unmarshal(invoiceAsBytes, &invoice)
	// Normally check that the specified argument is a valid status of invoice
	// we are skipping this check for this example
	invoice.Status = args[1]

	invoiceAsBytes, _ = json.Marshal(invoice)
	err := APIstub.PutState(args[0], invoiceAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to invoice Status holder: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}