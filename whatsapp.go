package main

import (
	"context"
	"fmt"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
	"os"
)

var wpClient *whatsmeow.Client

func waConnect() {
	container, err := sqlstore.New("sqlite3", "file:wapp.db?_foreign_keys=on", waLog.Noop)
	if err != nil {
		panic(err)
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	client := whatsmeow.NewClient(deviceStore, waLog.Noop)
	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err := client.Connect()
		if err != nil {
			panic(err)
		}
	}
	wpClient = client
}

func SendMessage(number string, msg string) error {
	_, err := wpClient.SendMessage(context.Background(), types.JID{
		User:   number,
		Server: types.DefaultUserServer,
	}, &waProto.Message{
		Conversation: proto.String(msg),
	})
	return err

}
