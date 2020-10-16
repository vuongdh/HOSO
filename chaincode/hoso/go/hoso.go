/*
 * ~/HOSO/fabric-samples/fabric-samples/chaincode/hoso/go
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

// Define the HOSO structure, with 22 properties.  Structure tags are used by encoding/json library
type HOSO struct {
	Malk   string `json:"malk"`
	Mabn   string `json:"mabn"`
	Maba   string `json:"maba"`
	Makb   string `json:"makb"`
	Hoten  string `json:"hoten"`
	Ngaysinh string `json:"ngaysinh"`
	Gioitinh  string `json:"gioitinh"`
	Cmnd  string `json:"cmnd"`
	Mathe  string `json:"mathe"`
	Diachi  string `json:"diachi"`
	Maicd  string `json:"maicd"`
	Chandoan  string `json:"chandoan"`
	Ghichu  string `json:"ghichu"`
	Lydo  string `json:"lydo"`
	Mabs  string `json:"mabs"`
	Tenbs  string `json:"tenbs"`
	Mach  string `json:"mach"`
	Nhietdo  string `json:"nhietdo"`
	Huyetap  string `json:"huyetap"`
	Cannang  string `json:"cannang"`
	Thuoc  string `json:"thuoc"`
	Canlamsang  string `json:"canlamsang"`
}

/*
 * The Init method is called when the Smart Contract "hoso" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "hoso"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryHOSO" {
		return s.queryHOSO(APIstub, args)

	} else if function == "initLedger" {
		return s.initLedger(APIstub)

	} else if function == "createHOSO" {
		return s.createHOSO(APIstub, args)

	} else if function == "queryAllHOSO" {
		return s.queryAllHOSO(APIstub)

	} else if function == "changeHOSO" {
		return s.changeHOSO(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryHOSO(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	hosoAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(hosoAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	hosos := []HOSO{
		HOSO{Malk: "92004_2020083125_2008040073",Mabn: "2020083125",Maba: "",Makb: "2008040073",Hoten: "Bùi Thị Tước",Ngaysinh: "1965-01-01",Gioitinh: "Nữ ", Cmnd: "", Mathe: "",Diachi: "Huyện Tam Bình, Vĩnh Long", Maicd: "K29",Chandoan: "Viêm dạ dày và tá tràng H.P(+)",Ghichu: "TT tinidazol 500mg 1vx2,Tetracylin 500mg 2vx2,Trymo 120mg 2vx2 trước ăn",Lydo: "",Mabs: "441",Tenbs: "BS.Lý Quốc Ánh",Mach: "",Nhietdo: "",Huyetap: "/",Cannang: "",Thuoc: "[{"tenhh":"Cimeverin","dvt":"viên","soluong":30.00,"cachuong":"Ngày uống 02 lần, mỗi lần 01 viên trước ăn","maicd":"K29","chandoan":"Viêm dạ dày và tá tràng H.P(+)"},{"tenhh":"KING KORE GOLD","dvt":"viên","soluong":15.00,"cachuong":"Sáng uống 1 viên","maicd":"K29","chandoan":"Viêm dạ dày và tá tràng H.P(+)"},{"tenhh":"MEDI-Levosulpirid","dvt":"viên","soluong":30.00,"cachuong":"Sáng uống 1 viên, chiều uống 1 viên","maicd":"K29","chandoan":"Viêm dạ dày và tá tràng H.P(+)"},{"tenhh":"RABARIS","dvt":"viên","soluong":30.00,"cachuong":"Sáng uống 1 viên, chiều uống 1 viên trước ăn","maicd":"K29","chandoan":"Viêm dạ dày và tá tràng H.P(+)"},{"tenhh":"VONSEnight","dvt":"viên","soluong":15.00,"cachuong":"tối uống 1 viên","maicd":"K29","chandoan":"Viêm dạ dày và tá tràng H.P(+)"}]", Canlamsang: "[{"tencls":"Xét nghiệm vi khuẩn HPYLORY qua hơi thở","dvt":"lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"K29","chandoan":"Viêm dạ dày và tá tràng "},{"tencls":"Siêu âm ổ bụng","dvt":"Lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"K29","chandoan":"Viêm dạ dày và tá tràng "},{"tencls":"Khám bệnh hạng I (Thu phí)","dvt":"Lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"K29","chandoan":"Viêm dạ dày và tá tràng H.P(+)"}]"},
		HOSO{Malk: "92004_2020083124_2008040080",Mabn: "2020083124", Maba: "",Makb: "2008040080", Hoten: "Nguyễn Hữu Tín", Ngaysinh: "2001-01-01",Gioitinh: "Nam",Cmnd: "", Mathe: "",Diachi: "Huyện Tam Bình, Vĩnh Long",Maicd: "L23",Chandoan: "Viêm da tiếp xúc dị ứng ",Ghichu: "",Lydo: "",Mabs: "441", Tenbs: "BS.Lý Quốc Ánh",Mach: "",Nhietdo: "",Huyetap: "/", Cannang: "", Thuoc: "[{"tenhh":"AMNONIMS","dvt":"viên","soluong":6.00,"cachuong":"Sáng uống 1 viên, chiều uống 1 viên","maicd":"L23","chandoan":"Viêm da tiếp xúc dị ứng "},{"tenhh":"KIDSOLON 4","dvt":"viên","soluong":6.00,"cachuong":"Ngày uống 02 lần, mỗi lần 02 viên","maicd":"L23","chandoan":"Viêm da tiếp xúc dị ứng "},{"tenhh":"RABARIS","dvt":"viên","soluong":3.00,"cachuong":"Sáng uống 1 viên","maicd":"L23","chandoan":"Viêm da tiếp xúc dị ứng "}]",Canlamsang : "[{"tencls":"Khám bệnh hạng I (Thu phí)","dvt":"Lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"L23","chandoan":"Viêm da tiếp xúc dị ứng "},{"tencls":"Toxocara (Giun đũa chó, mèo) Ab miễn dịch tự động","dvt":"Lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"L23","chandoan":"Viêm da tiếp xúc dị ứng "}]"},
		HOSO{Malk: "92004_2020083129_2008040095", Mabn: "2020083129", Maba: "", Makb: "2008040095",Hoten: "Đinh Kim Em", Ngaysinh: "1951-01-01",Gioitinh: "Nữ ",Cmnd: "",Mathe: "",Diachi: "Huyện Bình Minh, Vĩnh Long", Maicd: "I25.2", Chandoan: "Nhồi máu cơ tim cũ Bệnh Tăng huyết áp vô căn (nguyên phát); Tăng lipid máu hỗn hợp; Viêm dạ dày và tá tràng", Ghichu: "",Lydo: "",Mabs: "441",Tenbs: "BS.Lý Quốc Ánh",Mach: "",Nhietdo: "", Huyetap: "/", Cannang: "",Thuoc: "[{"tencls":"Khám bệnh hạng I (Thu phí)","dvt":"Lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"L23","chandoan":"Viêm da tiếp xúc dị ứng "},{"tencls":"Toxocara (Giun đũa chó, mèo) Ab miễn dịch tự động","dvt":"Lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"L23","chandoan":"Viêm da tiếp xúc dị ứng "}]", Canlamsang: "[{"tencls":"Điện tim thường","dvt":"Lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"I25.2","chandoan":"Nhồi máu cơ tim cũ Bệnh Tăng huyết áp vô căn (nguyên phát); Tăng lipid máu hỗn hợp; Viêm dạ dày và tá tràng"},{"tencls":"Khám bệnh hạng I (Thu phí)","dvt":"Lần","soluong":1.00,"mabs":"441","tenbs":"BS.Lý Quốc Ánh","maicd":"I25.2","chandoan":"Nhồi máu cơ tim cũ Bệnh Tăng huyết áp vô căn (nguyên phát); Tăng lipid máu hỗn hợp; Viêm dạ dày và tá tràng"}]"},

	}

	i := 0
	for i < len(hosos) {
		fmt.Println("i is ", i)
		hosoAsBytes, _ := json.Marshal(hosos[i])
		APIstub.PutState("HOSO"+strconv.Itoa(i), hosoAsBytes)
		fmt.Println("Added", hosos[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createHOSO(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	//if len(args) != 23 {
	//	return shim.Error("Incorrect number of arguments. Expecting 23")
	//}

	var hoso = HOSO{Malk: args[1],Mabn: args[2],Maba: args[3],Makb: args[4],Hoten: args[5],Ngaysinh: args[6],Gioitinh: args[7], Cmnd: args[8], Mathe: args[9],Diachi: args[10], Maicd: args[11],Chandoan: args[12],Ghichu: args[13],Lydo: args[14],Mabs: args[15],Tenbs: args[16],Mach: args[17],Nhietdo: args[18],Huyetap: args[19],Cannang: args[20], Thuoc: args[21], Canlamsang: args[22]}

	hosoAsBytes, _ := json.Marshal(hoso)
	APIstub.PutState(args[0], hosoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllHOSO(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "HOSO0"
	endKey := "HOSO999"

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

	fmt.Printf("- queryAllHOSO:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeHOSO(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	hosoAsBytes, _ := APIstub.GetState(args[0])
	hoso := HOSO{}

	json.Unmarshal(hosoAsBytes, &hoso)
	hoso.Hoten = args[1]

	hosoAsBytes, _ = json.Marshal(hoso)
	APIstub.PutState(args[0], hosoAsBytes)

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
