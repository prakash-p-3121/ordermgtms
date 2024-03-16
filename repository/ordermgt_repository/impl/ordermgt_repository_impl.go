package impl

import (
	"database/sql"
	"errors"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/mysqllib"
	"github.com/prakash-p-3121/ordermgtmodel"
	"sync"
)

type OrderRepositoryImpl struct {
	SingleStoreConnection *sql.DB
	ShardConnectionsMap   *sync.Map
}

func (repository *OrderRepositoryImpl) CreateOrder(shardID int64,
	idGenResp *idgenmodel.IDGenResp,
	req *ordermgtmodel.OrderCreateReq) errorlib.AppError {

	db, err := mysqllib.RetrieveShardConnectionByShardID(repository.ShardConnectionsMap, shardID)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}

	qry := `INSERT INTO orders (id, 
             id_bit_count, 
             product_id,
             seller_id,
             buyer_id,
             listing_id,
             currency,
             price,
             state_id
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) ;`
	_, err = db.Exec(qry, idGenResp.ID,
		idGenResp.BitCount,
		req.ProductID,
		req.SellerID,
		req.BuyerID,
		req.ListingID,
		req.Currency,
		req.Amount,
		ordermgtmodel.OrderCreated,
	)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return nil
}

func (repository *OrderRepositoryImpl) FindOrderByID(shardID int64, orderID string) (*ordermgtmodel.Order, errorlib.AppError) {
	db, err := mysqllib.RetrieveShardConnectionByShardID(repository.ShardConnectionsMap, shardID)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	var order ordermgtmodel.Order
	qry := `SELECT id, 
			id_bit_count, 
			product_id, 
			seller_id, 
			buyer_id, 
			listing_id, 
			currency, 
			price, 
			state_id, 
			created_at, 
			updated_at FROM orders WHERE id = ?;`
	row := db.QueryRow(qry, orderID)
	err = row.Scan(&order.ID,
		&order.IDBitCount,
		&order.ProductID,
		&order.SellerID,
		&order.BuyerID,
		&order.ListingID,
		&order.Currency,
		&order.Amount,
		&order.StateID,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &order, nil
}
