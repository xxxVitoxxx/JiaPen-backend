package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/config"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/http/rest"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/conn"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/mysqlDB"
)

func init() {
	rootCmd.AddCommand(messageCmd)
}

var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "message board",
	Long:  "message board",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Message Service Start")
		defer func() {
			if panicErr := recover(); panicErr != nil {
				log.Println(panicErr)
			}
		}()

		err := config.LoadRun()
		if err != nil {
			panic(err)
		}

		r := gin.Default()
		r.Use(gzip.Gzip(gzip.BestSpeed))

		mysqlDb := conn.CheckConnect()

		// message POST
		messageHandler := rest.NewMessageHandler(
			message.NewService(
				mysqlDB.NewMessageRepo(mysqlDb),
			),
		)
		messageHandler.Router(r)

		// message GET
		queryMessageHandler := rest.NewQueryMessageHandler(
			mysqlDB.NewQueryMessageRepo(mysqlDb),
		)
		queryMessageHandler.Router(r)

		go func() {
			if err := r.Run(fmt.Sprintf(":%s", viper.GetString("message.port"))); err != nil {
				log.Println("Error: ", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%s", viper.GetString("message.port")),
			Handler: r,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Println("Message Service Shutdown...", err)
		}

		log.Println("Message Service Exiting")
	},
}
