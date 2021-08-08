package connection

import (
	"os"

	"gopkg.in/mgo.v2"
)

func MongoDB() *mgo.Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")},
	})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return session
}
