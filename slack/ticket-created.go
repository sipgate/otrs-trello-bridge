package slack

import (
	"fmt"
	"github.com/nlopes/slack"
	"github.com/sipgate/otrs-bridge/contract"
	"github.com/sipgate/otrs-bridge/otrs"
	"github.com/spf13/viper"
)

type TicketCreatedInteractor struct{}

func (t *TicketCreatedInteractor) HandleTicketCreated(ticketID string, ticket otrs.Ticket) error {
	api := slack.New(viper.GetString("slack.apiToken"))
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Title:     ticket.Title,
		TitleLink: otrs.MakeTicketUrl(ticket),
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage(viper.GetString("slack.channelId"), "New OTRS Ticket", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	return err
}

func NewTicketCreatedInteractor() contract.TicketCreatedInteractor {
	return &TicketCreatedInteractor{}
}
