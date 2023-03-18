package main

import (
	"flag"
	"fmt"
	"local/config"
	"local/db"
	"local/svr"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	queryToken = flag.Bool("query", false, "query all access tokens")
	addToken   = flag.String("add", "", "add access token")
	delToken   = flag.String("del", "", "delete access token")
	updToken   = flag.String("upd", "", "update access token")
)

func main() {
	parseArgs()
	initEnv()
	db.Init(config.DBPath + "/" + config.DB_BAME)
	defer db.Uninit()

	if flag.NFlag() > 0 {
		runAsCln()
	} else {
		runAsSvr()
	}
}

func parseArgs() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-add <token>] | [del <token>] | [upd <oldToken,newToken>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Version:", VERSION)
		fmt.Println("Configure: ~/.config/msgcenter/config.json")
		fmt.Println("Database: ~/.local/share/msgcenter/msgcenter.db")
		fmt.Println("cert.pem: ~/.config/msgcenter/cert.pem")
		fmt.Println("key.pem: ~/.config/msgcenter/key.pem")
	}
	flag.Parse()
}

func runAsCln() {
	if *queryToken {
		if tokens, err := db.QueryAccessTokens(); err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println(tokens)
		}
	} else if *addToken != "" {
		if err := db.AddAccessToken(*addToken); err != nil {
			log.Fatalln(err)
		}
	} else if *delToken != "" {
		if err := db.DelAccessToken(*delToken); err != nil {
			log.Fatalln(err)
		}
	} else if *updToken != "" {
		tokens := strings.Split(*updToken, ",")
		if len(tokens) != 2 {
			log.Fatalln("Invalid format:", *updToken)
		}
		if err := db.UpdateAccessToken(tokens[0], tokens[1]); err != nil {
			log.Fatalln(err)
		}
	}
}

func runAsSvr() {
	svr.Start()
}

func initEnv() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
    config.Init()
}
