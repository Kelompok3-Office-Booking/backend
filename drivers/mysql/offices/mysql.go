package offices

import (
	"backend/businesses/offices"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type officeRepository struct {
	conn *gorm.DB
}

type imgsInterface struct {
	Id string
	Images string
}

func NewMySQLRepository(conn *gorm.DB) offices.Repository {
	return &officeRepository{
		conn: conn,
	}
}

func (or *officeRepository) GetAll() []offices.Domain {
	var rec []Office

	or.conn.Find(&rec)
	
	var imgsUrlPerID []imgsInterface

	or.conn.Raw("SELECT `offices`.`id`, GROUP_CONCAT( office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id").Scan(&imgsUrlPerID)

	officeDomain := []offices.Domain{}
	
	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) GetByID(id string) offices.Domain {
	var office Office

	or.conn.First(&office, "id = ?", id)
	
	var imagesString string
	
	// get office images
	querySQL := fmt.Sprintf("SELECT GROUP_CONCAT(office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id WHERE `offices`.`id` = %s GROUP BY offices.id", id)

	or.conn.Raw(querySQL).Scan(&imagesString)

	img := strings.Split(imagesString, " , ")
	office.Images = img

	return office.ToDomain()
}

func (or *officeRepository) Create(officeDomain *offices.Domain) offices.Domain {
	rec := FromDomain(officeDomain)

	result := or.conn.Create(&rec)

	result.Last(&rec)

	// insert to pivot table `office_images`
	for _, v := range rec.Images {
		querySQL := fmt.Sprintf("INSERT INTO `office_images`(`url`, `office_id`) VALUES ('%s', '%s')", v, strconv.Itoa(int(rec.ID)))
		or.conn.Table("office_images").Exec(querySQL)
	}

	return rec.ToDomain()
}

func (or *officeRepository) Update(id string, officeDomain *offices.Domain) offices.Domain {
	var office offices.Domain = or.GetByID(id)

	updatedOffice := FromDomain(&office)

	updatedOffice.Title = officeDomain.Title
	updatedOffice.Description = officeDomain.Description
	updatedOffice.OfficeType = officeDomain.OfficeType
	updatedOffice.OfficeLength = officeDomain.OfficeLength
	updatedOffice.PricePerHour = officeDomain.PricePerHour
	updatedOffice.OpenHour = officeDomain.OpenHour
	updatedOffice.CloseHour = officeDomain.CloseHour
	updatedOffice.Lat = officeDomain.Lat
	updatedOffice.Lng = officeDomain.Lng
	updatedOffice.Accommodate = officeDomain.Accommodate
	updatedOffice.WorkingDesk = officeDomain.WorkingDesk
	updatedOffice.MeetingRoom = officeDomain.MeetingRoom
	updatedOffice.PrivateRoom = officeDomain.PrivateRoom
	updatedOffice.City = officeDomain.City
	updatedOffice.District = officeDomain.District
	updatedOffice.Address = officeDomain.Address
	updatedOffice.Rate = officeDomain.Rate

	or.conn.Save(&updatedOffice)

	if len(officeDomain.Images) != 0 {
		queryDeleteImgs := fmt.Sprintf("DELETE FROM `office_images` WHERE `office_id` = %s", id)

		or.conn.Table("office_images").Exec(queryDeleteImgs)
	
		// insert to pivot table `office_images`
		for _, v := range officeDomain.Images {
			querySQL := fmt.Sprintf("INSERT INTO `office_images`(`url`, `office_id`) VALUES ('%s', '%s')", v, id)
			or.conn.Table("office_images").Exec(querySQL)
		}
	}

	return updatedOffice.ToDomain()
}

func (or *officeRepository) Delete(id string) bool {
	var office offices.Domain = or.GetByID(id)

	queryDeleteImgs := fmt.Sprintf("DELETE FROM `office_images` WHERE `office_id` = %s", id)

	resultDeletedImgs := or.conn.Table("office_images").Raw(queryDeleteImgs)

	if resultDeletedImgs.RowsAffected == 0 {
		return false
	} else {
		deletedOffice := FromDomain(&office)
		resultDeleteOffice := or.conn.Delete(&deletedOffice)
		
		if resultDeleteOffice.RowsAffected == 0 {
			return false
		} 
	}

	return true
}

func (or *officeRepository) SearchByCity(city string) []offices.Domain {
	var rec []Office

	or.conn.Find(&rec, "city = ?", city)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) SearchByRate(rate string) []offices.Domain {
	var rec []Office

	intRate, _ := strconv.Atoi(rate)

	if intRate == 5 {
		or.conn.Find(&rec, "rate = ?", rate)
	} else {
		or.conn.Where("rate >= ? AND rate < ?", rate, intRate+1).Order("rate desc, title").Find(&rec)
	}

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) SearchByTitle(title string) offices.Domain {
	var rec Office

	or.conn.First(&rec, "title = ?", title)

	return rec.ToDomain()
}
