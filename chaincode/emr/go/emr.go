/*
 * ~/EMR/fabric-samples/fabric-samples/chaincode/emr/go
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
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

// Define the EMR structure, with 6 properties.  Structure tags are used by encoding/json library
type EMR struct {
	Mabn   string `json:"mabn"`
	Hoten  string `json:"hoten"`
	Ngaysinh string `json:"ngaysinh"`
	Gioitinh  string `json:"gioitinh"`
	Cmnd  string `json:"cmnd"`
	Diachi  string `json:"diachi"`
}

/*
 * The Init method is called when the Smart Contract "emr" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "emr"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryEMR" {
		return s.queryEMR(APIstub, args)

	} else if function == "initLedger" {
		return s.initLedger(APIstub)

	} else if function == "createEMR" {
		return s.createEMR(APIstub, args)

	} else if function == "queryAllEMR" {
		return s.queryAllEMR(APIstub)

	} else if function == "changeEMR" {
		return s.changeEMR(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryEMR(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	emrAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(emrAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	emrs := []EMR{
		EMR{Mabn: "2020083198",Hoten: "Nguyễn Thị Mỹ Hương",Ngaysinh: "01/01/1969",Gioitinh: "Nữ",Cmnd: "",Diachi: "Phường An Thới, Quận Bình Thuỷ, Cần Thơ"},
		EMR{Mabn: "2020083191",Hoten: "Nguyễn Thị Sáu",Ngaysinh: "01/01/1939",Gioitinh: "Nữ", Cmnd: "", Diachi: "Phú Thạnh, Xã Long Phú, Huyện Tam Bình, Tỉnh Vĩnh Long"},
		EMR{Mabn: "2020083154",Hoten: "Đặng Thị Tám",Ngaysinh: "01/01/1951",Gioitinh: "Nữ", Cmnd: "", Diachi: "ấp thới hòa c, Xã Thới Xuân, Huyện Cờ Đỏ, Thành phố Cần Thơ"},

	}

	i := 0
	for i < len(emrs) {
		fmt.Println("i is ", i)
		emrAsBytes, _ := json.Marshal(emrs[i])
		APIstub.PutState("EMR"+strconv.Itoa(i), emrAsBytes)
		fmt.Println("Added", emrs[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createEMR(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	var emr = EMR{Mabn: args[1],Hoten: args[2],Ngaysinh: args[3],Gioitinh: args[4], Cmnd: args[5], Diachi: args[6]}

	emrAsBytes, _ := json.Marshal(emr)
	APIstub.PutState(args[0], emrAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllEMR(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "EMR0"
	endKey := "EMR999"

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

	fmt.Printf("- queryAllEMR:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeEMR(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	emrAsBytes, _ := APIstub.GetState(args[0])
	emr := EMR{}

	json.Unmarshal(emrAsBytes, &emr)
	emr.Hoten = args[1]

	emrAsBytes, _ = json.Marshal(emr)
	APIstub.PutState(args[0], emrAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
