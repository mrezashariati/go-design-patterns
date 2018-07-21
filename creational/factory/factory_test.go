package factory

import (
	"strings"
	"testing"
)

func TestCreatPaymentMethodCash(t *testing.T) {
	payment, err := GetPayment(Cash)
	if err != nil {
		t.Fatal("A payment method of type cash must exist")
	}

	msg := payment.Pay(10.5)
	if !strings.Contains(msg, "paid using cash") {
		t.Fatal("The cash payment method wasn'n right")
	}
	t.Log("LOG:", msg)
}
func TestCreatPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPayment(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type Debitcard must exist")
	}

	msg := payment.Pay(22.8)
	if !strings.Contains(msg, "paid using new credit card") {
		t.Fatal("The debitcard payment method wasn'n right")
	}
	t.Log("LOG:", msg)
}
func TestGetPeymantMethodNotExists(t *testing.T) {
	_, err := GetPayment(20)
	if err == nil {
		t.Fatal("A payment method with id 20 must return nil")
	}
	t.Log("LOG:", err)
}
