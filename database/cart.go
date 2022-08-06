package database

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrCantFindProduct    = errors.New("cant find the product")
	ErrCantDecodeProduct  = errors.New("cant find the product")
	ErrUserIdIsNotValid   = errors.New("this user id is not valid")
	ErrCantUpdateUser     = errors.New("cant add this product to cart")
	ErrCantRemoveItemCart = errors.New("cant remove this item from cart")
	ErrCantGetItem        = errors.New("cant get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cant purchase")
)

func AddProductToCart() gin.HandlerFunc {

}

func RemoveCartItem() gin.HandlerFunc {

}

func BuyItemFromCart() gin.HandlerFunc {

}

func InstantBuy() gin.HandlerFunc {

}
