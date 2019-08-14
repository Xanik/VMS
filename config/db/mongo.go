package dbs

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

//ConnectMongodbURL returns a connection to a mongodb instance through a connection string in the configuration file
func ConnectMongodbURL(url string) *mgo.Session {
	env := viper.GetString("env")

	session, _ := mgo.Dial(viper.GetString(env + ".db.mongo.url"))

	err := session.Ping()

	if err != nil {
		log.Printf("db connection error: %s", err.Error())
	}
	return session
}

//ConnectMongodb returns a connection to a mongodb instance through a connection options without ssl being true in the configuration file
func ConnectMongodb() *mgo.Session {
	env := viper.GetString("env")

	params := &mgo.DialInfo{
		Username: viper.GetString(env + ".db.mongo.user"),
		Password: viper.GetString(env + ".db.mongo.password"),
		Addrs:    viper.GetStringSlice(env + ".db.mongo.addrs"),
		Database: viper.GetString(env + ".db.mongo.db"),
	}

	session, _ := mgo.DialWithInfo(params)

	err := session.Ping()

	if err != nil {
		log.Printf("db connection error: %s", err.Error())
	}
	return session
}

//ConnectMongodbTLS returns a connection to a mongodb instance through a connection options with ssl set to true in the configuration file
func ConnectMongodbTLS() *mgo.Session {
	env := viper.GetString("env")

	params := &mgo.DialInfo{
		Username: viper.GetString(env + ".db.mongo.user"),
		Password: viper.GetString(env + ".db.mongo.password"),
		Addrs:    viper.GetStringSlice(env + ".db.mongo.addrs"),
		Database: viper.GetString(env + ".db.mongo.db"),
	}

	params.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), &tls.Config{})
		return conn, err
	}

	session, _ := mgo.DialWithInfo(params)

	err := session.Ping()

	fmt.Println(err)
	if err != nil {
		log.Printf("db connection error: %s", err.Error())
	}
	return session
}

//CreateUniqueIndex creates a unque mogodb index
func CreateUniqueIndex(c *mgo.Collection, data []string) error {

	for _, key := range data {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := c.EnsureIndex(index); err != nil {
			return err
		}
	}
	return nil
}
