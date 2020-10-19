package mongodb_test

import (
	"encoding/json"

	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/mongodb"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"

	//r "github.com/hiromaily/golibs/runtimes"
	"os"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//MongoDB Ver.3.x
//DONE
//TODO:読み込んだJSON(フォーマットがわからない)をとりあえず、insertするには？
//TODO:jsonの要素内の配列を増やすなど
//TODO:フォーマットがわからないJSONをmongoから読み込んだ場合は？(とりあえず、logをセットした場合とか)
//TODO:期限を付ける
//TODO:レコード件数の取得(n, err := coll.Count())
//TODO:Datetime (GMT)

//Not yet
//TODO:テーブル結合のロジックによるカバー

type Company struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	CompanyID int           `bson:"company_id"`
	Name      string        `bson:"name"`
}

type Work struct {
	Occupation string `bson:"occupation"`
	CompanyID  int    `bson:"company_id"`
}

type Address struct {
	ZipCode  string `bson:"zipcode"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
}

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Age       int           `bson:"age"`
	Address   Address       `bson:"address"`
	Works     []Work        `bson:"works"`
	CreatedAt time.Time     `bson:"createdAt"`
}

var (
	//Database Name For test
	testDbName = "testdb01"

	//Collection Name For test
	testColUser    = "user"
	testColCompany = "company"
	testColTeacher = "teacher"

	savedUserID string
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[MongoDB]")

	c := conf.GetConf().Mongo

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)
	if c.DbName != "" {
		//GetMongo().GetDB("hiromaily")
		GetMongo().GetDB(c.DbName)
	}
}

func teardown() {
	GetMongo().Close()
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// functions
//-----------------------------------------------------------------------------
func CreateCompanyData() error {
	mg := GetMongo()
	mg.GetCol(testColCompany)

	// test data
	companies := []Company{
		{CompanyID: 1, Name: "company01"},
		{CompanyID: 2, Name: "company02"},
		{CompanyID: 3, Name: "company03"},
		{CompanyID: 4, Name: "company04"},
		{CompanyID: 5, Name: "company05"},
		{CompanyID: 6, Name: "company06"},
		{CompanyID: 7, Name: "company07"},
		{CompanyID: 8, Name: "company08"},
		{CompanyID: 9, Name: "company09"},
		{CompanyID: 10, Name: "company10"},
	}

	bulk := mg.C.Bulk()
	for _, v := range companies {
		bulk.Insert(v)
	}
	//bulk.Insert(companies...)
	//->cannot use &companies (type *[]Company) as type []interface {} in argument to bulk.Insert
	_, err := bulk.Run()
	return err
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// Preparation
//-----------------------------------------------------------------------------
func TestCreateDatabase(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()

	//if database is not exsisting, create database
	mg.GetDB(testDbName)
}

/*
func TestCreateCollection(t *testing.T) {
	tu.SkipLog(t)
	mg := GetMongo()

	err := mg.CreateCol(testColUser)
	if err != nil {
		t.Errorf("mg.CreateCol(testColUser) / error: %s", err)
		//error: collection already exists
	}
}
*/

//set expire index
func TestSetExpireOnCollection(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	//var sessionExpire time.Duration = 60 * 1 //one minute
	var sessionExpire = 1 * time.Minute //one minute

	err := mg.SetExpireOnCollection(sessionExpire)

	if err != nil {
		t.Errorf("mg.C.EnsureIndex(sessionTTL) / error: %s", err)
	}
	//db.user.getIndexes()
}

//-----------------------------------------------------------------------------
// CREATE
//-----------------------------------------------------------------------------
// insert one record
func TestInsertOne(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	// test data
	works := []Work{{Occupation: "programmer", CompanyID: 1}, {Occupation: "programmer", CompanyID: 2}}
	address := Address{ZipCode: "1060047", Country: "Japan", City: "Tokyo", Address1: "港区南麻布2-9-7", Address2: "マンション101"}
	user := User{Name: "Harry", Age: 25, Address: address, Works: works, CreatedAt: time.Now()}

	// insert
	err := mg.C.Insert(&user)
	if err != nil {
		t.Errorf("mg.C.Insert(&user) / error:%s", err)
	}

	// check length
	cnt, err := mg.C.Count()
	lg.Debugf("count:%d, err = %s", cnt, err)

	//TODO: Local Time
	//time:2016-07-14 16:57:39.550187911 +0900 JST
	lg.Debugf("time:%v", user.CreatedAt.Local())
}

// insert multiple record at once
func TestBulkInsert(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	//#1 user data
	works1 := []Work{{Occupation: "lawyer", CompanyID: 3}, {Occupation: "lawyer", CompanyID: 4}}
	address1 := Address{ZipCode: "1530002", Country: "Japan", City: "Tokyo", Address1: "目黒区目黒1-2-3", Address2: "マンションXX"}
	user1 := User{Name: "Ren", Age: 25, Address: address1, Works: works1, CreatedAt: time.Now()}

	works2 := []Work{{Occupation: "programmer", CompanyID: 5}, {Occupation: "programmer", CompanyID: 6}}
	address2 := Address{ZipCode: "1230001", Country: "Japan", City: "Tokyo", Address1: "杉並区福田5-5-1", Address2: ""}
	user2 := User{Name: "Shana", Age: 22, Address: address2, Works: works2, CreatedAt: time.Now()}

	works3 := []Work{{Occupation: "programmer", CompanyID: 1}, {Occupation: "programmer", CompanyID: 5}}
	address3 := Address{ZipCode: "2590047", Country: "Japan", City: "Kanagawa", Address1: "逗子適当1-2-3", Address2: ""}
	user3 := User{Name: "Ken", Age: 31, Address: address3, Works: works3, CreatedAt: time.Now()}

	bulk := mg.C.Bulk()
	bulk.Insert(&user1)
	bulk.Insert(&user2, &user3)
	_, err := bulk.Run()
	if err != nil {
		t.Errorf("bulk.Insert(user), bulk.Run() / error:%s", err)
	}

	//#2 company
	mg.GetCol(testColCompany)
	err = CreateCompanyData()
	if err != nil {
		t.Errorf("CreateCompanyData(), bulk.Insert(company): error:%s", err)
	}
}

// insert from json file
func TestInsertJsonFile(t *testing.T) {
	//tu.SkipLog(t)

	//json
	fileData, err := LoadJSONFile(*tu.JSONFile)
	if err != nil {
		t.Fatal("Loading json file was failed.")
	}

	//var v []interface{}
	var v map[string]interface{}

	if err := json.Unmarshal(fileData, &v); err != nil {
		t.Fatal("Unmarshal json file was failed.")
	}

	//
	mg := GetMongo()
	mg.GetCol(testColTeacher)

	//err = mg.C.Insert(v...)
	err = mg.C.Insert(&v)
	if err != nil {
		t.Errorf("mg.C.Insert(&v) / error:%s", err)
	}
}

//-----------------------------------------------------------------------------
// READ
//-----------------------------------------------------------------------------
func TestGetOneDataByColumn(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	searchName := "Ken"

	user := new(User)

	//find
	err := mg.FindOne(bson.M{"name": searchName}, user)
	//mg.C.Find(bson.M{"name": searchName}).One(user)
	if err != nil {
		t.Errorf("mg.FindOne(bson.M{\"name\": searchName}, user) / error:%s", err)
		//error:not found
	}

	lg.Debugf("user_id is %v", user.ID)
	lg.Debugf("result user by find: %v", *user)

	//save
	savedUserID = GetObjectID(user.ID)
	lg.Debugf("savedUserID is %s", savedUserID)
}

func TestGetOneDataById(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	//This value is changeable as processing
	//userId := "5785f375340e7601939628b5"
	userID := savedUserID
	user := new(User)

	//find
	err := mg.C.Find(bson.M{"_id": bson.ObjectIdHex(userID)}).One(user)
	t.Logf("result user by find id: %v", *user)
	//mg.C.Find(bson.M{"name": searchName}).One(user)
	if err != nil {
		t.Errorf("mg.C.Find(bson.M{\"_id\": bson.ObjectIdHex(userId)}).One(user) / error:%s", err)
	}

	lg.Debugf("user_id is %v", user.ID)
	lg.Debugf("result user by find: %v", *user)
}

func TestGetAllDataByColumn(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	var users []User

	//#1 target is sinple element
	colQuerier := bson.M{"age": 25}
	err := mg.C.Find(colQuerier).All(&users)
	if err != nil {
		t.Errorf("mg.C.Find(colQuerier).All(&users) / error:%s", err)
	}
	if len(users) == 0 {
		t.Error("mg.C.Find(colQuerier).All(&users) / no data")
	}
	lg.Debugf("result users by find.all: length is %d,\n %+v", len(users), users)
	lg.Debug("- - - - - - - - - - - - - - - - - -")

	//#2 target is nested element
	users = nil
	colQuerier = bson.M{"address.zipcode": "1060047"}
	err = mg.C.Find(colQuerier).All(&users)
	if err != nil {
		t.Errorf("mg.C.Find(colQuerier).All(&users) / error:%s", err)
	}
	if len(users) == 0 {
		t.Error("mg.C.Find(colQuerier).All(&users) / no data")
	}
	lg.Debugf("result users by find.all: length is %d,\n %+v", len(users), users)
	lg.Debug("- - - - - - - - - - - - - - - - - -")

	//#3 target is nested and array element
	users = nil
	//bson.M{"categories": bson.M{"$elemMatch": bson.M{"slug": "general"}}}
	colQuerier = bson.M{"works": bson.M{"$elemMatch": bson.M{"occupation": "programmer"}}}
	err = mg.C.Find(colQuerier).All(&users)
	//err = mg.C.Find(nil).Select(colQuerier).All(&users)

	if err != nil {
		t.Errorf("mg.C.Find(colQuerier).All(&users) / error:%s", err)
		//Cannot use $elemMatch projection on a nested field
	}
	if len(users) == 0 {
		t.Error("mg.C.Find(colQuerier).All(&users) / no data")
	}
	lg.Debugf("result users by find.all: length is %d,\n %+v", len(users), users)
}

func TestGetAllData(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	//#1
	var users []User
	err := mg.C.Find(nil).Sort("age").All(&users)
	if err != nil {
		t.Errorf("mg.C.Find(nil).Sort(\"age\").All(&users) / error:%s", err)
	}
	//t.Logf("result users by find.all: %+v", users)

	//#2 unclear format of json
	mg.GetCol(testColTeacher)

	var v []map[string]interface{}
	err = mg.C.Find(nil).All(&v)
	if err != nil {
		t.Errorf("mg.C.Find(nil).All(&v) / error:%s", err)
		//result argument must be a slice address
	}
	lg.Debugf("result unclear map data by find.all: %+v", v)
	lg.Debugf("result url: %s", v[0]["url"])
}

//-----------------------------------------------------------------------------
// UPDATE
//-----------------------------------------------------------------------------
func TestUpdateOneDataByColumn(t *testing.T) {
	//tu.SkipLog(t)
	//when condition is by column, you should use UpdateAll()

	mg := GetMongo()
	mg.GetCol(testColUser)

	searchName := "Ken"

	// Update (only one record)
	colQuerier := bson.M{"name": searchName}
	updateData := bson.M{"$set": bson.M{"age": 18, "createdAt": time.Now()}}
	err := mg.C.Update(colQuerier, updateData)
	if err != nil {
		t.Errorf("mg.C.Update(colQuerier, updateData) / error:%s", err)
	}

	// check
	user := new(User)
	mg.FindOne(bson.M{"name": searchName}, user)
	lg.Debugf("result user by find: %v", *user)
}

func TestUpdateAllDataByColumn(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	// Update all
	colQuerier := bson.M{"age": 25}
	updateData := bson.M{"$set": bson.M{"age": 26, "createdAt": time.Now()}}
	_, err := mg.C.UpdateAll(colQuerier, updateData)
	if err != nil {
		t.Errorf("mg.C.UpdateAll(colQuerier, updateData) / error:%s", err)
	}

	// check
	var users []User
	mg.C.Find(bson.M{"age": 26}).All(&users)
	lg.Debugf("result user by find: %v", users)
}

func TestUpdateOneDataById(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	//TODO:this is must be gotten automatically
	//get top record

	//userId := "57871a3e340e7601939628e1"
	userID := savedUserID

	idQueryier := bson.ObjectIdHex(userID)

	//oids := make([]bson.ObjectId, len(ids))
	//for i := range ids {
	//	oids[i] = bson.ObjectIdHex(ids[i])
	//}
	//query := bson.M{"_id": bson.M{"$in": oids}}

	// #1. Update (only one record)
	updateData := bson.M{"$set": bson.M{"age": 14, "createdAt": time.Now()}}
	err := mg.C.UpdateId(idQueryier, updateData)
	if err != nil {
		t.Errorf("mg.C.UpdateId(idQueryier, updateData) / error:%s", err)
	}
	// check
	user := new(User)
	mg.C.Find(bson.M{"_id": bson.ObjectIdHex(userID)}).One(user)
	lg.Debugf("result user by find id: %v", *user)
	lg.Debug("- - - - - - - - - - - - - - - - - -")

	// #2. Update (nested element)
	updateData = bson.M{"$set": bson.M{"address.country": "UK", "createdAt": time.Now()}}
	err = mg.C.UpdateId(idQueryier, updateData)
	if err != nil {
		t.Errorf("mg.C.UpdateId(idQueryier, updateData) / error:%s", err)
	}
	// check
	user2 := new(User)
	mg.C.Find(bson.M{"_id": bson.ObjectIdHex(userID)}).One(user2)
	lg.Debugf("result user by find id: %v", *user2)
	lg.Debug("- - - - - - - - - - - - - - - - - -")

	// #3. Update (update by adding element on array)
	updateData = bson.M{"$push": bson.M{"works": bson.M{"occupation": "banker", "company_id": 9}}}
	err = mg.C.UpdateId(idQueryier, updateData)
	if err != nil {
		t.Errorf("mg.C.UpdateId(idQueryier, updateData) / error:%s", err)
	}
	// check
	mg.C.Find(bson.M{"_id": bson.ObjectIdHex(userID)}).One(user2)
	lg.Debugf("result user by find id: %v", *user2)
	lg.Debug("- - - - - - - - - - - - - - - - - -")
}

//-----------------------------------------------------------------------------
// BULK UPDATE
//-----------------------------------------------------------------------------
func TestBulkUpdateByColumn(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	// test data
	colQuerier1 := bson.M{"age": 26}
	updateData1 := bson.M{"$set": bson.M{"age": 27, "createdAt": time.Now()}}

	colQuerier2 := bson.M{"age": 27}
	updateData2 := bson.M{"$set": bson.M{"age": 28, "createdAt": time.Now()}}

	bulk := mg.C.Bulk()
	bulk.UpdateAll(colQuerier1, updateData1)
	bulk.UpdateAll(colQuerier2, updateData2)

	_, err := bulk.Run()
	if err != nil {
		t.Errorf("TestBulkInsert:Insert/ error:%s", err)
	}
}

//-----------------------------------------------------------------------------
// UPSERT
//-----------------------------------------------------------------------------
func TestUpsertOneData(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	// test data
	works := []Work{{Occupation: "programmer", CompanyID: 5}, {Occupation: "programmer", CompanyID: 10}, {Occupation: "programmer", CompanyID: 11}}
	address := Address{ZipCode: "1060047", Country: "Japan", City: "Tokyo", Address1: "港区南麻布1-2-3", Address2: "マンション555"}
	user := User{Name: "NewHarry", Age: 9, Address: address, Works: works, CreatedAt: time.Now()}

	colQuerier := bson.M{"address.zipcode": "1060047"}

	//func (c *Collection) Upsert(selector interface{}, update interface{}) (info *ChangeInfo, err error)
	_, err := mg.C.Upsert(colQuerier, user)
	if err != nil {
		t.Errorf("TestUpsertOneData:Upsert/ error:%s", err)
	}
}

//-----------------------------------------------------------------------------
// DELETE
//-----------------------------------------------------------------------------
func TestDeleteOneData(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	//Delete
	//Even if there are multiple result, one record is deleted.
	err := mg.C.Remove(bson.M{"name": "Ken"})
	if err != nil {
		t.Errorf("TestDeleteData:Remove / Error: %s", err)
	}
}

func TestDeleteMultipleData(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()
	mg.GetCol(testColUser)

	//Delete Multiple
	_, err := mg.C.RemoveAll(bson.M{"name": "Ken"})
	if err != nil {
		t.Errorf("TestDeleteMultipleData:RemoveAll / Error: %s", err)
	}
}

func TestDeleteAllData(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()

	//#1
	mg.GetCol(testColUser)
	err := mg.DelAllDocs("")
	if err != nil {
		t.Errorf("TestDeleteAllData:DelAllDocs / error: %s", err)
	}

	//#2
	mg.GetCol(testColTeacher)
	err = mg.DelAllDocs("")
	if err != nil {
		t.Errorf("TestDeleteAllData:DelAllDocs / error: %s", err)
	}
}

//-----------------------------------------------------------------------------
// Cleanup
//-----------------------------------------------------------------------------
func TestDropCollection(t *testing.T) {
	//tu.SkipLog(t)

	mg := GetMongo()

	//err := mg.DropCol("col01")
	err := mg.DropCol(testColUser)
	if err != nil {
		t.Errorf("DropCol(testColUser):Drop Collection / error: %s", err)
		//ns not found
	}
	err = mg.DropCol(testColTeacher)
	if err != nil {
		t.Errorf("DropCol(testColTeacher):Drop Collection / error: %s", err)
		//ns not found
	}
	err = mg.DropCol(testColCompany)
	if err != nil {
		t.Errorf("DropCol(testColCompany):Drop Collection / error: %s", err)
		//ns not found
	}
}
