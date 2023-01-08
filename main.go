package main

import (
	"context"
	"fmt"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func WAConnect() (*whatsmeow.Client, error) {
	container, err := sqlstore.New("sqlite3", "file:wapp.db?_foreign_keys=on", waLog.Noop)
	if err != nil {
		return nil, err
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
			return nil, err
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
			return nil, err
		}
	}
	return client, nil
}

func main() {
	wac, err := WAConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wac.Disconnect()

	msg := "Hayirli Cumalar"

	_, err = wac.SendMessage(context.Background(), types.JID{
		User:   "905325701373",
		Server: types.DefaultUserServer,
	}, &waProto.Message{
		Conversation: proto.String(msg),
	})

	if err != nil {
		panic(err)
	}
}
