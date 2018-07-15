package bitcointrade_test

import (
	"bitcointrade"
	"os"
	"testing"
)

func TestFeeBitcoin(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	cryptoService := &bitcointrade.Crypto{
		Crypto: "bitcoin",
	}
	resp, err := cryptoService.Fee()
	if err != nil {
		t.Errorf("TestFeeBitcoin => '%s'", err)
		return
	}
	if resp.Search("data").Data() == nil {
		t.Error("TestFeeBitcoin => 'data is null'")
	}
}
func TestListWidthDrawBitcoin(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	cryptoService := &bitcointrade.Crypto{
		Crypto: "bitcoin",
	}
	resp, err := cryptoService.ListWidthDraw()
	if err != nil {
		t.Errorf("TestListWidthDrawBitcoin => '%s'", err)
		return
	}

	if resp.Search("withdrawals").Data() == nil {
		t.Error("TestListWidthDrawBitcoin => 'withdrawals is null'")
	}
}
func TestListDepositsBitcoin(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	cryptoService := &bitcointrade.Crypto{
		Crypto: "bitcoin",
	}
	resp, err := cryptoService.ListDeposits()
	if err != nil {
		t.Errorf("TestListDepositsBitcoin => '%s'", err)
		return
	}

	if resp.Search("deposits").Data() == nil {
		t.Error("TestListDepositsBitcoin => 'deposits is null'")
	}
}

func TestCreateWidthDrawBitcoin(t *testing.T) {
	cryptoService := &bitcointrade.Crypto{
		Crypto: "bitcoin",
	}
	withdraw := &bitcointrade.WidthDraw{
		Destination: "mj9NZ98y1gwmgvxZytvqbMbUSo9M3E7CYN",
		Amount:      0.1,
		Fee:         0.11,
		FeeType:     "slow",
	}
	resp, err := cryptoService.CreateWidthDraw(withdraw)
	if err != nil {
		t.Errorf("TestCreateWidthDrawBitcoin => '%s'", err)
		return
	}

	if resp.Search("deposits").Data() == nil {
		t.Error("TestCreateWidthDrawBitcoin => 'deposits is null'")
	}

}

func TestFeeEthereum(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	cryptoService := &bitcointrade.Crypto{
		Crypto: "ethereum",
	}
	resp, err := cryptoService.Fee()
	if err != nil {
		t.Errorf("TestFeeEthereum => '%s'", err)
		return
	}
	if resp.Search("data").Data() == nil {
		t.Error("TestFeeEthereum => 'data is null'")
	}
}
func TestListWidthDrawEthereum(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	cryptoService := &bitcointrade.Crypto{
		Crypto: "ethereum",
	}
	resp, err := cryptoService.ListWidthDraw()
	if err != nil {
		t.Errorf("TestListWidthDrawEthereum => '%s'", err)
		return
	}

	if resp.Search("withdrawals").Data() == nil {
		t.Error("TestListWidthDrawEthereum => 'withdrawals is null'")
	}
}
func TestListDepositsEthereum(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	cryptoService := &bitcointrade.Crypto{
		Crypto: "ethereum",
	}
	resp, err := cryptoService.ListDeposits()
	if err != nil {
		t.Errorf("TestListDepositsEthereum => '%s'", err)
		return
	}

	if resp.Search("deposits").Data() == nil {
		t.Error("TestListDepositsEthereum => 'deposits is null'")
	}
}

func TestCreateWidthDrawEthereum(t *testing.T) {
	cryptoService := &bitcointrade.Crypto{
		Crypto: "ethereum",
	}
	withdraw := &bitcointrade.WidthDraw{
		Destination: "mj9NZ98y1gwmgvxZytvqbMbUSo9M3E7CYN",
		Amount:      0.1,
		Fee:         0.11,
		FeeType:     "slow",
	}
	resp, err := cryptoService.CreateWidthDraw(withdraw)
	if err != nil {
		t.Errorf("TestCreateWidthDrawEthereum => '%s'", err)
		return
	}

	if resp.Search("deposits").Data() == nil {
		t.Error("TestCreateWidthDrawEthereum => 'deposits is null'")
	}

}
