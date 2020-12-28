package main

import (
	"log"

	"github.com/tada-team/tdclient"
	"github.com/tada-team/tdclient/examples"
	"github.com/tada-team/tdproto"
)

const UNKNOWN = "unknown"

func main() {

	checkedUID := tdproto.NewJID("d-df2d36bf-3bbb-49f2-9a1e-5e513b7597d7")

	settings := examples.NewSettings()
	settings.RequireToken()
	settings.RequireTeam()
	settings.Parse()

	session, err := tdclient.NewSession(settings.Server)
	if err != nil {
		panic(err)
	}

	session.SetToken(settings.Token)
	session.SetVerbose(settings.Verbose)

	contacts, err := session.Contacts(settings.TeamUid)
	for _, contact := range contacts {
		//log.Println(contact.Jid.String(), checkedUID.String())
		if contact.Jid.String() == checkedUID.String() {
			log.Println("GOT IT!", contact.TeamStatus, contact.IsArchive, *contact.LastActivity)
			return
		}
	}
	log.Println("Not found")
}
