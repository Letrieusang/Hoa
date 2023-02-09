package Model

type Confession struct {
	Id      int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	Name    string `gorm:"type:VARCHAR(255)" json:"name"`
	Contact string `gorm:"type:VARCHAR(255)" json:"contact"`
	Content string `gorm:"type:TEXT" json:"content"`
}

type Advise struct {
	Id      int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	Name    string `gorm:"type:VARCHAR(255)" json:"name"`
	Content string `gorm:"type:TEXT" json:"content"`
}

func (Confession) TableName() string { return "confessions" }
func (Advise) TableName() string     { return "advises" }
