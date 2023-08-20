package app

type ApplicationGroup struct {
	*TransactionBackupUseCase
	ExchangeQryService
	*UserUseCase
}
