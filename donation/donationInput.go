package donation

type DonationIdInput struct {
	Id int `uri:"id" binding:"required"`
}
