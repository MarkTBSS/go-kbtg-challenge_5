package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets() ([]Wallet, error)
	GetByType(walletType string) ([]Wallet, error)
	GetByUserID(id string) ([]Wallet, error)
	CreateWallet(wallet Wallet) (Wallet, error)
}

func New(database Storer) *Handler {
	return &Handler{store: database}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
//
// @Summary		Get all wallets
// @Description	Get all wallets
// @Tags		wallet
// @Accept		json
// @Produce		json
// @Success		200	{object}	Wallet
// @Router		/api/v1/wallets [get]
// @Failure		500	{object}	Err
func (handler *Handler) WalletHandler(context echo.Context) error {
	wallets, err := handler.store.Wallets()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, wallets)
}

// QueryParamHandler
//
// @Summary 	Get all wallet_type
// @Description Get all wallet_type
// @Tags 		wallet_type
// @Param 		wallet_type query string false "wallet type" Enums(Savings, Credit Card, Crypto Wallet)
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} Wallet
// @Router 		/api/v1/wallets/query [get]
// @Failure 	500 {object} Err
func (handler *Handler) QueryParamHandler(context echo.Context) error {
	walletType := context.QueryParam("wallet_type")
	wallets, err := handler.store.GetByType(walletType)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, wallets)
}

// PathParamHandler
//
//	@Summary		Get all user_id
//	@Description	Get all user_id
//	@Tags			user_id
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/users/:id/wallets [get]
//	@Failure		500	{object}	Err
func (handler *Handler) PathParamHandler(context echo.Context) error {
	id := context.Param("id")
	wallets, err := handler.store.GetByUserID(id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, wallets)
}

func (handler *Handler) BindingDataHandler(context echo.Context) error {
	var wallet Wallet
	err := context.Bind(&wallet)
	if err != nil {
		return context.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	//fmt.Println(wallet)
	walletResponse, err := handler.store.CreateWallet(wallet)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, walletResponse)
}
