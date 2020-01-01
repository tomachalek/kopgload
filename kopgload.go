package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"kops/dataload"
	"kops/db"
	"log"

	"github.com/tomachalek/vertigo/v2"
)

type Conf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"dbname"`
}

func loadConf(confPath string) *Conf {
	rawData, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}
	var conf Conf
	err2 := json.Unmarshal(rawData, &conf)
	if err2 != nil {
		log.Fatal(err2)
	}
	return &conf
}

func main() {
	flag.Parse()
	conf := loadConf(flag.Arg(0))
	db := db.ConnectDB(conf.Host, conf.User, conf.Password, conf.DB)
	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	pc := &vertigo.ParserConf{
		InputFilePath:         "/home/tomas/work/data/corpora/vertical/vertikala",
		Encoding:              "utf-8",
		StructAttrAccumulator: "comb",
	}
	loader := dataload.NewLoader(txn)
	loader.Prepare()
	for i := 0; i < 20; i++ {
		err = vertigo.ParseVerticalFile(pc, loader)
	}
	log.Println("ERROR: ", err)
	loader.Finish()
	txn.Commit()
	db.Close()
}
