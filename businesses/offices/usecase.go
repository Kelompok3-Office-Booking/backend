package offices

type OfficeUsecase struct {
	officeRepository Repository
}

func NewOfficeUsecase(or Repository) Usecase {
	return &OfficeUsecase{
		officeRepository: or,
	}
}

func (n *OfficeUsecase) Create(input *Domain) Domain {
	return n.officeRepository.Create(input)
}

func (n *OfficeUsecase) GetAll() []Domain {
	return n.officeRepository.GetAll()
}

func (n *OfficeUsecase) GetByID(id string) Domain {
	return n.officeRepository.GetByID(id)
}

func (n *OfficeUsecase) Delete(id string) bool {
	return n.officeRepository.Delete(id)
}

func (n *OfficeUsecase) SearchByCity(city string) []Domain {
	return n.officeRepository.SearchByCity(city)
}

func (n *OfficeUsecase) SearchByRate(rate string) []Domain {
	return n.officeRepository.SearchByRate(rate)
}
