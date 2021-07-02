/*
 * @Author: your name
 * @Date: 2021-06-01 15:45:39
 * @LastEditTime: 2021-06-16 19:51:41
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /mychain/main.go
 */
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//论文溯源绑定结构体
type PaperTraceStruct struct {
	Hashtitle string `json:"paper_Info"` //论文的Hashtitle
	Txid      string `json:"tx_id"`      //论文绑定的事务号:Signature_Title_ConferenceName_Timestamp
}

//软著/专利溯源绑定结构体
type PatentOrSoftwareTraceStruct struct {
	HashTitle string `json:"hash_title"`
	Txid      string `json:"tx_id"` //绑定的事务号 Title_Timestamp
}

//论文结构11
type PaperCom struct {
	Authors        []string `json:"authors"`      //所有作者(中文名) 用于查询
	ConferenceName string   `json:"short_name"`   //会议/期刊名称缩写
	ProjectNum     string   `json:"project_num"`  //项目号
	ProjectFund    string   `json:"project_fund"` //项目基金
	CreateBy       string   `json:"create_by"`    //上传者id
	CreateTime     string   `json:"create_time"`  //上传日期
	UpdateBy       string   `json:"update_by"`    //上次更新人员id
	UpdateTime     string   `json:"update_time"`  //上次更新时间
	ModifyCount    int      `json:"modify_count"` //上链后修改次数
	IsTop80        int      `json:"is_top80"`     //是否为top80 0-否 1-是
	Exception      int      `json:"exception"`    //异常类型 0-无异常 1-题目重复 2-多次修改 3-题目重复并多次修改
	Operation      string   `json:"operation"`    //操作类型
}

//会议论文13
type Conference struct {
	PaperCommon    PaperCom `json:"paper_common"`        //论文公共字段信息
	Signature      string   `json:"signature"`           //数字签名
	PaperId        string   `json:"paper_id"`            //对应的paperId
	Title          string   `json:"paper_title"`         //论文题目
	HashTitle      string   `json:"hash_title"`          //论文题目hash值
	TimeRange      string   `json:"start_end_time"`      //会议开始-结束时间
	Place          string   `json:"place"`               //会议举办地点
	PageNumRange   string   `json:"page_num_starttoend"` //起始-结束页码
	ConferenceType string   `json:"type"`                //会议类别:CCF A B C等
	IsElectronic   int      `json:"is_electronic"`       //是否为电子版 0-否 1-是
	Reporter       string   `json:"reporter"`            //报告人姓名
	LastTxId       string   `json:"last_tx_id"`          //上次事务号
	ThisTxId       string   `json:"this_tx_id"`          //本次事务号
}

//期刊论文17
type Journal struct {
	PaperCommon  PaperCom `json:"paper_common"` //论文公共字段信息
	Signature    string   `json:"signature"`    //数字签名
	PaperId      string   `json:"paper_id"`     //对应的paperId
	Title        string   `json:"paper_title"`  //论文题目
	HashTitle    string   `json:"hash_title"`   //论文题目hash值
	Year         string   `json:"year"`         //期刊年份
	VolumeNum    string   `json:"volume_num"`   //期刊卷号
	IssueNum     string   `json:"issue_num"`    //期刊期号
	Doi          string   `json:"doi"`
	PageNumRange string   `json:"page_num_starttoend"` //起始-结束页码
	PublishTime  string   `json:"publish_time"`        //发表日期
	Issn         string   `json:"issn"`                //ISSN/ISBN:具体编号
	Type1        string   `json:"type1"`               //期刊类别：IEEE Trans ACM Trans等
	Type2        string   `json:"type2"`               //期刊类别：CCF A B C等
	FirstPublish int      `json:"first_publish"`       //是否为首发
	LastTxId     string   `json:"last_tx_id"`          //上次事务号
	ThisTxId     string   `json:"this_tx_id"`          //本次事务号
}

//软著和专利的公共信息8
type SoftAndPatentCom struct {
	Authors     []string `json:"authors"`
	Remarks     string   `json:"remarks"`
	CreateBy    string   `json:"create_by"`
	CreateTime  string   `json:"create_time"`
	UpdateBy    string   `json:"update_by"` //上次更新人员id
	UpdateTime  string   `json:"update_time"`
	ModifyCount int      `json:"modify_count"` //上链后修改次数
	Exception   int      `json:"exception"`    //异常类型 0-无异常 1-题目重复 2-多次修改
	Operation   string   `json:"operation"`    //操作类型

}

//软著//12
type Software struct {
	Common         SoftAndPatentCom `json:"SoftAndPatent_Com"` //公共信息
	Id             string           `json:"software_id"`
	Title          string           `json:"title"`              //版权名称
	HashTitle      string           `json:"hash_title"`         //软著titel的hash值
	CertificateNum string           `json:"certificate_number"` //证书号
	RegisterNum    string           `json:"register_num"`       //登记号
	RegisterData   string           `json:"register_date"`      //登记日期
	Institution    string           `json:"institution"`        //单位
	Project        string           `json:"project"`            //依托项目
	Status         int              `json:"status"`             //0-拟申请 1-已申请
	LastTxId       string           `json:"last_tx_id"`         //上次事务号
	ThisTxId       string           `json:"this_tx_id"`         //本次事务号
}

//专利20
type Patent struct {
	Common       SoftAndPatentCom `json:"SoftAndPatent_Com"` //公共信息
	Id           string           `json:"primary_id"`
	Title        string           `json:"title"`      //专利名称
	HashTitle    string           `json:"hash_title"` //专利titel的hash值
	Status       int              `json:"status"`     //0-拟定申请 1-申请未出版 2-已出版
	ApplyNum     string           `json:"apply_num"`  //专利申请号
	ApplyData    string           `json:"apply_date"`
	Institution  string           `json:"institution"`    //专利权人，单位
	Group        string           `json:"group"`          //组别
	Project      string           `json:"project"`        //依托项目
	Agent        string           `json:"agent"`          //代理单位
	PatentNum    string           `json:"patent_num"`     //专利号
	PublishData  string           `json:"publish_data"`   //授权公布日
	IsUSA        int              `json:"is_usa"`         //0-不是美国专利  1-是美国专利
	IsRefuse     int              `json:"is_refuse"`      //0-没有被驳回 1-被驳回
	RelativeId   string           `json:"relative_id"`    //isUSA为1时，对应国内专利id号
	RefuseReason string           `json:"refuse_reason"`  //isRefuse为1时，被驳回原因
	PayRange     string           `json:"pay_starttoend"` //缴维护费开始-结束时间
	LastTxId     string           `json:"last_tx_id"`     //上次事务号
	ThisTxId     string           `json:"this_tx_id"`     //本次事务号
}

// 定义智能合约结构体
type SmartContract struct {
}

// 在链码初始化过程中调用Init来初始化任何数据
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Paper_chaincode Init")
	return shim.Success(nil)
}

// 在链码每个事务中，Invoke会被调用。
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Paper_chaincode Invoke")

	function, args := stub.GetFunctionAndParameters()
	if function == "addConference" {
		return t.addConference(stub, args)
	} else if function == "addJournal" {
		return t.addJournal(stub, args)
	} else if function == "addSoftware" {
		return t.addSoftware(stub, args)
	} else if function == "addPatent" {
		return t.addPatent(stub, args)
	} else if function == "traceBackwardPaper" { //论文溯源
		return t.traceBackwardPaper(stub, args)
	} else if function == "traceBackwardConOrPatent" { //软著和专利溯源
		return t.traceBackwardConOrPatent(stub, args)
	} else if function == "traceBackwardPaperFromTxid" { //paperbyTxid
		return t.traceBackwardPaperFromTxid(stub, args)
	} else if function == "traceBackwardConOrPatentFromTxid" { //软著和专利溯源byTxid
		return t.traceBackwardConOrPatentFromTxid(stub, args)
	} else if function == "GetPaper" { //软著和专利溯源byTxid
		return t.GetPaper(stub, args)
	}

	fmt.Println("invoke did not find func: " + function)
	return shim.Error("Invalid Smart Contract function name.")
}

//对字符串str计算sha256
func sha256Str(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	sum := hash.Sum(nil)
	result := hex.EncodeToString(sum)
	return result
}

//添加会议论文，作者名字命名格式一律为authorxxx
func (t *SmartContract) addConference(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 21 {
		return shim.Error("Incorrect number of arguments. Expecting 13")
	}

	var num int //一共有几个作者
	num = 0
	var end int //作者结束的下标
	end = -1
	//生成作者数组
	var authors []string
	for n, str := range args {
		if len(str) > 6 {
			if str[0:6] == "author" {
				authors = append(authors, str[6:])
				end = n
				num = num + 1
			}
		}

	}

	var i int
	i = end + 1
	//生成论文公共结构体
	ModifyCountAsInt, err := strconv.Atoi(args[i+7])
	IsTop80AsInt, err := strconv.Atoi(args[i+8])
	ExceptionAsInt, err := strconv.Atoi(args[i+9])
	PaperCom := &PaperCom{authors, args[i], args[i+1], args[i+2], args[i+3], args[i+4], args[i+5],
		args[i+6], ModifyCountAsInt, IsTop80AsInt, ExceptionAsInt, args[i+10]}

	i = i + 11

	//生成会议论文结构体
	IsElectronicAsInt, err := strconv.Atoi(args[i+8])

	//获取{thistxid,lasttxid} args[i+3]是hashtitle
	AlltxId := getTxIdPaper(i, args, *PaperCom, stub)

	conference := &Conference{*PaperCom, args[i], args[i+1], args[i+2], args[i+3], args[i+4], args[i+5],
		args[i+6], args[i+7], IsElectronicAsInt, args[i+9], AlltxId[1], AlltxId[0]}
	conferenceAsBytes, err := json.Marshal(conference)
	if err != nil {
		return shim.Error(err.Error())
	}

	//更新溯源结构体上链
	//key是paperhashtitle
	hashtitleKey := "paper" + conference.HashTitle
	PaperTraceStructNew := &PaperTraceStruct{conference.HashTitle, AlltxId[0]} //hashtitle:thistxid
	PaperTraceStructNewAsBytes, err := json.Marshal(PaperTraceStructNew)
	err = stub.PutState(hashtitleKey, PaperTraceStructNewAsBytes) //上链的key是hashtitle
	//会议论文上链 key是thistxid
	err = stub.PutState(AlltxId[0], conferenceAsBytes)
	if err != nil {
		return shim.Error("cannot putstate conference")
	}

	return shim.Success(conferenceAsBytes)
}

//添加期刊论文
func (t *SmartContract) addJournal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 25 {
		return shim.Error("Incorrect number of arguments. Expecting 15")
	}

	var num int //一共有几个作者
	num = 0
	var end int //作者结束的下标
	end = -1
	//生成作者数组
	var authors []string
	for n, str := range args {
		if len(str) > 6 {
			if str[0:6] == "author" {
				authors = append(authors, str[6:])
				end = n
				num = num + 1
			}
		}

	}

	var i int
	i = end + 1
	//生成论文公共结构体
	ModifyCountAsInt, err := strconv.Atoi(args[i+7])
	IsTop80AsInt, err := strconv.Atoi(args[i+8])
	ExceptionAsInt, err := strconv.Atoi(args[i+9])
	PaperCom := &PaperCom{authors, args[i], args[i+1], args[i+2], args[i+3], args[i+4], args[i+5],
		args[i+6], ModifyCountAsInt, IsTop80AsInt, ExceptionAsInt, args[i+10]}

	i = i + 11
	//生成期刊论文结构体
	FirstPublishAsInt, err := strconv.Atoi(args[i+8])

	//获取{thistxid,lasttxid}
	AlltxId := getTxIdPaper(i, args, *PaperCom, stub)

	journal := &Journal{*PaperCom, args[i], args[i+1], args[i+2], args[i+3], args[i+4], args[i+5],
		args[i+6], args[i+7], args[i+8], args[i+9], args[i+10], args[i+11], args[i+12], FirstPublishAsInt, AlltxId[1], AlltxId[0]}

	journalAsBytes, err := json.Marshal(journal)
	if err != nil {
		return shim.Error(err.Error())
	}

	//更新溯源结构体上链
	//key是paperhashtitle
	hashtitleKey := "paper" + journal.HashTitle
	PaperTraceStructNew := &PaperTraceStruct{journal.HashTitle, AlltxId[0]} //hashtitle;thistxid
	PaperTraceStructNewAsBytes, err := json.Marshal(PaperTraceStructNew)
	err = stub.PutState(hashtitleKey, PaperTraceStructNewAsBytes) //上链的key是hashtitle
	//期刊论文上链 key是thistxid
	err = stub.PutState(AlltxId[0], journalAsBytes)

	if err != nil {
		return shim.Error("cannot putstate journal")
	}

	return shim.Success(journalAsBytes)
}

//论文：计算本次事务号，获得上次事务号{thistxid,lastid}
func getTxIdPaper(i int, args []string, com PaperCom, stub shim.ChaincodeStubInterface) [2]string {

	var AlltxId [2]string

	signature := args[i]
	title := args[i+2]
	hashtitle := args[i+3] //之前存储的溯源结构体存储的key是paperhashtitle
	hashtitleKey := "paper" + hashtitle
	ConferenceName := com.ConferenceName
	timestamp := strconv.FormatInt(time.Now().Unix(), 10) //获取string类型的时间戳
	//论文绑定的事务号:Signature_Title_ConferenceName_Timestamp
	txstr := signature + "_" + title + "_" + ConferenceName + "_" + timestamp
	thistxid := sha256Str(txstr) //计算本次事务号
	fmt.Printf("thistxid: %s", thistxid)
	//获取上次事务号
	var lastid string
	paperTraceStructAsbytes, err := stub.GetState(hashtitleKey) //之前存储的溯源结构体信息的value信息（字节数组）
	if err != nil {                                             //没有这个论文的之前事务号这个结构体上链过
		lastid = "head"
	} else if paperTraceStructAsbytes != nil {
		paperTraceStruct := PaperTraceStruct{}
		err = json.Unmarshal(paperTraceStructAsbytes, &paperTraceStruct)
		lastid = paperTraceStruct.Txid //上次事务号
	}

	AlltxId[0] = thistxid
	AlltxId[1] = lastid

	return AlltxId

}

//软著/专利：计算本次事务号，获得上次事务号{thistxid,lastid}
func getTxId(i int, args []string, stub shim.ChaincodeStubInterface) [2]string {

	var AlltxId [2]string

	title := args[i+1]
	timestamp := strconv.FormatInt(time.Now().Unix(), 10) //获取string类型的时间戳
	//论文绑定的事务号:Title_Timestamp
	txstr := title + "_" + timestamp
	thistxid := sha256Str(txstr) //计算本次事务号

	//获取上次事务号
	var lastid string
	hashtitle := args[i+2]                                    //之前存储的溯源结构体存储的key是hashtitle
	patentTraceStructAsbytes, err := stub.GetState(hashtitle) //之前存储的溯源结构体信息的value信息（字节数组）
	if err != nil {                                           //没有这个论文的之前事务号这个结构体上链过
		lastid = "head"
	} else if patentTraceStructAsbytes != nil {
		paperTraceStruct := PatentOrSoftwareTraceStruct{}
		err = json.Unmarshal(patentTraceStructAsbytes, &paperTraceStruct)
		lastid = paperTraceStruct.Txid //上次事务号
	}

	AlltxId[0] = thistxid
	AlltxId[1] = lastid

	return AlltxId

}

//添加软著
func (t *SmartContract) addSoftware(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 17 {
		return shim.Error("Incorrect number of arguments. Expecting 13")
	}

	var num int //一共有几个作者
	num = 0
	var end int //作者结束的下标
	end = -1
	//生成作者数组
	var authors []string
	for n, str := range args {
		if len(str) > 6 {
			if str[0:6] == "author" {
				authors = append(authors, str[6:])
				end = n
				num = num + 1
			}
		}

	}
	var i int
	i = end + 1
	//生成软著公共结构体
	ModifyCountAsInt, err := strconv.Atoi(args[i+5])
	ExceptionAsInt, err := strconv.Atoi(args[i+6])
	SoftwareCom := &SoftAndPatentCom{authors, args[i], args[i+1], args[i+2], args[i+3], args[i+4], ModifyCountAsInt, ExceptionAsInt, args[1+7]}

	i = i + 8
	//生成软著结构体
	StatusAsInt, err := strconv.Atoi(args[i+8])

	//获取{thistxid,lasttxid}
	AlltxId := getTxId(i, args, stub)

	software := &Software{*SoftwareCom, args[i], args[i+1], args[i+2], args[i+3], args[i+4], args[i+5],
		args[i+6], args[i+7], StatusAsInt, AlltxId[1], AlltxId[0]}

	softwareAsBytes, err := json.Marshal(software)
	if err != nil {
		return shim.Error(err.Error())
	}

	//更新溯源结构体上链
	TraceStructNew := &PatentOrSoftwareTraceStruct{args[i+2], AlltxId[0]} //hashtitle:thistxid
	TraceStructNewAsBytes, err := json.Marshal(TraceStructNew)
	err = stub.PutState(args[i+2], TraceStructNewAsBytes)
	//key是软著的thistxid 上链
	err = stub.PutState(AlltxId[0], softwareAsBytes)

	if err != nil {
		return shim.Error("cannot putstate software")
	}

	return shim.Success(softwareAsBytes)

}

//添加专利
func (t *SmartContract) addPatent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 25 {
		return shim.Error("Incorrect number of arguments. Expecting 21")
	}

	var num int //一共有几个作者
	num = 0
	var end int //作者结束的下标
	end = -1
	//生成作者数组
	var authors []string
	for n, str := range args {
		if len(str) > 6 {
			if str[0:6] == "author" {
				authors = append(authors, str[6:])
				end = n
				num = num + 1
			}
		}

	}
	var i int
	i = end + 1
	//生成专利公共结构体
	ModifyCountAsInt, err := strconv.Atoi(args[i+5])
	ExceptionAsInt, err := strconv.Atoi(args[i+6])
	PatentCom := &SoftAndPatentCom{authors, args[i], args[i+1], args[i+2], args[i+3], args[i+4], ModifyCountAsInt, ExceptionAsInt, args[1+7]}

	i = i + 8
	//生成专利结构体
	StatusAsInt, err := strconv.Atoi(args[i+3])
	IsUSAAsInt, err := strconv.Atoi(args[i+12])
	IsRefuseAsInt, err := strconv.Atoi(args[i+13])

	//获取{thistxid,lasttxid}
	AlltxId := getTxId(i, args, stub)

	patent := &Patent{*PatentCom, args[i], args[i+1], args[i+2], StatusAsInt, args[i+4], args[i+5],
		args[i+6], args[i+7], args[i+8], args[i+9], args[i+10], args[i+11], IsUSAAsInt, IsRefuseAsInt,
		args[i+14], args[i+15], args[i+16], AlltxId[1], AlltxId[0]}

	patentAsbytes, err := json.Marshal(patent)
	if err != nil {
		return shim.Error(err.Error())
	}

	//更新溯源结构体上链
	TraceStructNew := &PatentOrSoftwareTraceStruct{args[i+2], AlltxId[0]} //hashtitle:thistxid
	TraceStructNewAsBytes, err := json.Marshal(TraceStructNew)
	err = stub.PutState(args[i+2], TraceStructNewAsBytes)
	//key是专利的thistxid 上链
	err = stub.PutState(AlltxId[0], patentAsbytes)

	if err != nil {
		return shim.Error("cannot putstate patent")
	}

	return shim.Success(patentAsbytes)
}

//会议和期刊论文的溯源 args {hashtitle,type} type是种类，1是会议，2是期刊
func (t *SmartContract) traceBackwardPaper(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	//获取这个hashtitle绑定的txid
	var paperTraceStruct PaperTraceStruct
	hashtitleKey := "paper" + args[0]
	paperTraceStructAsBytes, err := stub.GetState(hashtitleKey)
	if err != nil { //如果找不到论文hashtitle和txid的绑定，那就是这个论文数据从未上链过
		return shim.Error(err.Error())
	} else if paperTraceStructAsBytes == nil { //目标数据不存在
		return shim.Error("Failed! data does not exist")
	}

	err = json.Unmarshal(paperTraceStructAsBytes, &paperTraceStruct)
	TxId := paperTraceStruct.Txid
	if TxId == "head" {
		fmt.Println("This is the first time")
	}

	//根据获取的txid作为key去查找论文上传结构体里面的信息
	if args[1] == "1" { //是会议论文的溯源
		var paper Conference
		// var lastpaper Conference
		paperAsBytes, err := stub.GetState(TxId)
		if err != nil {
			return shim.Error(err.Error())
		} else if paperAsBytes == nil { //目标数据不存在
			return shim.Error("Failed! data does not exist")
		}
		err = json.Unmarshal(paperAsBytes, &paper)
		// lastTxId := paper.LastTxId                       //通过本次事务号取得上次事务号
		// lastpaperAsBytes, err := stub.GetState(lastTxId) //通过上次事务号取得上一次论文结构体
		// err = json.Unmarshal(lastpaperAsBytes, &lastpaper)

		fmt.Printf("get Conference paper:%s", paper.Title)
		return shim.Success(paperAsBytes)

	} else if args[1] == "2" { //是期刊论文的溯源
		var paper Journal
		// var lastpaper Journal
		paperAsBytes, err := stub.GetState(TxId)
		if err != nil {
			return shim.Error(err.Error())
		} else if paperAsBytes == nil { //目标数据不存在
			return shim.Error("Failed! data does not exist")
		}
		err = json.Unmarshal(paperAsBytes, &paper)
		// lastTxId := paper.LastTxId                       //通过本次事务号取得上次事务号
		// lastpaperAsBytes, err := stub.GetState(lastTxId) //通过上次事务号取得上一次论文结构体
		// err = json.Unmarshal(lastpaperAsBytes, &paper)

		fmt.Printf("get Journal paper:%s\n", paper.Title)
		return shim.Success(paperAsBytes)
	}

	return shim.Success(nil)
}

//(通过事务号)会议和期刊论文的溯源 args {txid,type} type是种类，1是会议，2是期刊
func (t *SmartContract) traceBackwardPaperFromTxid(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	//根据获取的txid作为key去查找论文上传结构体里面的信息
	if args[1] == "1" { //是会议论文的溯源
		var paper Conference
		//var lastpaper Conference
		paperAsBytes, err := stub.GetState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		} else if paperAsBytes == nil { //目标数据不存在
			return shim.Error("Failed! data does not exist")
		}
		err = json.Unmarshal(paperAsBytes, &paper)
		// lastTxId := paper.LastTxId                       //通过本次事务号取得上次事务号
		// lastpaperAsBytes, err := stub.GetState(lastTxId) //通过上次事务号取得上一次论文结构体
		// err = json.Unmarshal(lastpaperAsBytes, &lastpaper)

		fmt.Printf("get Conference paper:%s", paper.Title)
		return shim.Success(paperAsBytes)

	} else if args[1] == "2" { //是期刊论文的溯源
		var paper Journal
		//var lastpaper Journal
		paperAsBytes, err := stub.GetState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		} else if paperAsBytes == nil { //目标数据不存在
			return shim.Error("Failed! data does not exist")
		}
		err = json.Unmarshal(paperAsBytes, &paper)
		// lastTxId := paper.LastTxId                       //通过本次事务号取得上次事务号
		// lastpaperAsBytes, err := stub.GetState(lastTxId) //通过上次事务号取得上一次论文结构体
		// err = json.Unmarshal(lastpaperAsBytes, &lastpaper)

		fmt.Printf("get Journal paper:%s", paper.Title)
		return shim.Success(paperAsBytes)
	}

	return shim.Success(nil)
}

//软著和专利的溯源 args {hashtitle,type} type是种类，1是软著，2是专利
func (t *SmartContract) traceBackwardConOrPatent(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	//获取这个hashtitle绑定的txid
	var TraceStruct PatentOrSoftwareTraceStruct
	TraceStructAsBytes, err := stub.GetState(args[0])
	if err != nil { //如果找不到论文hashtitle和txid的绑定，那就是这个论文数据从未上链过
		return shim.Error(err.Error())
	} else if TraceStructAsBytes == nil { //目标数据不存在
		return shim.Error("Failed! data does not exist")
	}

	err = json.Unmarshal(TraceStructAsBytes, &TraceStruct)
	TxId := TraceStruct.Txid
	if TxId == "head" {
		fmt.Println("This is the first time")
	}

	//根据获取的txid作为key去查找上传结构体里面的信息
	if args[1] == "1" { //是软著的溯源
		var paper Software
		//var lastpaper Software
		paperAsBytes, err := stub.GetState(TxId)
		if err != nil {
			return shim.Error(err.Error())
		} else if paperAsBytes == nil { //目标数据不存在
			return shim.Error("Failed! data does not exist")
		}
		err = json.Unmarshal(paperAsBytes, &paper)
		// lastTxId := paper.LastTxId                       //通过本次事务号取得上次事务号
		// lastpaperAsBytes, err := stub.GetState(lastTxId) //通过上次事务号取得上一次结构体
		// err = json.Unmarshal(lastpaperAsBytes, &lastpaper)

		fmt.Printf("get software name:%s", paper.Title)
		return shim.Success(paperAsBytes)

	} else if args[1] == "2" { //是专利的溯源
		var paper Patent
		//var lastpaper Patent
		paperAsBytes, err := stub.GetState(TxId)
		if err != nil {
			return shim.Error(err.Error())
		} else if paperAsBytes == nil { //目标数据不存在
			return shim.Error("Failed! data does not exist")
		}
		err = json.Unmarshal(paperAsBytes, &paper)
		// lastTxId := paper.LastTxId                       //通过本次事务号取得上次事务号
		// lastpaperAsBytes, err := stub.GetState(lastTxId) //通过上次事务号取得上一次结构体
		// err = json.Unmarshal(lastpaperAsBytes, &lastpaper)

		fmt.Printf("get Journal paper:%s", paper.Title)
		return shim.Success(paperAsBytes)
	}

	return shim.Success(nil)
}

//(通过事务号)软著和专利的溯源 args {txid,type} type是种类，1是软著，2是专利
func (t *SmartContract) traceBackwardConOrPatentFromTxid(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	//根据获取的txid作为key去查找上传结构体里面的信息
	if args[1] == "1" { //是软著的溯源
		var paper Software
		//var lastpaper Software
		paperAsBytes, err := stub.GetState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		} else if paperAsBytes == nil { //目标数据不存在
			return shim.Error("Failed! data does not exist")
		}
		err = json.Unmarshal(paperAsBytes, &paper)
		// lastTxId := paper.LastTxId                       //通过本次事务号取得上次事务号
		// lastpaperAsBytes, err := stub.GetState(lastTxId) //通过上次事务号取得上一次结构体
		// err = json.Unmarshal(lastpaperAsBytes, &lastpaper)

		fmt.Printf("get software name:%s", paper.Title)
		return shim.Success(paperAsBytes)

	} else if args[1] == "2" { //是专利的溯源
		var paper Patent
		//var lastpaper Patent
		paperAsBytes, err := stub.GetState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		} else if paperAsBytes == nil { //目标数据不存在
			return shim.Error("Failed! data does not exist")
		}
		err = json.Unmarshal(paperAsBytes, &paper)
		// lastTxId := paper.LastTxId                       //通过本次事务号取得上次事务号
		// lastpaperAsBytes, err := stub.GetState(lastTxId) //通过上次事务号取得上一次结构体
		// err = json.Unmarshal(lastpaperAsBytes, &lastpaper)

		fmt.Printf("get Journal paper:%s", paper.Title)
		return shim.Success(paperAsBytes)
	}

	return shim.Success(nil)
}

//查询会议论文
func (t *SmartContract) GetPaper(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	paperAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("cannot getstate paper")
	}
	if paperAsBytes == nil {
		return shim.Success([]byte("not found the paper"))
	}

	return shim.Success([]byte(paperAsBytes))
}

// //查询期刊论文
// func (t *SmartContract) GetJournal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	paperAsBytes, err := stub.GetState(args[0])
// 	if err != nil {
// 		return shim.Error("cannot getstate paper")
// 	}
// 	if paperAsBytes == nil {
// 		return shim.Success([]byte("not found the paper"))
// 	}

// 	return shim.Success([]byte(paperAsBytes))
// }

// //查询软著
// func (t *SmartContract) GetSoftware(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	softwareAsBytes, err := stub.GetState(args[0])
// 	if err != nil {
// 		return shim.Error("cannot getstate saftware")
// 	}
// 	if softwareAsBytes == nil {
// 		return shim.Success([]byte("not found the software"))
// 	}

// 	return shim.Success([]byte(softwareAsBytes))
// }

// //查询专利
// func (t *SmartContract) GetPatent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	patentAsBytes, err := stub.GetState(args[0])
// 	if err != nil {
// 		return shim.Error("cannot getstate patent")
// 	}
// 	if patentAsBytes == nil {
// 		return shim.Success([]byte("not found the patent"))
// 	}

// 	return shim.Success([]byte(patentAsBytes))
// }

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
