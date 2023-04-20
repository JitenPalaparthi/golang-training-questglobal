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

func (c *Contact) Create(ctx context.Context, in *pb.ContactCreateMessage) (*pb.GeneralResponse, error) {
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

func (c *Contact) Delete(ctx context.Context, in *pb.ByIDRequest) (*pb.GeneralResponse, error) {
	out := new(pb.GeneralResponse)

	r, err := c.Contact.Delete(int(in.ID))
	if err != nil {
		return nil, err
	}
	if r > 0 {
		out.Code = 202
		out.Message = "Sccessfully deleted"
		return out, nil
	} else {
		out.Code = 400
		out.Message = "invalid id or no data"
		return out, nil
	}
}

func (c *Contact) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.GeneralResponse, error) {
	out := new(pb.GeneralResponse)

	_, err := c.Contact.Update(int(in.ID), in.Data.AsMap())

	if err != nil {
		return nil, err
	}
	out.Code = 200
	out.Message = "Sccessfully updated"
	return out, nil
}

func (c *Contact) GetBy(ctx context.Context, in *pb.ByIDRequest) (*pb.ContactMessage, error) {
	out := new(pb.ContactMessage)

	contact, err := c.Contact.GetBy(int(in.ID))
	if err != nil {
		return nil, err
	}

	out.ID = contact.ID
	out.Name = contact.Name
	out.Address = contact.Address
	out.Email = contact.Email
	out.PhoneNumber = contact.PhoneNumber
	out.Status = contact.Status
	out.LastModified = contact.LastModified

	return out, nil
}

func (c *Contact) GetAll(ctx context.Context, in *pb.NoIn) (*pb.ContactMessages, error) {

	out := new(pb.ContactMessages)

	contacts, err := c.Contact.GetAll()
	if err != nil {
		return nil, err
	}

	for _, v := range contacts {
		c := new(pb.ContactMessage)
		c.ID = v.ID
		c.Name = v.Name
		c.Address = v.Address
		c.Email = v.Email
		c.PhoneNumber = v.PhoneNumber
		c.Status = v.Status
		c.LastModified = v.LastModified
		out.Contacts = append(out.Contacts, c)

	}
	return out, nil
}
