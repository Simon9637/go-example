package main

import (
	"log"
	"net/http"
	"time"
	"errors"
	"goStudyProject/router"
	"github.com/spf13/pflag"
	"goStudyProject/config"
	"github.com/spf13/viper"
	"fmt"
)

func main() {
	// Parse conf.yml
	pflag.Parse()

	// Init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Register routers
	router := router.NewRouter()
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("THe router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deploged successfully.")
	}()

	log.Printf("Start to listening the incomming requests on http port %s", viper.GetString("port"))
	log.Printf(http.ListenAndServe(viper.GetString("port"), router).Error())
}

var (
	// 解析默认文件
	cfg = pflag.StringP("conf", "c", "", "Usage")
)

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		//resp, err := http.Get("http://localhost:8088" + "/check/health")
		resp, err := http.Get(fmt.Sprintf("http://localhost%s/check/health", viper.GetString("port")))
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping
		log.Print("Waiting for the router, Retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect the router.")
}
