package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter(c *Config) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://axogc.net:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"X-Access-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/enums", WithAuthBind(c, GetEnums))
	r.GET("/user-hours/:userId/:days", WithAuthBind(c, StatUserHours))

	r.GET("/roles", WithAuthBind(c, ListRole))
	r.DELETE("/roles/:roleId", WithAuthBind(c, DeleteRole))

	r.PATCH("/client-contacts/:clientContactId", WithAuthBind(c, EditClientContact))
	r.DELETE("/client-contacts/:clientContactId", WithAuthBind(c, DeleteClientContact))

	r.GET("/categories", WithAuthBind(c, ListCategory))
	r.POST("/categories", WithAuthBind(c, AddCategory))
	r.GET("/categories/:categoryId", WithAuthBind(c, GetCategory))
	r.PATCH("/categories/:categoryId", WithAuthBind(c, EditCategory))
	r.DELETE("/categories/:categoryId", WithAuthBind(c, DeleteCategory))

	r.GET("/project-statuses", WithAuthBind(c, ListProjectStatus))

	r.GET("/projects", WithAuthBind(c, ListProject))
	r.GET("/projects/:projectId", WithAuthBind(c, GetProject))
	r.PATCH("/projects/:projectId", WithAuthBind(c, EditProject))
	r.DELETE("/projects/:projectId", WithAuthBind(c, DeleteProject))
	r.GET("/projects/:projectId/tasks", WithAuthBind(c, ListProjectTask))
	r.POST("/projects/:projectId/tasks", WithAuthBind(c, AddTask))
	r.GET("/projects/:projectId/invoices", WithAuthBind(c, ListProjectInvoice))
	r.POST("/projects/:projectId/invoices", WithAuthBind(c, AddInvoice))
	r.GET("/projects/:projectId/attachments", WithAuthBind(c, ListProjectAttachment))
	r.POST("/projects/:projectId/attachments", WithAuthBind(c, AddProjectAttachment))

	r.GET("/records", WithAuthBind(c, ListRecord))
	r.GET("/records/:recordId", WithAuthBind(c, GetRecord))
	r.PATCH("/records/:recordId", WithAuthBind(c, EditRecord))
	r.DELETE("/records/:recordId", WithAuthBind(c, DeleteRecord))

	r.GET("/clients", WithAuthBind(c, ListClient))
	r.POST("/clients", WithAuthBind(c, AddClient))
	r.GET("/clients/:clientId", WithAuthBind(c, GetClient))
	r.PATCH("/clients/:clientId", WithAuthBind(c, EditClient))
	r.DELETE("/clients/:clientId", WithAuthBind(c, DeleteClient))
	r.GET("/clients/:clientId/contacts", WithAuthBind(c, ListClientContact))
	r.POST("/clients/:clientId/contacts", WithAuthBind(c, AddClientContact))
	r.GET("/clients/:clientId/projects", WithAuthBind(c, ListClientProject))
	r.POST("/clients/:clientId/projects", WithAuthBind(c, AddProject))

	r.GET("/users", WithAuthBind(c, ListUser))
	r.POST("/users", WithAuthBind(c, AddUser))
	r.GET("/users/:userId", WithAuthBind(c, GetUser))
	r.PATCH("/users/:userId", WithAuthBind(c, EditUser))
	r.DELETE("/users/:userId", WithAuthBind(c, DeleteUser))
	r.GET("/users/:userId/records", WithAuthBind(c, ListUserRecord))
	r.GET("/users/:userId/attendances", WithAuthBind(c, ListUserAttendance))

	r.GET("/profiles", WithAuthBind(c, ListProfile))
	r.DELETE("/profiles/:profileId", WithAuthBind(c, DeleteProfile))
	r.POST("/users/:userId/profiles", WithAuthBind(c, AddProfile))

	r.GET("/task-statuses", WithAuthBind(c, ListTaskStatus))

	r.GET("/materials/:materialId/tasks", WithAuthBind(c, ListTask))
	r.GET("/tasks", WithAuthBind(c, ListTask))
	r.GET("/tasks/:taskId", WithAuthBind(c, GetTask))
	r.PATCH("/tasks/:taskId", WithAuthBind(c, EditTask))
	r.DELETE("/tasks/:taskId", WithAuthBind(c, DeleteTask))
	r.POST("/tasks/:taskId/records", WithAuthBind(c, AddRecord))
	r.GET("/tasks/:taskId/records", WithAuthBind(c, ListTaskRecord))
	r.GET("/tasks/:taskId/materials", WithAuthBind(c, ListTaskMaterial))

	r.GET("/suppliers", WithAuthBind(c, ListSupplier))
	r.POST("/suppliers", WithAuthBind(c, AddSupplier))
	r.GET("/suppliers/:supplierId", WithAuthBind(c, GetSupplier))
	r.PATCH("/suppliers/:supplierId", WithAuthBind(c, EditSupplier))
	r.DELETE("/suppliers/:supplierId", WithAuthBind(c, DeleteSupplier))
	r.GET("/suppliers/:supplierId/contacts", WithAuthBind(c, ListSupplierContact))
	r.GET("/suppliers/:supplierId/purchases", WithAuthBind(c, ListSupplierPurchase))

	r.PATCH("/supplier-contacts/:supplierContactId", WithAuthBind(c, EditSupplierContact))
	r.DELETE("/supplier-contacts/:supplierContactId", WithAuthBind(c, DeleteSupplierContact))
	r.POST("/suppliers/:supplierId/supplier-contacts", WithAuthBind(c, AddSupplierContact))

	r.POST("/photos", WithAuthBind(c, AddPhoto))
	r.DELETE("/photos/:photoId", WithAuthBind(c, DeletePhoto))

	r.GET("/application-statuses", WithAuthBind(c, ListApplicationStatus))
	r.GET("/application-types", WithAuthBind(c, ListApplicationType))

	r.GET("/applications", WithAuthBind(c, ListApplication))
	r.POST("/applications", WithAuthBind(c, AddApplication))
	r.GET("/applications/:applicationId", WithAuthBind(c, GetApplication))
	r.DELETE("/applications/:applicationId", WithAuthBind(c, DeleteApplication))
	r.PATCH("/applications/:applicationId/reply", WithAuthBind(c, ReplyApplication))
	r.PATCH("/applications/:applicationId/approve", WithAuthBind(c, ApproveApplication))
	r.PATCH("/applications/:applicationId/reject", WithAuthBind(c, RejectApplication))

	r.GET("/purchase-statuses", WithAuthBind(c, ListPurchaseStatus))

	r.GET("/purchases", WithAuthBind(c, ListPurchase))
	r.GET("/purchases/:purchaseId", WithAuthBind(c, GetPurchase))
	r.PATCH("/purchases/:purchaseId", WithAuthBind(c, EditPurchase))
	r.DELETE("/purchases/:purchaseId", WithAuthBind(c, DeletePurchase))
	r.POST("/suppliers/:supplierId/purchases", WithAuthBind(c, AddPurchase))

	r.GET("/invoice-statuses", WithAuthBind(c, ListInvoiceStatus))

	r.GET("/invoices", WithAuthBind(c, ListInvoice))
	r.GET("/invoices/:invoiceId", WithAuthBind(c, GetInvoice))
	r.PATCH("/invoices/:invoiceId", WithAuthBind(c, EditInvoice))
	r.DELETE("/invoices/:invoiceId", WithAuthBind(c, DeleteInvoice))

	r.GET("/material-statuses", WithAuthBind(c, ListMaterialStatus))

	r.GET("/materials", WithAuthBind(c, ListMaterial))
	r.GET("/materials/:materialId", WithAuthBind(c, GetMaterial))
	r.PATCH("/materials/:materialId", WithAuthBind(c, EditMaterial))
	r.DELETE("/materials/:materialId", WithAuthBind(c, DeleteMaterial))
	r.POST("/categories/:categoryId/materials", WithAuthBind(c, AddMaterial))

	r.GET("/attendance-types", WithAuthBind(c, ListAttendanceType))

	r.GET("/attendances", WithAuthBind(c, ListAttendance))
	r.GET("/attendances/:attendanceId", WithAuthBind(c, GetAttendance))
	r.GET("/attendances/:attendanceId/users", WithAuthBind(c, ListAttendanceUser))

	r.NoRoute(func(c *gin.Context) { c.JSON(404, Res("接口不存在", nil)) })

	return r
}
