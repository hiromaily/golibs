package sqs

//TODO:work in progress
import (
	"flag"
	"fmt"
	"log"
	sqslib "oden.dac.co.jp/sallytools/common/libs/aws/sqs"
	"oden.dac.co.jp/sallytools/common/libs/config"
	logex "oden.dac.co.jp/sallytools/common/libs/log"
	//"oden.dac.co.jp/sallytools/common/libs/utils"
	"os"
	"runtime"
)

// WARNING
// この機能はunitestから呼び出し戻り値を取得し、利用しているため、むやみにfmtやlogパッケージから
// 出力しないようにお願いします。

//command line
var (
	mode  = flag.Int("mode", 1, "")
	toml  = flag.String("toml", "", "toml file path")
	n     = flag.Int("n", 1, "amount of sqs messages for setting on SQS")
	msg   = flag.String("msg", "send message dayo", "text data on sqs messages")
	loc   = flag.Int("loc", 0, "data for including location")
	ot    = flag.String("ot", "0", "operationType for sqs messages")
	ct    = flag.String("ct", "0", "operationType for sqs messages")
	debug = flag.Int("debug", 1, "")
)
var usage = `Usage: boom [options...] <url>

Options:
  -mode       1:create, 2:purge
  -toml       tomlファイルのパス
  -n          SQSに作成するデータ数
  -msg        sqsにセットするtextメッセージ
  -loc        0:何も設定しない、1:位置情報をtrueにする、2:位置情報をfalseにする
  -ot         sqsの属性値である operationTypeを設定する。0の場合、tomlより取得した値をセット
  -ct         sqsの属性値である contentTypeを設定する。  0の場合、tomlより取得した値をセット
  -debug      1:show something logs.

e.g. for how to use
  // mode1 is for specific sqs data
  $ ./createsqs -loc 1 -debug 1
`

//TODO:otの値によって処理を分岐せねばならない。
func getContentBody(text string) string {
	//get config
	conf := config.GetConfInstance()

	//connection info
	conInfo := config.CreateConInfo(conf.Request.Mid)

	//encode Json data
	if *ot != "2" {
		bytesMessage, err := conInfo.CreateSendMessage(text, 2)
		if err != nil {
			panic(err)
		}
		return string(bytesMessage)
	} else {
		bytesMessage, err := conInfo.CreateSendMessageForOpType2(text)
		if err != nil {
			panic(err)
		}
		return string(bytesMessage)
	}
}

// Set SQS on AWS
func setSQSData(num int, msg string) {
	conf := config.GetConfInstance()
	//conf.Aws.Sqs.QueueName -> test_message

	//1.オブジェクト作成
	sqslib.New()
	//2.sqsにqueueがあるかチェック
	//3.なければ作成
	inputParams := sqslib.CreateInputParam(conf.Aws.Sqs.QueueName)
	sendMsgRes, err := sqslib.CreateNewQueue(inputParams)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("sendMsgRes : %v \n", sendMsgRes)

	//4.deadMessage用も同じく、
	//5.なければ作成
	inputParamsForDead := sqslib.CreateInputParam(conf.Aws.Sqs.DeadQueueName)
	deadMsgRes, err := sqslib.CreateNewQueue(inputParamsForDead)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("deadMsgRes : %v \n", deadMsgRes)

	//6.メッセージを作成
	//body := getContentBody("send message dayo")
	body := getContentBody(msg)
	acid := config.GetConfInstance().Request.Acid
	if num == 1 {
		sendInputParams := sqslib.CreateSendMessageInput(sendMsgRes.QueueUrl, &body, &acid, *ot, *ct)

		//7.メッセージを送信
		sqslib.SendMessageToQueue(sendInputParams)
	} else {
		var bulkCount int = num / 10
		if bulkCount >= 1 {
			for i := 0; i < bulkCount; i++ {
				sendBulkProcedure(sendMsgRes.QueueUrl, &body, &acid, 10)
			}
			if num%10 != 0 {
				sendBulkProcedure(sendMsgRes.QueueUrl, &body, &acid, num%10)
			}
		} else {
			sendBulkProcedure(sendMsgRes.QueueUrl, &body, &acid, num)
		}
	}
}

func sendBulkProcedure(url *string, body *string, acid *string, num int) {
	//10通まで
	sendInputBatchParams := sqslib.CreateSendMessageBatchInput(url, body, acid, *ot, *ct, num)

	//メッセージを送信
	sqslib.SendMultipleMessagesToQueue(sendInputBatchParams)

}

func purgeSQSData() {
	sqslib.New()
	conf := config.GetConfInstance()

	inputParams := sqslib.CreateInputParam(conf.Aws.Sqs.QueueName)
	sendMsgRes, err := sqslib.CreateNewQueue(inputParams)
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(*sendMsgRes.QueueUrl)
	sqslib.PurgeQueue(sendMsgRes.QueueUrl)
}

func getSQSAttributes() {
	sqslib.New()
	conf := config.GetConfInstance()

	inputParams := sqslib.CreateInputParam(conf.Aws.Sqs.QueueName)
	sendMsgRes, err := sqslib.CreateNewQueue(inputParams)
	if err != nil {
		panic(err.Error())
	}

	//check attribute
	params := sqslib.CreateAttributesParams(sendMsgRes.QueueUrl)
	resp, err := sqslib.GetQueueAttributes(params)

	//fmt.Printf("%s", resp)
	//fmt.Printf("%v", resp.Attributes)
	//fmt.Printf("%s", *resp.Attributes["ApproximateNumberOfMessages"])

	//暫定対応
	fmt.Printf("%s,%s", *resp.Attributes["ApproximateNumberOfMessages"], *resp.Attributes["ApproximateNumberOfMessagesNotVisible"])

	//ApproximateNumberOfMessages: "3",

}

// handling command line
func handleCmdline() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, runtime.NumCPU()))
	}

	flag.Parse()

	//log.Debug("flag.NArg(): " + strconv.Itoa(flag.NArg()))
	logex.Debug("flag.NArg(): ", flag.NArg())
	logex.Debug("mode: ", *mode)
	logex.Debug("toml: ", *toml)
	logex.Debug("n: ", *n)
	logex.Debug("msg: ", *msg)
	logex.Debug("loc: ", *loc)
	logex.Debug("ot: ", *ot)
	logex.Debug("ct: ", *ct)
	logex.Debug("debug: ", *debug)
}

//init
func init() {
	logex.InitializeDefaultLog(2, log.Ltime) //From debug log

	//handle command line
	handleCmdline()
}

// main
func main() {
	//Timer Start
	//t := utils.TimeInfo{}
	//t.StartTimer()

	//get toml path
	if *toml != "" {
		config.SetTomlPath(*toml)
	}

	//-loc オプション -> 0:何も設定しない、1:位置情報をtrueにする、2:位置情報をfalseにする
	//get config
	//conf := config.GetConfInstance()
	if *loc == 1 {
		config.ChangeConfigLocationEnable(true)
	} else if *loc == 2 {
		config.ChangeConfigLocationEnable(false)
	}

	//パラメータに応じて、初期値を設定
	conf := config.GetConfInstance()
	if *ot == "0" {
		*ot = conf.Aws.Sqs.MsgAttr.OpType
	}
	if *ct == "0" {
		*ct = conf.Aws.Sqs.MsgAttr.OpType
	}

	if *mode == 1 {
		//set sqs data
		setSQSData(*n, *msg)
	} else if *mode == 2 {
		getSQSAttributes()
	} else if *mode == 3 {
		purgeSQSData()
	}

	//Timer End
	//t.StopTimer()
}
