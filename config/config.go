package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ListenAddr string
	TestMode   bool
	EnableTLS  bool
}

const DB_BAME string = "msgcenter.db"

var ConfPath, DBPath, CertFile, KeyFile string

var AppConf Config = Config{ListenAddr: ":8001", TestMode: false, EnableTLS: false}

func Init() {
	home, e := os.LookupEnv("HOME")
	if !e {
		log.Fatalln("can't find $HOME")
	}

	ConfPath = home + "/.config/msgcenter"
	DBPath = home + "/.local/share/msgcenter"
	CertFile = home + "/.config/msgcenter/cert.pem"
	KeyFile = home + "/.config/msgcenter/key.pem"
	for _, path := range []string{ConfPath, DBPath} {
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(path, os.ModePerm); err != nil {
					log.Fatalln("Error:", err)
				}
			} else {
				log.Fatalln("Error:", err)
			}
		}
	}
	saveDefaultCert()
	saveDefaultKey()
	loadConf()
}

func loadConf() {
	config := ConfPath + "/config.json"
	if _, err := os.Stat(config); err != nil {
		if os.IsNotExist(err) {
			defaultConf(config)
		} else {
			log.Fatalln("Error:", err)
		}
	} else {
		if js, err := os.ReadFile(config); err != nil {
			log.Fatalln("Error:", err)
		} else {
			if err := json.Unmarshal(js, &AppConf); err != nil {
				log.Fatalln("Error:", err)
			}
		}
	}
}

func defaultConf(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	if err := enc.Encode(AppConf); err != nil {
		log.Fatalln("Error:", err)
	}
}
