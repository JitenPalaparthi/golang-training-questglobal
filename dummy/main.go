package handler

import (
	"fmt"
	"net/http"
)

func PurchaseItem(w http.ResponseWriter, r *http.Request) {

	details, err := database.Purchase(Item) // this is to store data in the database and assume it as purchase

	if err != nil {
		return err // do some abt err
	}

	// write code to send emailway

	// sync way of Microservice communication

	err := emailService.Send(details.Email, details.Item+" purchased.Booking id:", details.ItemId)

	// examples of failures
	// what if email server is failed to send email
	// what if email server is taking 1 or 2 mins to send email
	// what if email server has been changed//
	// what if you update new email server

	// async

	kafka.Produce("topic-item-purchased", details)

	// kafka must be up and run 24/7 .. That is taken care by kafka admin which HA

	// to send email, you create  saperate service which is a consumer of kafaka
	// a server is running as a saperate application which is reading data from kafaka and sending emails
	// no dependency on your application.

	if err != nil {
		return err // do some abt err
	}

	fmt.Fprintln(w, details)

}
