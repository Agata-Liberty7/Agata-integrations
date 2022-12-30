package stats

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx/v3"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"gorm.io/gorm"
)

var (
	logChan = make(chan string)
)

func InitRoutes(f *gin.RouterGroup) {

	go func() {
		log.Println(<-logChan)
	}()

	f.GET("/:country/count", func(c *gin.Context) {
		country := c.Params.ByName("country")

		ctxDB := c.MustGet("ctxDatabase").(*gorm.DB)

		res := &[]types.MetadataRaw{}

		p, _ := getParser(country)
		supporedSections := p.AvailableSections()

		ctxDB.
			Model(&types.Stats{}).
			Select("section, count(*) as count, DATE(updated_at) as time").
			Where("country = ? AND section in ?", strings.ToUpper(country), supporedSections).
			Group("section, DATE(updated_at)").
			Find(&res)

		data := make(map[string]types.Metadata)

		// fill data
		for _, r := range *res {
			data[r.Section] = types.Metadata{
				Count: r.Count,
				Time:  r.Time,
			}
		}

		// fill empty data
		for _, Type := range SECTIONS {
			_, ok := data[Type]
			if !ok {
				for _, r := range supporedSections {
					if r == Type {
						data[Type] = types.Metadata{Count: 0}
					}
				}
			}
		}

		c.JSON(http.StatusOK, data)
	})

	f.GET("/:country/:section", func(c *gin.Context) {
		section := c.Params.ByName("section")
		country := c.Params.ByName("country")

		ctxDB := c.MustGet("ctxDatabase").(*gorm.DB)

		v := types.StatResponses{}
		ctxDB.Model(&types.Stat{}).
			Where("section LIKE ? and country LIKE ?", section, strings.ToUpper(country)).
			Find(&v)
		c.JSON(http.StatusOK, v)
	})

	f.GET("/:country/:section/clean", func(c *gin.Context) {
		section := c.Params.ByName("section")
		country := c.Params.ByName("country")

		ctxDB := c.MustGet("ctxDatabase").(*gorm.DB)
		ctxDB.
			Where("section LIKE ? and country LIKE ?", section, strings.ToUpper(country)).
			Delete(&types.Stats{})

		c.JSON(http.StatusOK, nil)
	})

	f.GET("/:country/:section/fetch", func(c *gin.Context) {
		section := c.Params.ByName("section")
		country := c.Params.ByName("country")

		// log.Println(">> fetching", section, country)

		logChan <- ">> fetching " + section + " " + country

		// // TODO: common FetchSection
		// // value, err := fedstat.FetchSection(section)
		// value, err := fetchCountrySection(country, section)

		p, err := getParser(country)
		value, err := p.Get[section]()

		// value, err := fetchSection(section, country)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		ctxDB := c.MustGet("ctxDatabase").(*gorm.DB)

		ctxDB.
			Where("section LIKE ? and country LIKE ?", section, strings.ToUpper(country)).
			Delete(&types.Stats{})
		ctxDB.Save(&value)

		v := types.StatResponses{}
		ctxDB.Model(&types.Stat{}).
			Where("section LIKE ? and country LIKE ?", section, strings.ToUpper(country)).
			Find(&v)

		// c.JSON(http.StatusOK, v)
		c.JSON(http.StatusOK, v)
	})

	// Manual input
	f.POST("/:country/:section/manual", func(c *gin.Context) {
		section := c.Params.ByName("section")
		country := c.Params.ByName("country")

		log.Println(">> manual export of ", section, country)

		sizeString := c.Request.FormValue("size")
		size, _ := strconv.Atoi(sizeString)
		file, _, err := c.Request.FormFile("upload")
		if err != nil {
			log.Println(err)
		}

		x, err := xlsx.OpenReaderAt(file, int64(size))
		if err != nil {
			log.Println(err)
		}

		sheet := x.Sheets[0]
		data := types.Stats{}

		for row := 3; row < sheet.MaxRow; row++ {
			newItem := types.Stat{
				Section: section,
				Country: strings.ToUpper(country),
			}

			cell, err := sheet.Cell(row, 0)
			if err == nil {
				newItem.Timestamp = cell.String()
			}

			cell, err = sheet.Cell(row, 1)
			if err == nil {
				newItem.Location = cell.String()
			}

			cell, err = sheet.Cell(row, 2)
			if err == nil {
				newItem.Name = cell.String()
			}

			cell, err = sheet.Cell(row, 3)
			if err == nil {
				newItem.Value = cell.String()
			}

			cell, err = sheet.Cell(row, 4)
			if err == nil {
				newItem.Unit = cell.String()
			}

			data = append(data, newItem)
		}

		ctxDB := c.MustGet("ctxDatabase").(*gorm.DB)
		ctxDB.Save(&data)

		v := types.StatResponses{}
		ctxDB.Model(&types.Stat{}).
			Where("section LIKE ? and country LIKE ?", section, strings.ToUpper(country)).
			Find(&v)

		// value, err := fetchSection(section, country)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, v)
	})
}
