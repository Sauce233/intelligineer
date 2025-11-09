package main

import "time"

var Tables = []any{
	new(Client),
	new(ClientContact),
	new(Attachment),
	new(ProjectStatus),
	new(Project),
	new(InvoiceStatus),
	new(Invoice),
	new(TaskStatus),
	new(Task),
	new(Supplier),
	new(SupplierContact),
	new(Category),
	new(MaterialStatus),
	new(Material),
	new(PurchaseStatus),
	new(Purchase),
	new(PurchaseMaterial),
	new(TaskMaterial),
	new(User),
	new(AttendanceType),
	new(Attendance),
	new(UserAttendance),
	new(ProfileType),
	new(Profile),
	new(Role),
	new(UserRole),
	new(Record),
	new(Photo),
	new(ApplicationType),
	new(ApplicationStatus),
	new(Application),
}

// Color 是下列之一
// red orange amber yellow lime green emerald teal cyan sky blue
// indigo violet purple fuchsia pink rose slate gray zinc neutral stone

type Client struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	Name      string    `gorm:"type:VARCHAR(100);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:客户名称"`
	Address   string    `gorm:"type:VARCHAR(255);not null;comment:地址"`
	Profile   string    `gorm:"type:TEXT;comment:簡介"`
}

type ClientContact struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	Name      string    `gorm:"type:VARCHAR(100);not null;comment:姓名"`
	Phone     string    `gorm:"type:VARCHAR(20);not null;comment:电话号码"`
	Email     string    `gorm:"type:VARCHAR(100);not null;comment:邮箱"`
	Profile   string    `gorm:"type:TEXT;not null;comment:簡介"`
	ClientID  uint      `gorm:"not null;index;comment:客户标识"`
	Client    Client    `gorm:"constraint:OnDelete:CASCADE"`
}

type Attachment struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	Name      string    `gorm:"type:VARCHAR(100);not null;comment:名称"`
	Filename  string    `gorm:"type:VARCHAR(255);not null;comment:文件名"`
	ProjectID uint      `gorm:"not null;index;comment:项目标识"`
	Project   Project   `gorm:"constraint:OnDelete:CASCADE"`
}

type ProjectStatus struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

type Project struct {
	ID        uint          `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time     `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time     `gorm:"not null;comment:更新时间"`
	Name      string        `gorm:"type:VARCHAR(191);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:名称"`
	ClientID  uint          `gorm:"not null;index;comment:客户标识"`
	Client    Client        `gorm:"constraint:OnDelete:RESTRICT"`
	Profile   string        `gorm:"type:TEXT;not null;comment:簡介"`
	Offer     float64       `gorm:"DECIMAL(10,2);not null;comment:报价"`
	Budget    float64       `gorm:"DECIMAL(10,2);not null;comment:成本"`
	StatusID  uint          `gorm:"not null;index;comment:状态标识"`
	Status    ProjectStatus `gorm:"constraint:OnDelete:RESTRICT"`
}

type InvoiceStatus struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

type Invoice struct {
	ID        uint          `gorm:"primarykey;comment:发票ID"`
	CreatedAt time.Time     `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time     `gorm:"not null;comment:更新时间"`
	ProjectID uint          `gorm:"not null;comment:项目标识"`
	Project   Project       `gorm:"constraint:OnDelete:CASCADE"`
	Name      string        `gorm:"type:VARCHAR(191);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:名称"`
	Number    string        `gorm:"type:VARCHAR(255);not null;comment:号码"`
	Profile   string        `gorm:"type:TEXT;not null;comment:簡介"`
	Amount    float64       `gorm:"type:decimal(10,2);comment:金额"`
	IssueDate time.Time     `gorm:"not null;comment:开具时间"`
	DueDate   time.Time     `gorm:"not null;comment:到期时间"`
	StatusID  uint          `gorm:"not null;index;comment:状态标识"`
	Status    InvoiceStatus `gorm:"constraint:OnDelete:RESTRICT"`
}

// pending, in_progress, completed
type TaskStatus struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

// start_time 和 end_time 如果未开始或未结束，那么为预估时间
type Task struct {
	ID        uint       `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time  `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time  `gorm:"not null;comment:更新时间"`
	Name      string     `gorm:"type:VARCHAR(100);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:名称"`
	Profile   string     `gorm:"type:TEXT;not null;comment:簡介"`
	ProjectID uint       `gorm:"not null;index;comment:项目标识"`
	Project   Project    `gorm:"constraint:OnDelete:RESTRICT"`
	StartTime time.Time  `gorm:"not null;comment:开始时间"`
	EndTime   time.Time  `gorm:"not null;comment:结束时间"`
	StatusID  uint       `gorm:"not null;index;comment:状态标识"`
	Status    TaskStatus `gorm:"constraint:OnDelete:RESTRICT"`
}

type Supplier struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	Name      string    `gorm:"type:VARCHAR(255);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:名称"`
	Address   string    `gorm:"type:varchar(255);comment:地址"`
	Profile   string    `gorm:"type:TEXT;comment:簡介"`
}

type SupplierContact struct {
	ID         uint      `gorm:"primarykey;comment:标识"`
	CreatedAt  time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt  time.Time `gorm:"not null;comment:更新时间"`
	Name       string    `gorm:"type:VARCHAR(100);not null;comment:姓名"`
	Phone      string    `gorm:"type:VARCHAR(20);not null;comment:电话号码"`
	Email      string    `gorm:"type:VARCHAR(100);not null;comment:邮箱"`
	Profile    string    `gorm:"type:TEXT;not null;comment:簡介"`
	SupplierID uint      `gorm:"not null;index;comment:供应商标识"`
	Supplier   Supplier  `gorm:"constraint:OnDelete:CASCADE"`
}

type Category struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	Name      string    `gorm:"type:VARCHAR(50);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:名称"`
	Profile   string    `gorm:"type:TEXT;not null;comment:簡介"`
}

type MaterialStatus struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

type Material struct {
	ID         uint           `gorm:"primarykey;comment:标识"`
	CreatedAt  time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt  time.Time      `gorm:"not null;comment:更新时间"`
	CategoryID uint           `gorm:"not null;index;comment:类型标识"`
	Category   Category       `gorm:"constraint:OnDelete:RESTRICT"`
	StatusID   uint           `gorm:"not null;index;comment:状态标识"`
	Status     MaterialStatus `gorm:"constraint:OnDelete:RESTRICT"`
	Name       string         `gorm:"type:VARCHAR(100);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:物料名"`
	Quantity   float64        `gorm:"type:DECIMAL(10,2);not null;comment:数量"`
	Unit       string         `gorm:"type:VARCHAR(10);not null;comment:单位"`
	Profile    string         `gorm:"type:VARCHAR(255);not null;comment:簡介"`
}

type PurchaseStatus struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

type Purchase struct {
	ID         uint           `gorm:"primarykey;comment:ID"`
	CreatedAt  time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt  time.Time      `gorm:"not null;comment:更新时间"`
	Name       string         `gorm:"type:VARCHAR(192);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:名称"`
	Profile    string         `gorm:"type:TEXT;not null;comment:簡介"`
	SupplierID uint           `gorm:"not null;index;comment:供应商标识"`
	Supplier   Supplier       `gorm:"constraint:OnDelete:CASCADE"`
	StatusID   uint           `gorm:"not null;index;comment:状态标识"`
	Status     PurchaseStatus `gorm:"constraint:OnDelete:RESTRICT"`
}

type PurchaseMaterial struct {
	PurchaseID uint     `gorm:"primarykey;comment:采购标识"`
	Purchase   Purchase `gorm:"constraint:OnDelete:CASCADE"`
	MaterialID uint     `gorm:"primarykey;comment:材料标识"`
	Material   Material `gorm:"constraint:OnDelete:CASCADE"`
	Quantity   float64  `gorm:"type:DECIMAL(10,2);not null;comment:数量"`
	Price      float64  `gorm:"type:DECIMAL(10,2);not null;comment:单价"`
}

// Quantity 表消耗掉的材料数量
// 如果改变了材料的状态，在 Notes 标注，例如“造成起重机轻度损坏”
type TaskMaterial struct {
	TaskID     uint     `gorm:"primarykey;comment:任务标识"`
	Task       Task     `gorm:"constraint:OnDelete:CASCADE"`
	MaterialID uint     `gorm:"primarykey;comment:材料标识"`
	Material   Material `gorm:"constraint:OnDelete:CASCADE"`
	Quantity   float64  `gorm:"type:DECIMAL(10,2);not null;comment:数量"`
	Profile    string   `gorm:"type:VARCHAR(255);not null;comment:簡介"`
}

type User struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	Password  string    `gorm:"type:CHAR(60);not null;comment:密码Hash"`
	Name      string    `gorm:"type:varchar(100);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:姓名"`
	Email     string    `gorm:"type:varchar(128);unique;comment:電郵"`
	Phone     string    `gorm:"type:VARCHAR(20);not null;comment:手機號"`
	Profile   string    `gorm:"type:TEXT;not null;comment:簡介"`
}

// clock in, clock out, overtime, training, others
type AttendanceType struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

type Attendance struct {
	ID        uint           `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"not null;comment:更新时间"`
	Name      string         `gorm:"type:VARCHAR(50);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:名称"`
	Location  string         `gorm:"type:VARCHAR(255);not null;comment:地点描述"`
	TypeID    uint           `gorm:"not null;index;comment:类型标识"`
	Type      AttendanceType `gorm:"constraint:OnDelete:RESTRICT"`
	Free      bool           `gorm:"not null;comment:自由打卡模式"`
	StartTime time.Time      `gorm:"not null;comment:开始时间"`
	EndTime   time.Time      `gorm:"not null;comment:结束时间"`
}

// pending, present, excused
type UserAttendanceStatus struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

type UserAttendance struct {
	UserID       uint                 `gorm:"primarykey;comment:用户标识"`
	User         User                 `gorm:"constraint:OnDelete:CASCADE"`
	AttendanceID uint                 `gorm:"primarykey;comment:签到标识"`
	Attendance   Attendance           `gorm:"constraint:OnDelete:CASCADE"`
	Time         *time.Time           `gorm:"comment:签到时间"`
	StatusID     uint                 `gorm:"not null;index;comment:状态标识"`
	Status       UserAttendanceStatus `gorm:"constraint:OnDelete:RESTRICT"`
}

type ProfileType struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:用戶標識"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名稱"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:顏色"`
}

type Profile struct {
	ID        uint        `gorm:"primarykey;comment:標識"`
	CreatedAt time.Time   `gorm:"not null;comment:創建時間"`
	UpdatedAt time.Time   `gorm:"not null;comment:更新時間"`
	Name      string      `gorm:"type:VARCHAR(50);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:名稱"`
	TypeID    uint        `gorm:"not null;index;comment:類型標識"`
	Type      ProfileType `gorm:"constraint:OnDelete:RESTRICT"`
	UserID    uint        `gorm:"not null;index;comment:用戶標識"`
	User      User        `gorm:"constraint:OnDelete:CASCADE"`
	Profile   string      `gorm:"type:VARCHAR(255);not null;comment:簡介"`
	Filename  string      `gorm:"type:VARCHAR(255);not null;comment:文件名"`
}

type Role struct {
	ID    uint   `gorm:"primarykey;comment:标识"`
	Name  string `gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

type UserRole struct {
	UserID uint `gorm:"primarykey;comment:用户标识"`
	User   User `gorm:"constraint:OnDelete:CASCADE"`
	RoleID uint `gorm:"primarykey;comment:角色标识"`
	Role   Role `gorm:"constraint:OnDelete:CASCADE"`
}

type Record struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:记录时间"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	TaskID    uint      `gorm:"not null;index;comment:任务标识"`
	Task      Task      `gorm:"constraint:OnDelete:CASCADE"`
	UserID    uint      `gorm:"not null;index;comment:用户标识"`
	User      User      `gorm:"constraint:OnDelete:CASCADE"`
	Title     string    `gorm:"type:VARCHAR(191);not null;comment:标题"`
	Content   string    `gorm:"type:TEXT;comment:内容"`
	Minutes   uint      `gorm:"not null;comment:时长"`
}

type Photo struct {
	ID        uint      `gorm:"primarykey;comment:标识和文件名"`
	CreatedAt time.Time `gorm:"not null;comment:上传时间"`
	RecordID  uint      `gorm:"index;not null;comment:记录标识"`
	Record    Record    `gorm:"constraint:OnDelete:CASCADE"`
}

type ApplicationType struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

// pending, approved, rejected
type ApplicationStatus struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:标识"`
	Name  string `json:"name" gorm:"type:VARCHAR(10);not null;comment:名称"`
	Color string `json:"color" gorm:"type:VARCHAR(10);not null;comment:颜色"`
}

type Application struct {
	ID          uint              `gorm:"primarykey;comment:标识"`
	CreatedAt   time.Time         `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time         `gorm:"not null;comment:更新时间"`
	ApplicantID uint              `gorm:"not null;index;comment:申请者"`
	Applicant   User              `gorm:"constraint:OnDelete:CASCADE"`
	ApproverID  *uint             `gorm:"index;comment:回复者"`
	Approver    *User             `gorm:"constraint:OnDelete:SET NULL"`
	Title       string            `gorm:"type:VARCHAR(191);not null;index:,class:FULLTEXT,option:WITH PARSER ngram;comment:标题"`
	Content     string            `gorm:"type:TEXT;not null;comment:内容"`
	Reply       string            `gorm:"type:TEXT;not null;comment:批复"`
	TypeID      uint              `gorm:"not null;index;comment:类型标识"`
	Type        ApplicationType   `gorm:"constraint:OnDelete:RESTRICT"`
	StatusID    uint              `gorm:"not null;index;comment:状态标识"`
	Status      ApplicationStatus `gorm:"constraint:OnDelete:RESTRICT"`
}
