package cassandra

import (
	"github.com/gocql/gocql"
)

//https://github.com/gocql/gocql

// CassInfo is for session data of Cassandora
type CassInfo struct {
	Session *gocql.Session
}

var cassInfo CassInfo

// New is to create Cassandora instance
func New(hosts []string, port int, keyspace string) (err error) {
	// connect to the cluster
	//cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	if port == 0 {
		cluster.Port = 9042
	} else {
		cluster.Port = port
	}
	//cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum

	cassInfo.Session, err = cluster.CreateSession()
	return err
}

// GetCass is to get Cassandora instance singleton architecture
func GetCass() *CassInfo {
	if cassInfo.Session == nil {
		return nil
		//panic("Before call this, call New in addition to arguments")
	}
	return &cassInfo
}

// SetKeySpace is to change keyspace(database)
// TODO:What happened?
//func (cs *CassInfo) SetKeySpace(keyspace string) error {
//	err := cs.Session.Query(fmt.Sprintf("use %s", keyspace)).Exec()
//	return err
//}

// Close is to close connection
func (cs *CassInfo) Close() {
	cs.Session.Close()
}
