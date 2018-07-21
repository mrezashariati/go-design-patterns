/* Factory Package implemnets factory method design pattern whith a simple example */
package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPayment(m int) (PaymentMethod, error) {
	switch m{
	case Cash:
		return new(CashPM),nil
		
	case DebitCard:
		return new(CreditCardPM),nil

	default :
		return nil,errors.New(fmt.Sprintf("PaymentMethod %d not recongnized\n",m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}
type CreditCardPM struct{}

func (c *CashPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid using cash\n",amount)
}

func (d *DebitCardPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid using debitcard\n",amount)
}

func (c *CreditCardPM) Pay(amount float64) string {
	return fmt.Sprintf("%#0.2f paid using new credit card implementation\n",amount)
}