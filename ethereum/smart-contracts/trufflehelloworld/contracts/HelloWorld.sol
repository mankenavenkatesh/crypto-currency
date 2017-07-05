pragma solidity ^0.4.4;

contract HelloWorld {
    uint public balance;
    
    // constructor function
    // Runs once automatically when contract is deployed on ethereum blockchain
    function HelloWorld(){
        balance = 1000;
    }
    
    function deposit(uint _value) returns(uint _newValue){
        balance+=_value;
        return balance;        
    }
    
}