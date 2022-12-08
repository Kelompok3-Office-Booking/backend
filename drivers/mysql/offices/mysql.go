package offices

import (
	"backend/businesses/offices"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type officeRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) offices.Repository {
	return &officeRepository{
		conn: conn,
	}
}

func (or *officeRepository) GetAll() []offices.Domain {
	var rec []Office

	or.conn.Find(&rec)
	
	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT(office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)

	officeDomain := []offices.Domain{}
	
	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		for _, fac := range officeFacilitiesPerID {
			if strconv.Itoa(int(office.ID)) == fac.Id {
				f_id := fac.F_id
				facilitesId := strings.Split(f_id, " , ")
				f_desc := fac.F_desc
				facilitesDesc := strings.Split(f_desc, " , ")
				f_slug := fac.F_slug
				facilitiesSlug := strings.Split(f_slug, " , ")

				office.FacilitiesId =  facilitesId
				office.FacilitiesDesc = facilitesDesc
				office.FacilitesSlug = facilitiesSlug
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

	var fac facilities

	querySQL = fmt.Sprintf("SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` WHERE `offices`.`id` = %s", id)

	or.conn.Raw(querySQL).Scan(&fac)

	f_id := fac.F_id
	facilitesId := strings.Split(f_id, " , ")
	f_desc := fac.F_desc
	facilitesDesc := strings.Split(f_desc, " , ")
	f_slug := fac.F_slug
	facilitiesSlug := strings.Split(f_slug, " , ")

	office.FacilitiesId =  facilitesId
	office.FacilitiesDesc = facilitesDesc
	office.FacilitesSlug = facilitiesSlug

	return office.ToDomain()
}

func (or *officeRepository) Create(officeDomain *offices.Domain) offices.Domain {
	var result *gorm.DB

	rec := FromDomain(officeDomain)

	facilitiesIdList := []int{}

	for _, v := range rec.FacilitiesId {
		id, _ := strconv.Atoi(v)
		facilitiesIdList = append(facilitiesIdList, id)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(facilitiesIdList)))

	err := or.conn.Transaction(func(tx *gorm.DB) error {
		result = tx.Create(&rec)
		result.Last(&rec)
		
		// insert to pivot table `office_images`
		for _, v := range rec.Images {
			querySQL := fmt.Sprintf("INSERT INTO `office_images`(`url`, `office_id`) VALUES ('%s', '%s')", v, strconv.Itoa(int(rec.ID)))
			
			if err := tx.Table("office_images").Exec(querySQL).Error; err != nil {
				return err
			}
		}

		// insert to pivot table `office_facilities`
		for _, v := range facilitiesIdList {
			querySQL := fmt.Sprintf("INSERT INTO `office_facilities`(`facilities_id`, `office_id`) VALUES ('%d','%d')", v, rec.ID)
			if err := tx.Table("office_facilities").Exec(querySQL).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		rec.ID = 0
		return rec.ToDomain()
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
	updatedOffice.Price = officeDomain.Price
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

	if office.ID == 0 {
		return false
	}

	deletedOffice := FromDomain(&office)
	result := or.conn.Delete(&deletedOffice)

	if result.RowsAffected == 0 {
		return false
	}

	queryDeleteImgs := fmt.Sprintf("DELETE FROM `office_images` WHERE `office_id` = '%d'", deletedOffice.ID)
	or.conn.Table("office_images").Exec(queryDeleteImgs)

	queryDeleteFac := fmt.Sprintf("DELETE FROM `office_facilities` WHERE `office_id` = '%d'", deletedOffice.ID)
	or.conn.Table("office_facilities").Exec(queryDeleteFac)

	return true
}

func (or *officeRepository) SearchByCity(city string) []offices.Domain {
	var rec []Office

	or.conn.Find(&rec, "city = ?", city)

	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT(office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		for _, fac := range officeFacilitiesPerID {
			if strconv.Itoa(int(office.ID)) == fac.Id {
				f_id := fac.F_id
				facilitesId := strings.Split(f_id, " , ")
				f_desc := fac.F_desc
				facilitesDesc := strings.Split(f_desc, " , ")
				f_slug := fac.F_slug
				facilitiesSlug := strings.Split(f_slug, " , ")

				office.FacilitiesId =  facilitesId
				office.FacilitiesDesc = facilitesDesc
				office.FacilitesSlug = facilitiesSlug
			}
		}

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

	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT( office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		for _, fac := range officeFacilitiesPerID {
			if strconv.Itoa(int(office.ID)) == fac.Id {
				f_id := fac.F_id
				facilitesId := strings.Split(f_id, " , ")
				f_desc := fac.F_desc
				facilitesDesc := strings.Split(f_desc, " , ")
				f_slug := fac.F_slug
				facilitiesSlug := strings.Split(f_slug, " , ")

				office.FacilitiesId =  facilitesId
				office.FacilitiesDesc = facilitesDesc
				office.FacilitesSlug = facilitiesSlug
			}
		}

		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) SearchByTitle(title string) []offices.Domain {
	var rec []Office

	or.conn.Find(&rec, "title = ?", title)

	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT(office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		for _, fac := range officeFacilitiesPerID {
			if strconv.Itoa(int(office.ID)) == fac.Id {
				f_id := fac.F_id
				facilitesId := strings.Split(f_id, " , ")
				f_desc := fac.F_desc
				facilitesDesc := strings.Split(f_desc, " , ")
				f_slug := fac.F_slug
				facilitiesSlug := strings.Split(f_slug, " , ")

				office.FacilitiesId =  facilitesId
				office.FacilitiesDesc = facilitesDesc
				office.FacilitesSlug = facilitiesSlug
			}
		}

		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) GetOffices() []offices.Domain {
	var rec []Office

	or.conn.Find(&rec, "office_type = ?", "office")

	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT( office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		for _, fac := range officeFacilitiesPerID {
			if strconv.Itoa(int(office.ID)) == fac.Id {
				f_id := fac.F_id
				facilitesId := strings.Split(f_id, " , ")
				f_desc := fac.F_desc
				facilitesDesc := strings.Split(f_desc, " , ")
				f_slug := fac.F_slug
				facilitiesSlug := strings.Split(f_slug, " , ")

				office.FacilitiesId =  facilitesId
				office.FacilitiesDesc = facilitesDesc
				office.FacilitesSlug = facilitiesSlug
			}
		}

		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) GetCoworkingSpace() []offices.Domain {
	var rec []Office

	or.conn.Find(&rec, "office_type = ?", "coworking space")

	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT( office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		for _, fac := range officeFacilitiesPerID {
			if strconv.Itoa(int(office.ID)) == fac.Id {
				f_id := fac.F_id
				facilitesId := strings.Split(f_id, " , ")
				f_desc := fac.F_desc
				facilitesDesc := strings.Split(f_desc, " , ")
				f_slug := fac.F_slug
				facilitiesSlug := strings.Split(f_slug, " , ")

				office.FacilitiesId =  facilitesId
				office.FacilitiesDesc = facilitesDesc
				office.FacilitesSlug = facilitiesSlug
			}
		}

		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) GetMeetingRooms() []offices.Domain {
	var rec []Office

	or.conn.Find(&rec, "office_type = ?", "meeting room")

	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT( office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		for _, fac := range officeFacilitiesPerID {
			if strconv.Itoa(int(office.ID)) == fac.Id {
				f_id := fac.F_id
				facilitesId := strings.Split(f_id, " , ")
				f_desc := fac.F_desc
				facilitesDesc := strings.Split(f_desc, " , ")
				f_slug := fac.F_slug
				facilitiesSlug := strings.Split(f_slug, " , ")

				office.FacilitiesId =  facilitesId
				office.FacilitiesDesc = facilitesDesc
				office.FacilitesSlug = facilitiesSlug
			}
		}

		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) GetRecommendation() []offices.Domain {
	var rec []Office

	or.conn.Order("rate desc, title, description").Find(&rec)
	
	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT(office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)

	officeDomain := []offices.Domain{}
	
	for _, office := range rec {
		for _, v := range imgsUrlPerID {
			if strconv.Itoa(int(office.ID)) == v.Id {
				url := v.Images
				img := strings.Split(url, " , ")
				office.Images = img
			}
		}

		for _, fac := range officeFacilitiesPerID {
			if strconv.Itoa(int(office.ID)) == fac.Id {
				f_id := fac.F_id
				facilitesId := strings.Split(f_id, " , ")
				f_desc := fac.F_desc
				facilitesDesc := strings.Split(f_desc, " , ")
				f_slug := fac.F_slug
				facilitiesSlug := strings.Split(f_slug, " , ")

				office.FacilitiesId =  facilitesId
				office.FacilitiesDesc = facilitesDesc
				office.FacilitesSlug = facilitiesSlug
			}
		}

		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) GetNearest(lat string, long string) []offices.Domain {
	var rec []Office

	or.conn.Find(&rec)
	
	var imgsUrlPerID []imgs

	queryGetImgs := "SELECT `offices`.`id`, GROUP_CONCAT( office_images.url ORDER BY office_images.id SEPARATOR ' , ') AS images FROM offices INNER JOIN office_images on offices.id = office_images.office_id GROUP BY offices.id"
	or.conn.Raw(queryGetImgs).Scan(&imgsUrlPerID)

	var officeFacilitiesPerID []facilities
	queryGetFacilities := "SELECT `offices`.`id`,GROUP_CONCAT(`office_facilities`.`facilities_id` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_id,GROUP_CONCAT(`facilities`.`description` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_desc,GROUP_CONCAT(`facilities`.`slug` ORDER BY `office_facilities`.`facilities_id` SEPARATOR ' , ') AS f_slug FROM `offices` INNER JOIN `office_facilities` ON `offices`.`id`=`office_facilities`.`office_id` INNER JOIN `facilities` ON `office_facilities`.`facilities_id`=`facilities`.`id` GROUP BY `offices`.`id`"
	or.conn.Raw(queryGetFacilities).Scan(&officeFacilitiesPerID)
	
	// find nearest logic here
	var distance []distance
	queryGetDistance := fmt.Sprintf("SELECT `offices`.`id`, CAST((SQRT(POW(69.1 * (`lat` - %s), 2) + POW(69.1 * (%s -`lng`) * COS(`lat` / 57.3), 2))) AS decimal(16,2)) AS distance FROM `offices` ORDER BY distance,`offices`.`id`;", lat, long)
	or.conn.Raw(queryGetDistance).Scan(&distance)

	officeDomain := []offices.Domain{}

	for _, d := range distance {
		for _, office := range rec {
			for _, v := range imgsUrlPerID {
				if strconv.Itoa(int(office.ID)) == v.Id {
					url := v.Images
					img := strings.Split(url, " , ")
					office.Images = img
				}
			}
	
			for _, fac := range officeFacilitiesPerID {
				if strconv.Itoa(int(office.ID)) == fac.Id {
					f_id := fac.F_id
					facilitesId := strings.Split(f_id, " , ")
					f_desc := fac.F_desc
					facilitesDesc := strings.Split(f_desc, " , ")
					f_slug := fac.F_slug
					facilitiesSlug := strings.Split(f_slug, " , ")
	
					office.FacilitiesId =  facilitesId
					office.FacilitiesDesc = facilitesDesc
					office.FacilitesSlug = facilitiesSlug
				}
			}

			if strconv.Itoa(int(office.ID)) == d.Id {
				office.Distance = d.Distance
				officeDomain = append(officeDomain, office.ToDomain())
			}
		}
	}

	return officeDomain
}