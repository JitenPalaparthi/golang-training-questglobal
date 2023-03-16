package handler

import (
	"context"
	"time"

	pb "gitlab.stackroute.in/JitenP/golang-training-questglobal/contactspb"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"
)

type Contact struct {
	pb.UnimplementedContactServer
	database.Contact //promoted field
}

func (c *Contact) Create(ctx context.Context, in *pb.ContactIn) (*pb.GeneralResponse, error) {
	contact := new(models.Contact)
	contact.Email = in.Email
	contact.Address = in.Address
	contact.Name = in.Name
	contact.PhoneNumber = in.PhoneNumber
	contact.Status = "created"
	contact.LastModified = time.Now().Unix()
	_, err := c.Add(contact)

	if err != nil {
		return nil, err
	}
	out := new(pb.GeneralResponse)
	out.Code = 201
	out.Message = "Sccessfully inserted"
	return out, nil
}
