/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strings"
)

// Roaming Solution is a Smart Contract between telecom operator to exchange roaming call details records and settlement billing. 

type RoamingSolutionChaincode struct {
}

type WriteCallEventDetails struct {

		VPMN string `json:"vpmn"`
		HPMN string `json:"hpmn"`
		CallType string `json:"calltype"`
		SimChargeableSubsciber string `json:"simchargeablesubsciber"`
		CallEventStartTimeSatmap string `json:"calleventstarttimesatmap"`             
		TotalCallEventDuration string `json:"totalcalleventduration"`
		NetworkLocation string `json:"networklocation"`
		ImeiEquipmentIdentifier string `json:"imeiequipmentidentifier"`
		TeleServiceCode string `json:"teleservicecode"`
		ChargedItem string `json:"chargeditem"`
		ExchangeRateCode string `json:"exchangeratecode"`
		ChargeType string `json:"chargetype"`
		Charge string `json:"charge"`
		ChargeableUnits string `json:"chargeableunits"`
		ChargedUnits string `json:"chargedunits"`
		LocalTimeStamp string `json:"localtimestamp"`               
		Status string `json:"status"`
}

// Init method will be called during deployment.

func (t *RoamingSolutionChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	fmt.Println("Init Chaincode...done")

	return nil, nil
}


// Invoke function

func (t *RoamingSolutionChaincode) WriteCallEventDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("WriteCallEventDetails Invoke Begins...")
		
	if len(args) != 16 {
		return nil, errors.New("Incorrect number of arguments. Expecting 16")
	}
	status1 := "WriteCallEventDetailsCompleted"
	key := args[0]+args[15]
	WriteCallEventDetailsObj := WriteCallEventDetails{VPMN: args[0], HPMN: args[1], CallType: args[2], SimChargeableSubsciber: args[3], CallEventStartTimeSatmap: args[4], TotalCallEventDuration: args[5],NetworkLocation: args[6], ImeiEquipmentIdentifier: args[7], TeleServiceCode: args[8], ChargedItem: args[9], ExchangeRateCode: args[10], ChargeType: args[11], Charge: args[12],ChargeableUnits: args[13],ChargedUnits: args[14], LocalTimeStamp: args[15], Status: status1}
	err := stub.PutState(key,[]byte(fmt.Sprintf("%s",WriteCallEventDetailsObj)))
			if err != nil {
				return nil, err
			}
			
	fmt.Println("WriteCallEventDetails Invoke ends...")
	return nil, nil 
}



// in args this will take three values - VPMN and HPMN and LocalTimeStamp
func (t *RoamingSolutionChaincode) EntitlementFromHPMN(stub shim.ChaincodeStubInterface, argsVpmn []string) ([]byte, error) {
      
        if len(argsVpmn) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	    }

			
		key := argsVpmn[1]+argsVpmn[15]		
	        valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		} else if len(valAsbytes) == 0{
			jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		
		HPMNCallEventDetails := fmt.Sprintf("%s", valAsbytes)
		
		HPMNCallEventDetails = strings.Trim(HPMNCallEventDetails,"{")
		HPMNCallEventDetails = strings.Trim(HPMNCallEventDetails,"}")
		HPMNCallEventDetails = strings.Trim(HPMNCallEventDetails,"[")
		HPMNCallEventDetails = strings.Trim(HPMNCallEventDetails,"]")
		
		argsNew := strings.Split(HPMNCallEventDetails, " ")
		
		fmt.Println("HPMN Call Event Details Structure",argsNew)

		//Some HPMN Call Event Validation logic
						
		status1 := "HPMNApproved"
		CallEventDetailsFromHPMNObj := WriteCallEventDetails{VPMN: argsNew[0], HPMN: argsNew[1], CallType: argsNew[2], SimChargeableSubsciber: argsNew[3], CallEventStartTimeSatmap: argsNew[4], TotalCallEventDuration: argsNew[5],NetworkLocation: argsNew[6], ImeiEquipmentIdentifier: argsNew[7], TeleServiceCode: argsNew[8], ChargedItem: argsNew[9], ExchangeRateCode: argsNew[10], ChargeType: argsNew[11], Charge: argsNew[12],ChargeableUnits: argsNew[13],ChargedUnits: argsNew[14], LocalTimeStamp: argsNew[15], Status: status1}
        
		fmt.Println("VPMN+HPMN Call Event Details Structure",CallEventDetailsFromHPMNObj)
	
		err = stub.PutState(key,[]byte(fmt.Sprintf("%s",CallEventDetailsFromHPMNObj)))
			if err != nil {
				return nil, err
			}
			
			
		fmt.Println("Invoke EntitlementFromHPMN Chaincode... end") 
		return nil,nil		

}

// args should be three values - VPMN and LocalTimeStamp

func (t *RoamingSolutionChaincode) VPMNQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2 argument")
    }

    key = args[0]+args[15]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query VPMN Call Event Written confirmation ... end") 
    return valAsbytes, nil 

}

// args should be three values - HPMN and LocalTimeStamp

func (t *RoamingSolutionChaincode) EntitlementFromHPMNQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2 argument")
    }

    key = args[1]+args[15]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query EntitlementFromHPMNQuery ... end") 
    return valAsbytes, nil 

}


// Invoke Function

func (t *RoamingSolutionChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke Roaming Solution Chaincode... start") 
	
	// Handle different functions
	if function == "WriteCallEventDetails" {
		return t.WriteCallEventDetails (stub, args)
	}else if function == "EntitlementFromHPMN" {
		return t.EntitlementFromHPMN(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'WriteCallEventDetails' or 'EntitlementFromHPMN' but found '" + function + "'")
	}
	
	
	fmt.Println("Invoke Roaming Solution Chaincode... end") 
	
	return nil,nil;
}

// Query to get HPMN Call Event Details

func (t *RoamingSolutionChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query Roaming Solution Chaincode... start") 

	
	if function == "EntitlementFromHPMNQuery" {
		return t.EntitlementFromHPMNQuery(stub, args)
        }else if function == "VPMNQuery" {
		return t.VPMNQuery(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'EntitlementFromHPMNQuery' or 'VPMNQuery' but found '" + function + "'")
	} 
	
	// else we can query WorldState to fetch value
	
	var key, jsonResp string
    var err error

    if len(args) < 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }
	fmt.Println(len(args))
	if len(args) == 2 {
	   key = args[1]+args[15]
	} else {
	   key = args[0]
	}
    
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query EntitlementFromHPMN Chaincode... end") 
    return valAsbytes, nil 
  
	
}

func main() {
	err := shim.Start(new(RoamingSolutionChaincode))
	if err != nil {
		fmt.Println("Error starting RoamingSolutionChaincode: %s", err)
	}
}
