package storage

import (
	"log"

	"github.com/gocql/gocql"
)

const (
	CREATE_KEYSPACE = " CREATE KEYSPACE IF NOT EXISTS " + KEYSPACE_NAME + " WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };"
	CREATE_TABLE    = "create table if not exists candy_shop_db.candy (id uuid PRIMARY KEY, candyname text, candy_price float);"
	KEYSPACE_NAME   = "candy_shop_db"
)

func SetupStorage() (*gocql.Session, error) {
	cluster := createCluster("127.0.0.1:9042", KEYSPACE_NAME)
	session, err := cluster.CreateSession()
	//defer session.Close()
	if err != nil {
		log.Fatal("FATAL", err)
		return nil, err
	}
	createTable(session)
	return session, nil
}

func createCluster(host string, keyspace string) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(host)
	createKeyspace(keyspace, cluster)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.One
	return cluster
}

func createKeyspace(keyspace string, cluster *gocql.ClusterConfig) {
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("FATAL", err)
	}
	defer session.Close()
	if err != nil {
		log.Fatal("FATAL", err)
	}
	if err := session.Query(CREATE_KEYSPACE).Exec(); err != nil {
		log.Fatal("FATAL", err)
	}
	log.Println("INFO", "Configurado keyspace: "+keyspace)
}

func createTable(session *gocql.Session) {
	log.Println("INFO", "Creando tabla si no existe...")
	session.Query(CREATE_TABLE).Exec()
}
