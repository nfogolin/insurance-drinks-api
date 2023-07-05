package requests

type GetDrinkByIdRequest struct {
	Id int64 `uri:"drinkId" binding:"required"`
}
