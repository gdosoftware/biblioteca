package model


type Channel struct {
	Id            string `bson:"id"`
	MerchantId    string `bson:"merchantId"`
	Name          string `bson:"name"`
}
