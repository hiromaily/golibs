package mongodb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoDB Ver.3.x
//https://gist.github.com/border/3489566

//Query and Projection Operators
//https://docs.mongodb.com/manual/reference/operator/query/

//TODO:Mongo session sometimes disconnect and it's not recover automatically.

// MongoInfo is for MongoDB instance
type MongoInfo struct {
	Session  *mgo.Session
	Db       *mgo.Database
	C        *mgo.Collection
	mongoURL string
}

var (
	mgInfo      MongoInfo
	savedDbName string
)

//-----------------------------------------------------------------------------
// Settings
//-----------------------------------------------------------------------------

// New is to create instance for singleton
func New(host, db, user, pass string, port uint16) {
	if mgInfo.Session == nil {
		//[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
		//mgInfo.session, _ = mgo.Dial("mongodb://user:pass@localhost:port/test")
		mgInfo.mongoURL = createMongoURL(host, db, user, pass, port)

		mgInfo.getMongoSession(0)
		//mgInfo.Session.SetMode(mgo.Monotonic, true)
	}
}

// NewInstance makes a new instance
func NewInstance(host, db, user, pass string, port uint16) (*MongoInfo, error) {
	var err error
	mg := &MongoInfo{}

	mg.mongoURL = createMongoURL(host, db, user, pass, port)

	mg.Session, err = mgo.Dial(mg.mongoURL)
	if err != nil {
		return nil, err
	}
	return mg, nil
}

func createMongoURL(host, db, user, pass string, port uint16) string {
	var mongoURL string
	if db == "" {
		mongoURL = fmt.Sprintf("mongodb://%s:%d", host, port)
	} else {
		savedDbName = db
		if user != "" && pass != "" {
			mongoURL = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, pass, host, port, db)
		} else {
			mongoURL = fmt.Sprintf("mongodb://%s:%d/%s", host, port, db)
		}
	}
	return mongoURL
}

func (mi *MongoInfo) getMongoSession(rtnSession uint8) *mgo.Session {
	if mgInfo.Session == nil {
		var err error
		mgInfo.Session, err = mgo.Dial(mi.mongoURL)
		if err != nil {
			panic(err)
		}
	}
	if rtnSession == 1 {
		return mgInfo.Session.Clone()
	}
	return nil
}

// GetMongo is to get instance. singleton architecture
//  This is for singleton design pattern
func GetMongo() *MongoInfo {
	if mgInfo.Session == nil {
		//panic("Before call this, call New in addition to arguments")
		mgInfo.getMongoSession(0)
	}
	return &mgInfo
}

// Close is to close connection
func (mi *MongoInfo) Close() {
	mi.Session.Close()
}

//-----------------------------------------------------------------------------
// Database
//-----------------------------------------------------------------------------

// GetDB is reset session.DB object to mi.db
func (mi *MongoInfo) GetDB(dbName string) *mgo.Database {
	savedDbName = dbName
	//mi.db = mi.session.DB("test")
	//mi.Db = mi.Session.DB(dbName)
	mi.Db = mi.getMongoSession(1).DB(dbName)
	return mi.Db
}

// DropDB is to drop database
func (mi *MongoInfo) DropDB(dbName string) error {
	//err := mi.Session.DB(dbName).DropDatabase()
	err := mi.getMongoSession(1).DB(dbName).DropDatabase()
	return err
}

//-----------------------------------------------------------------------------
// Collection
//-----------------------------------------------------------------------------

// SetExpireOnCollection is to set expired date to collection
func (mi *MongoInfo) SetExpireOnCollection(sessionExpire time.Duration) error {
	sessionTTL := mgo.Index{
		Key:         []string{"createdAt"},
		Unique:      false,
		DropDups:    false,
		Background:  true,
		ExpireAfter: sessionExpire,
	} // sessionExpire is a time.Duration
	fmt.Println(sessionExpire)

	err := mi.C.EnsureIndex(sessionTTL)

	return err
}

// create collection
/*
func (mi *MongoInfo) CreateCol(colName string) error {
	//gopkg.in/mgo.v2/bson.DocElem composite literal uses unkeyed fields
	err := mi.Session.Run(bson.D{{"create", colName}}, nil)
	if err == nil {
		mi.C = mi.Db.C(colName)
	}
	return err
}
*/

// GetCol is to get and set collection
func (mi *MongoInfo) GetCol(colName string) *mgo.Collection {
	if mi.Db == nil {
		if savedDbName == "" {
			mi.Db = mi.getMongoSession(1).DB(savedDbName)
		} else {
			panic("mongo db instance is nil.")
		}
	}
	mi.C = mi.Db.C(colName)
	return mi.C
}

// DropCol is to drop collection
func (mi *MongoInfo) DropCol(colName string) (err error) {
	err = mi.Db.C(colName).DropCollection()
	mi.C = nil
	return
}

//-----------------------------------------------------------------------------
// Document
//-----------------------------------------------------------------------------

// GetCount is to get count
func (mi *MongoInfo) GetCount() int {
	cnt, _ := mi.C.Count()
	return cnt
}

// FindOne is query to find one
func (mi *MongoInfo) FindOne(bd bson.M, data interface{}) error {

	//p := new(Person) //return is address of Person??
	//if name == "" {
	//	mi.C.Find(bson.M{}).One(data)
	//} else {
	//	mi.C.Find(bson.M{"name": name}).One(data)
	//}
	return mi.C.Find(bd).One(data)
}

// DelAllDocs is to delete all documents record from collection. Version3.x
func (mi *MongoInfo) DelAllDocs(colName string) (err error) {
	if colName != "" {
		//mi.Db.C(colName).Remove(bson.M{})
		_, err = mi.Db.C(colName).RemoveAll(bson.M{})
	} else {
		//mi.C.Remove(bson.M{})
		_, err = mi.C.RemoveAll(bson.M{})
	}
	return
}

//-----------------------------------------------------------------------------
// Util
//-----------------------------------------------------------------------------

// ConvertDateTime is to convert datetime GMT
// MongoDB stores times in UTC by default
func ConvertDateTime() {
	//user.CreatedAt.Local()
}

// GetObjectID is to get ObjectId as string
func GetObjectID(ID bson.ObjectId) string {
	//bson.ObjectId
	return ID.Hex()
}

// LoadJSONFile is to load JSON file
func LoadJSONFile(filePath string) ([]byte, error) {
	// Loading jsonfile
	if filePath == "" {
		err := errors.New("nothing JSON file")
		return nil, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
