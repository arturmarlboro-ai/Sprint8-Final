package structs

type Parcel struct {
	Number    int
	Client    int
	Status    string
	Address   string
	CreatedAt string
}

const (
	ParcelStatusRegistered = "registered"
	ParcelStatusSent       = "sent"
	ParcelStatusDelivered  = "delivered"
)
