package postgres

import (
	"log"
	"time"

	"github.com/MarkTBSS/go-kbtg-challenge_5/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (postgres *Postgres) Wallets() ([]wallet.Wallet, error) {
	rows, err := postgres.Database.Query("SELECT * FROM user_wallet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (postgres *Postgres) GetByType(walletType string) ([]wallet.Wallet, error) {
	query := `SELECT * FROM user_wallet WHERE wallet_type = $1`
	rows, err := postgres.Database.Query(query, walletType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}
func (postgres *Postgres) GetByUserID(id string) ([]wallet.Wallet, error) {
	query := `SELECT * FROM user_wallet WHERE user_id = $1`
	rows, err := postgres.Database.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (postgres *Postgres) CreateWallet(wallet wallet.Wallet) (wallet.Wallet, error) {
	query := `
		INSERT INTO user_wallet (user_id, user_name, wallet_name, wallet_type, balance)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, user_name, wallet_name, wallet_type, balance, created_at
	`
	//fmt.Println(wallet)
	row := postgres.Database.QueryRow(
		query,
		wallet.UserID,
		wallet.UserName,
		wallet.WalletName,
		wallet.WalletType,
		wallet.Balance,
	)
	err := row.Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.UserName,
		&wallet.WalletName,
		&wallet.WalletType,
		&wallet.Balance,
		&wallet.CreatedAt,
	)
	if err != nil {
		log.Fatal(err)
	}
	return wallet, err
}
