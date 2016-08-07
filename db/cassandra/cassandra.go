package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

//https://github.com/gocql/gocql

type CassInfo struct {
	Session *gocql.Session
}

var cassInfo CassInfo

func New(hosts []string, keyspace string) {
	// connect to the cluster
	//cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Port = 9042
	cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum
	cassInfo.Session, _ = cluster.CreateSession()
}

// singleton architecture
func GetCass() *CassInfo {
	if cassInfo.Session == nil {
		return nil
		//panic("Before call this, call New in addtion to arguments")
	}
	return &cassInfo
}

// change keyspace(database)
// TODO:What happend?
func (cs *CassInfo) SetKeySpace(keyspace string) error {
	err := cs.Session.Query(fmt.Sprintf("use %s", keyspace)).Exec()
	return err
}

// close
func (cs *CassInfo) Close() {
	cs.Session.Close()
}
