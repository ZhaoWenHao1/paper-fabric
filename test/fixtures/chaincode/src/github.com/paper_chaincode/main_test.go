/*
 * @Author: your name
 * @Date: 2021-06-11 08:34:02
 * @LastEditTime: 2021-06-11 10:47:09
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /undefined/Users/apple/go/src/github.com/chaincode/mychain/main_test.go
 */
package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func mockInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func addConference(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addConference"), []byte(args[0]), []byte(args[1]),
		[]byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7]),
		[]byte(args[8]), []byte(args[9]), []byte(args[10]), []byte(args[11]), []byte(args[12]), []byte(args[13]),
		[]byte(args[14]), []byte(args[15]), []byte(args[16]), []byte(args[17]), []byte(args[18]), []byte(args[19]),
		[]byte(args[20])})

	if res.Status != shim.OK {
		fmt.Println("addConference failed:", string(res.Message))
		t.FailNow()
	}
}
func addJournal(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addJournal"), []byte(args[0]), []byte(args[1]),
		[]byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7]),
		[]byte(args[8]), []byte(args[9]), []byte(args[10]), []byte(args[11]), []byte(args[12]), []byte(args[13]),
		[]byte(args[14]), []byte(args[15]), []byte(args[16]), []byte(args[17]), []byte(args[18]), []byte(args[19]),
		[]byte(args[20]), []byte(args[21]), []byte(args[22]), []byte(args[23]), []byte(args[24])})

	if res.Status != shim.OK {
		fmt.Println("addJournal failed:", string(res.Message))
		t.FailNow()
	}
}
func addSoftware(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addSoftware"), []byte(args[0]), []byte(args[1]),
		[]byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7]),
		[]byte(args[8]), []byte(args[9]), []byte(args[10]), []byte(args[11]), []byte(args[12]), []byte(args[13]),
		[]byte(args[14]), []byte(args[15]), []byte(args[16])})

	if res.Status != shim.OK {
		fmt.Println("addSoftware failed:", string(res.Message))
		t.FailNow()
	}
}
func addPatent(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addPatent"), []byte(args[0]), []byte(args[1]),
		[]byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7]),
		[]byte(args[8]), []byte(args[9]), []byte(args[10]), []byte(args[11]), []byte(args[12]), []byte(args[13]),
		[]byte(args[14]), []byte(args[15]), []byte(args[16]), []byte(args[17]), []byte(args[18]), []byte(args[19]),
		[]byte(args[20]), []byte(args[21]), []byte(args[22]), []byte(args[23]), []byte(args[24])})

	if res.Status != shim.OK {
		fmt.Println("addPatent failed:", string(res.Message))
		t.FailNow()
	}
}
func traceBackwardPaper(t *testing.T, stub *shim.MockStub, args []string) {

}
func traceBackwardConOrPatent(t *testing.T, stub *shim.MockStub, args []string) {

}

func TestAddConference(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	addConference(t, stub, []string{"author1", "ConferenceName", "projectNum11", "projectFund11",
		"createBy11", "createTime11", "updateBy11", "updateTime11", "1", "1",
		"0", "Operation","Signature", "PaperId", "Title", "HashTitle", "TimeRange", "Place", "PageNumRange",
		"ConferenceType", "0", "Reporter"}) //21
}

func TestAddJournal(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	addJournal(t, stub, []string{"author1p", "JournalName", "projectNum11p", "projectFund11p",
		"createBy11p", "createTime11p", "updateBy11p", "updateTime11p", "1", "1",
		"0", "Operation", "Signaturep", "PaperIdp", "Titlep", "HashTitlep", "Yearp", "VolumeNump",
		"IssueNump", "doip", "PageNumRange", "PublishTime", "Issn", "Type1", "Type2", "1"}) //25
}

func TestAddSoftware(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	addSoftware(t, stub, []string{"author1", "Remarks", "createBy11", "createTime11", "updateBy11",
		"updateTime11", "1", "0", "Operation", "PaperId", "Title", "HashTitle", "CertificateNum", "RegisterNum", "RegisterData",
		"Institution", "Project", "0"})
}

func TestAddPatent(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	addConference(t, stub, []string{"author1", "Remarks", "createBy11", "createTime11", "updateBy11",
		"updateTime11", "1", "0", "Operation", "PaperId", "Title", "HashTitle", "1", "ApplyNum", "ApplyData", "Institution",
		"Group", "Project", "Agent", "PatentNum", "PublishData", "1", "1", "RelativeId", "RefuseReason", "payRange"})
}

// func TestTraceBackwardPaper(t *testing.T) {
// 	smartContract := new(SmartContract)
// 	stub := shim.NewMockStub("SmartContract", smartContract)
// 	mockInit(t, stub, nil)
// 	addConference(t, stub, []string{"1", "department_software"})
// 	addConference(t, stub, []string{"2", "department_test"})
// }

// func TestTraceBackwardConOrPatent(t *testing.T) {
// 	smartContract := new(SmartContract)
// 	stub := shim.NewMockStub("SmartContract", smartContract)
// 	mockInit(t, stub, nil)
// 	addConference(t, stub, []string{"1", "department_software"})
// 	addConference(t, stub, []string{"2", "department_test"})
// }
