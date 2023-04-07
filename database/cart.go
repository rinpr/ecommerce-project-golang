package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can't find the product!")
	ErrCantDecodeProducts = errors.New("Can't find the product!")
	ErrUserIdIsNotValid   = errors.New("This user is not valid.")
	ErrCantUpdateUser     = errors.New("Can't add this product to the cart.")
	ErrCantRemoveItemCart = errors.New("Can't remove this product to the cart.")
	ErrCantGetItem        = errors.New("Unable to get the item from the cart.")
	ErrCantBuyCartItem    = errors.New("Unable to update the purchase.")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
