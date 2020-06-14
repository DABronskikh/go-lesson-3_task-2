package card

type Card = struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

type Transaction = struct {
	Id     int64
	Amount int64
	Date   int64
	MCC    string
	status string
	Type   string
	Status string
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func SumByMCC(transactions []Transaction, mcc []string) (total int64) {

	for _, transaction := range transactions {
		if isMCC(transaction.MCC, mcc) {
			total += transaction.Amount
		}
	}

	return total
}

func isMCC(mcc string, mccArr []string) bool {
	for _, candidate := range mccArr {
		if candidate == mcc {
			return true
		}
	}
	return false
}

func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5912": "Аптеки",
	}

	value, ok := mcc[code]
	if ok {
		return value
	}

	return "Категория не указана"
}
