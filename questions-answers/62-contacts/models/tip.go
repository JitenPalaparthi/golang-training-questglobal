package models

type Tip struct {
	ID           uint64     `json:"id"` // automatically taken as primary key
	TipText      string     `json:"tipText" gorm:"column:tipText"`
	TechnologyID uint64     `json:"technologyID" gorm:"column:technologyID"`
	Technology   Technology `json:"technology" gorm:"foreignKey:technologyID"`
	Status       string     `json:"status"`
	LastModified int64      `json:"lastModified" gorm:"column:lastModified"`
}

type Technology struct {
	ID           uint64 `json:"id"` // automatically taken as primary key
	Tech         string `json:"tech"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified"`
}

type Userinterests struct {
	ID uint64 `json:"id"` // automatically taken as primary key

	UserID uint64 `json:"userID" gorm:"column:userID"`
	User   User   `json:"user" gorm:"foreignKey:userID"`

	TechnologyID uint64     `json:"technologyID" gorm:"column:technologyID"`
	Technology   Technology `json:"technology" gorm:"foreignKey:technologyID"`
	Status       string     `json:"status"`
	LastModified int64      `json:"lastModified" gorm:"column:lastModified"`
}

type Testmaster struct {
	ID            uint64 `json:"id"` // automatically taken as primary key
	Name          string `json:"name"`
	Description   string `json:"description"`
	NoOfQuestions uint8  `json:"noOfQuestions" gorm:"column:noOfQuestions"`
	Duration      int16  `json:"duration"`
	Status        string `json:"status"`
	LastModified  int64  `json:"lastModified" gorm:"column:lastModified"`
}

type Questionmaster struct {
	ID           uint64     `json:"id"` // automatically taken as primary key
	TestmasterID uint64     `json:"testmasterID" gorm:"column:testmasterID"`
	Testmaster   Testmaster `json:"testMaster" gorm:"foreignKey:testmasterID"`
	Question     string     `json:"question"`
	Points       uint8      `json:"points"`
	Status       string     `json:"status"`
	LastModified int64      `json:"lastModified" gorm:"column:lastModified"`
}

type Answermaster struct {
	ID uint64 `json:"id"` // automatically taken as primary key

	TestmasterID uint64     `json:"testmasterID" gorm:"column:testmasterID"`
	Testmaster   Testmaster `json:"testMaster" gorm:"foreignKey:testmasterID"`

	QuestionmasterID uint64         `json:"questionmasterID" gorm:"column:questionmasterID"`
	Questionmaster   Questionmaster `json:"questionmaster" gorm:"foreignKey:questionmasterID"`

	Option       string `json:"option"`
	IsAnswer     bool   `json:"isAnswer" gorm:"column:isAnswer"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified"`
}

type Usertest struct {
	ID     uint64 `json:"id"` // automatically taken as primary key
	TestID uint64 `json:"testID" gorm:"column:testID"`

	UserID uint64 `json:"userID" gorm:"column:userID"`
	User   User   `json:"user" gorm:"foreignKey:userID"`

	TestmasterID uint64     `json:"testmasterID" gorm:"column:testmasterID"`
	TestMaster   Testmaster `json:"testMaster" gorm:"foreignKey:testmasterID"`

	QuestionmasterID uint64         `json:"questionmasterID" gorm:"column:questionmasterID"`
	Questionmaster   Questionmaster `json:"questionmaster" gorm:"foreignKey:questionmasterID"`

	AnswermasterID uint64       `json:"answermasterID" gorm:"column:answermasterID"`
	Answermaster   Answermaster `json:"answerMaster" gorm:"foreignKey:answermasterID"`

	TestOn       uint64 `json:"testOn" gorm:"column:testOn"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified"`
}
