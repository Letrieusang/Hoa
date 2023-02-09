package Model

type AdviseRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

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

type ConfessionRequest struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Content string `json:"content"`
}

func (Confession) TableName() string { return "confessions" }
func (Advise) TableName() string     { return "advises" }

func (c AdviseRequest) ConvertToAdvise() Advise {
	return Advise{
		Name:    c.Name,
		Content: c.Content,
	}
}

func (c ConfessionRequest) ConvertToConfess() Confession {
	return Confession{
		Name:    c.Name,
		Content: c.Content,
		Contact: c.Contact,
	}
}
