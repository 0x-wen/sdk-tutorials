package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	auction "github.com/cosmos/sdk-tutorials/tutorials/nameservice/base/x/auction"
)

// GetNameRecord 获取指定名称记录，与内存储存交互
func (k Keeper) GetNameRecord(ctx sdk.Context, name string) (auction.Name, error) {
	nameRecord, err := k.Names.Get(ctx, name)
	if err != nil {
		return auction.Name{}, err
	}

	return nameRecord, nil
}

func (k Keeper) reserveName(ctx sdk.Context, nr auction.Name) error {
	owner, _ := sdk.AccAddressFromBech32(nr.Owner)
	// 转账给模块账号
	if err := k.executeTransfer(ctx, owner, nr.Amount); err != nil {
		return err
	}
	// k 设置存储, [auction.Name.Name] -> auction.Name
	if err := k.Names.Set(ctx, nr.Name, nr); err != nil {
		// TODO: Custom error
		return err
	}

	return nil
}

func (k Keeper) checkAvailability(ctx sdk.Context, name string) error {
	hasName, err := k.Names.Has(ctx, name)
	if hasName {
		return fmt.Errorf("%s is already reserved", name)
	}
	if err != nil {
		return fmt.Errorf("unexpected error validating reservation :: %w", err)
	}
	return nil
}

func (k Keeper) validateAndFormat(ctx sdk.Context, msg *auction.MsgBid) (auction.Name, error) {
	nr := auction.Name{}

	sender, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nr, err
	}

	resolveAddr, err := sdk.AccAddressFromBech32(msg.ResolveAddress)
	if err != nil {
		return nr, fmt.Errorf("invalid resolve address %s", msg.ResolveAddress)
	}

	if err := k.checkSufficientBalance(ctx, sender, msg.Amount); err != nil {
		return nr, err
	}

	return auction.Name{Name: msg.Name, Owner: sender.String(), ResolveAddress: resolveAddr.String(), Amount: msg.Amount}, nil
}

func (k Keeper) checkSufficientBalance(ctx sdk.Context, senderAddr sdk.AccAddress, amount sdk.Coins) error {
	found, payment := amount.Find("uatom")
	if !found {
		return fmt.Errorf("invalid coin denom - %s, default denom: %s", amount, "uatom")
	}

	balance := k.bk.GetBalance(ctx, senderAddr, payment.Denom)
	_, insufficientBalance := sdk.Coins{balance}.SafeSub(payment)
	if insufficientBalance {
		return fmt.Errorf("price for name exceeds account balance: %s < %s", balance, payment)
	}
	return nil
}

func (k Keeper) executeTransfer(ctx sdk.Context, senderAddr sdk.AccAddress, amount sdk.Coins) error {
	err := k.bk.SendCoinsFromAccountToModule(ctx, senderAddr, auction.ModuleName, amount)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) GetNames(ctx sdk.Context) ([]*auction.Name, error) {
	iter, err := k.Names.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	var names []*auction.Name
	values, err := iter.Values()
	if err != nil {
		return nil, err
	}
	for i := range values {
		names = append(names, &values[i])
	}
	return names, nil
}
