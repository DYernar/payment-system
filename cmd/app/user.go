package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"payment.system.com/domain"
	pb "payment.system.com/proto"
)

func (app *application) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	login := ps.ByName("login")
	user, err := app.UserUsecases.GetUserByLogin(login)

	if err != nil {
		app.Logger.Printf("Get user error: %v", err)
		if errors.Is(err, domain.ErrorUserLoginDoesntExists) {
			app.badRequest(w, r, err.Error())
			return
		}
		app.serverError(w, r, err)
		return
	}

	// Get wallets from parsing service

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(app.config.Server.Grpc.Parsing, opts...)

	if err != nil {
		app.Logger.Printf("error dialing grpc %v", err)
		app.serverError(w, r, err)
		return
	}
	defer conn.Close()
	app.Logger.Println("Connected to grpc parsing...")

	client := pb.NewWalletsServiceClient(conn)

	ctx := context.Background()

	iinStruct := pb.Iin{
		Iin: user.Iin,
	}

	resp, err := client.GetWallets(ctx, &iinStruct)

	if err != nil {
		app.Logger.Printf("Get wallets error %v", err)
		app.serverError(w, r, err)
		return
	}

	type wallet struct {
		Id      int64   `json:"id"`
		Name    string  `json:"name"`
		Number  int64   `json:"number"`
		Balance float64 `json:"balance"`
		Iin     string  `json:"iin"`
	}

	wallets := []wallet{}
	for _, w := range resp.Wallets {
		wallets = append(wallets, wallet{
			Id:      w.Id,
			Name:    w.Name,
			Number:  w.Number,
			Balance: w.Balance,
			Iin:     w.Iin,
		})
	}

	err = app.writeJson(w, http.StatusOK, envelope{"user": user, "wallets": wallets}, nil)

	if err != nil {
		app.Logger.Printf("get user err %v", err)
		app.serverError(w, r, err)
		return
	}
}
