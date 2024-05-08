/*
Copyright © 2022 Openerd <jsmoon@openerd.com>
*/
package cmd

import (
  "context"
  "fmt"
  "git.dev.opnd.io/gc/backend-admin/pkg/config/db"
  "git.dev.opnd.io/gc/backend-admin/pkg/router"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "git.dev.opnd.io/gc/backend-admin/pkg/config"
  "git.dev.opnd.io/gc/backend-admin/pkg/logger"
  "github.com/labstack/echo/v4"
  "github.com/spf13/cobra"
  "golang.org/x/sync/errgroup"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
  Use:   "serve",
  Short: "serve",
  Long:  `serve`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("serve called")
    // panic 대응
    defer func() {
      if r := recover(); r != nil {
        logger.Logger.Error(r)
      }
    }()

    config.Init()

    db.Init()
    //kvstore.Init()
    //s3.Init()

    e := echo.New()

    router.SetupBaseHandler(e, nil)

    //grpcServer := router.NewGrpcServer()
    //
    //router.RouterSetup(e, grpcServer)

    mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
    defer stop()

    g, gCtx := errgroup.WithContext(mainCtx)
    g.Go(func() error {
      logger.Logger.Info("backend Start")
      e.HideBanner = true
      if config.CertKeyPair != nil {
        if err := e.StartTLS(fmt.Sprintf(":%d", config.Config.Port), config.Config.TLS.CertFile, config.Config.TLS.KeyFile); err != nil && err != http.ErrServerClosed {
          logger.Logger.Error("shutting down the server : ", err)
          return err
        }
      } else {
        if err := e.Start(fmt.Sprintf(":%d", config.Config.Port)); err != nil && err != http.ErrServerClosed {
          logger.Logger.Error("shutting down the server : ", err)
          return err
        }
      }
      return nil
    })

    g.Go(func() error {
      <-gCtx.Done()
      //grpcServer.GracefulStop()

      ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
      defer cancel()

      if err := e.Shutdown(ctx); err != nil {
        logger.Logger.Error(err)
        return err
      }
      logger.Logger.Info("backend graceful shutting down")
      return nil
    })

    if err := g.Wait(); err != nil {
      logger.Logger.Errorf("exit reason: %s \n", err)
    }

  },
}

func init() {
  rootCmd.AddCommand(serveCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // serveCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
