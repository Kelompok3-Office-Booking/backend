package transactions

type transactionUsecase struct {
	transactionRepository Repository
}

func NewTransactionUsecase(tr Repository) Usecase {
	return &transactionUsecase{
		transactionRepository: tr,
	}
}

func (tu *transactionUsecase) GetAll() []Domain {
	return tu.transactionRepository.GetAll()
}

func (tu *transactionUsecase) Create(transactionDomain *Domain) Domain {
	return tu.transactionRepository.Create(transactionDomain)
}
