package fedstat

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
// 	"gorm.io/gorm"
// )

// func InitRoutes(r *gin.Engine) {
// 	f := r.Group("/fedstat")

// 	f.GET("/:section/clean", func(c *gin.Context) {
// 		section := c.Params.ByName("section")

// 		ctxDB := c.MustGet("ctxDatabase").(*gorm.DB)
// 		ctxDB.Where("section LIKE ? and country LIKE ?", section, COUNTRY).Delete(&types.Stats{})

// 		c.JSON(http.StatusOK, nil)
// 	})

// 	f.GET("/:section/fetch", func(c *gin.Context) {
// 		section := c.Params.ByName("section")

// 		value, err := FetchSection(section)
// 		if err != nil {
// 			log.Println("err", err)
// 			c.JSON(http.StatusInternalServerError, err)
// 		}

// 		ctxDB := c.MustGet("ctxDatabase").(*gorm.DB)

// 		ctxDB.
// 			Where("section LIKE ? and country LIKE ?", section, COUNTRY).
// 			Delete(&types.Stats{})
// 		ctxDB.Save(&value)

// 		v := types.StatResponses{}
// 		ctxDB.Model(&types.Stat{}).
// 			Where("section LIKE ? and country LIKE ?", section, COUNTRY).
// 			Find(&v)

// 		c.JSON(http.StatusOK, v)
// 	})
// }
