package inputs

type DonationIdInput struct {
	Id int `uri:"id" binding:"required"`
}
