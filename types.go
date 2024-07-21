package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AuthorID      primitive.ObjectID `json:"author_id,omitempty" bson:"author_id,omitempty"`
	AuthorName    string             `json:"author_name,omitempty" bson:"author_name,omitempty"`
	Date          string             `json:"date,omitempty" bson:"date,omitempty"`
	Sport         string             `json:"sport,omitempty" bson:"sport,omitempty"`
	ImageSrc      string             `json:"image_src,omitempty" bson:"image_src,omitempty"`
	FigureCaption string             `json:"figure_caption,omitempty" bson:"figure_caption,omitempty"`
	FigureCreator string             `json:"figure_creator,omitempty" bson:"figure_creator,omitempty"`
	FigureCredit  string             `json:"figure_credit,omitempty" bson:"figure_credit,omitempty"`
	Content       []string           `json:"content,omitempty" bson:"content,omitempty"`
	Alt           string             `json:"alt_text,omitempty" bson:"alt_text,omitempty"`
	Title         string             `json:"title,omitempty" bson:"title,omitempty"`
}
