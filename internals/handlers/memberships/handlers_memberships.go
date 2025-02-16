package memberships

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//? Handler adalah struct yang meng-embed *gin.Engine.
//* Artinya, Handler memiliki semua metode dari gin.Engine seperti GET, POST, Group, dll.
// Dengan menggunakan *gin.Engine, Handler dapat langsung mengakses instance router Gin
type Handler struct {
	*gin.Engine
	db *gorm.DB
}

//? Fungsi ini menerima *gin.Engine sebagai parameter dan mengembalikannya dalam *Handler.
// Ini berguna agar instance gin.Engine yang sudah dibuat di main.go dapat digunakan kembali di dalam Handler.
// instance gin.Engine yang dibuat di main.go


//? api adalah parameter fungsi di NewHandler(api *gin.Engine), yang berarti api adalah variabel bertipe pointer ke gin.Engine
func NewHandler(api *gin.Engine, db *gorm.DB) *Handler {
	// Bagian ini adalah cara untuk membuat dan mengembalikan pointer ke struct Handler.
	// Ini mengembalikan alamat memori dari struct Handler, jadi kita bisa menggunakannya tanpa membuat salinan baru.
	// Artinya kita membuat struct Handler baru dan mengisi field Engine dengan nilai api, yang berasal dari parameter.
	return &Handler{Engine: api, db: db}
}
 

func (h *Handler) RegisterRoutes() {

	//  Membuat grup route dengan prefix /memberships.
	route:= h.Engine.Group("/memberships")
	// bisa dengan route := h.Group("/memberships")

	route.GET("/ping", h.Ping)
}



//* Analoginya
// type Mobil struct {
// 	Merek string
// }

// Membuat struct biasa
// mobil1 := Mobil{Merek: "Toyota"}

// Membuat pointer ke struct
// mobil2 := &Mobil{Merek: "Honda"}