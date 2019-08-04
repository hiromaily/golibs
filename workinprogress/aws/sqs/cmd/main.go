package main

//TODO:work in progress
import (
	"flag"
	"fmt"
	"os"
	"runtime"

	conf "github.com/hiromaily/golibs/config"
	lg "github.com/hiromaily/golibs/log"
	"github.com/hiromaily/golibs/workinprogress/aws/sqs"
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
	/*
		conf := conf.GetConf()
		//connection info
		conInfo := conf.CreateConInfo(conf.Request.Mid)

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
	*/
	return ""
}

// Set SQS on AWS
func setSQSData(num int, msg string) {
	conf := conf.GetConf()
	//conf.Aws.Sqs.QueueName -> test_message

	//1.オブジェクト作成
	sqs.New()
	//2.sqsにqueueがあるかチェック
	//3.なければ作成
	inputParams := sqs.CreateInputParam(conf.Aws.Sqs.QueueName)
	sendMsgRes, err := sqs.CreateNewQueue(inputParams)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("sendMsgRes : %v \n", sendMsgRes)

	//4.deadMessage用も同じく、
	//5.なければ作成
	inputParamsForDead := sqs.CreateInputParam(conf.Aws.Sqs.DeadQueueName)
	deadMsgRes, err := sqs.CreateNewQueue(inputParamsForDead)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("deadMsgRes : %v \n", deadMsgRes)

	//6.メッセージを作成
	//body := getContentBody("send message dayo")
	body := getContentBody(msg)
	acid := ""
	if num == 1 {
		sendInputParams := sqs.CreateSendMessageInput(sendMsgRes.QueueUrl, &body, &acid, *ot, *ct)

		//7.メッセージを送信
		sqs.SendMessageToQueue(sendInputParams)
	} else {
		var bulkCount = num / 10
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
	sendInputBatchParams := sqs.CreateSendMessageBatchInput(url, body, acid, *ot, *ct, num)

	//メッセージを送信
	sqs.SendMultipleMessagesToQueue(sendInputBatchParams)

}

func purgeSQSData() {
	sqs.New()
	conf := conf.GetConf()

	inputParams := sqs.CreateInputParam(conf.Aws.Sqs.QueueName)
	sendMsgRes, err := sqs.CreateNewQueue(inputParams)
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(*sendMsgRes.QueueUrl)
	sqs.PurgeQueue(sendMsgRes.QueueUrl)
}

func getSQSAttributes() {
	sqs.New()
	conf := conf.GetConf()

	inputParams := sqs.CreateInputParam(conf.Aws.Sqs.QueueName)
	sendMsgRes, err := sqs.CreateNewQueue(inputParams)
	if err != nil {
		panic(err.Error())
	}

	//check attribute
	params := sqs.CreateAttributesParams(sendMsgRes.QueueUrl)
	resp, err := sqs.GetQueueAttributes(params)
	if err != nil {
		panic(err.Error())
	}

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
	lg.Debug("flag.NArg(): ", flag.NArg())
	lg.Debug("mode: ", *mode)
	lg.Debug("toml: ", *toml)
	lg.Debug("n: ", *n)
	lg.Debug("msg: ", *msg)
	lg.Debug("loc: ", *loc)
	lg.Debug("ot: ", *ot)
	lg.Debug("ct: ", *ct)
	lg.Debug("debug: ", *debug)
}

//init
func init() {
	lg.InitializeLog(lg.DebugStatus, lg.NoDateNoFile, "[SQL]", "", "hiromaily")

	//handle command line
	handleCmdline()
}

func main() {
	//Timer Start
	//t := utils.TimeInfo{}
	//t.StartTimer()

	//get toml path
	if *toml != "" {
		conf.SetTOMLPath(*toml)
	}

	//-loc オプション -> 0:何も設定しない、1:位置情報をtrueにする、2:位置情報をfalseにする
	//get config
	//conf := config.GetConf()
	//パラメータに応じて、初期値を設定
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
