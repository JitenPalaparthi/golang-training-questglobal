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

type UserInterests struct {
	ID uint64 `json:"id"` // automatically taken as primary key

	UserID uint64 `json:"userID" gorm:"column:userID"`
	User   User   `json:"user" gorm:"foreignKey:userID"`

	TechnologyID uint64     `json:"technologyID" gorm:"column:technologyID"`
	Technology   Technology `json:"technology" gorm:"foreignKey:technologyID"`
	Status       string     `json:"status"`
	LastModified int64      `json:"lastModified" gorm:"column:lastModified"`
}

type TestMaster struct {
	ID            uint64 `json:"id"` // automatically taken as primary key
	Name          string `json:"name"`
	Description   string `json:"description"`
	NoOfQuestions uint8  `json:"noOfQuestions" gorm:"column:noOfQuestions"`
	Duration      int16  `json:"duration"`
	Status        string `json:"status"`
	LastModified  int64  `json:"lastModified" gorm:"column:lastModified"`
}

type QuestionMaster struct {
	ID           uint64     `json:"id"` // automatically taken as primary key
	TestMasterID uint64     `json:"testMasterID" gorm:"column:testMasterID"`
	TestMaster   TestMaster `json:"testMaster" gorm:"foreignKey:testMasterID"`
	Question     string     `json:"question"`
	Points       uint8      `json:"points"`
	Status       string     `json:"status"`
	LastModified int64      `json:"lastModified" gorm:"column:lastModified"`
}

type AnswerMaster struct {
	ID uint64 `json:"id"` // automatically taken as primary key

	TestMasterID uint64     `json:"testMasterID" gorm:"column:testMasterID"`
	TestMaster   TestMaster `json:"testMaster" gorm:"foreignKey:testMasterID"`

	QuestionMasterID uint64         `json:"questionMasterID" gorm:"column:questionMasterID"`
	QuestionMaster   QuestionMaster `json:"questionMaster" gorm:"foreignKey:questionMasterID"`

	Option       string `json:"option"`
	IsAnswer     bool   `json:"isAnswer" gorm:"column:isAnswer"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified"`
}

type UserTest struct {
	ID     uint64 `json:"id"` // automatically taken as primary key
	TestID uint64 `json:"testID" gorm:"column:testID"`

	UserID uint64 `json:"userID" gorm:"column:userID"`
	User   User   `json:"user" gorm:"foreignKey:userID"`

	TestMasterID uint64     `json:"testMasterID" gorm:"column:testMasterID"`
	TestMaster   TestMaster `json:"testMaster" gorm:"foreignKey:testMasterID"`

	QuestionMasterID uint64         `json:"questionMasterID" gorm:"column:questionMasterID"`
	QuestionMaster   QuestionMaster `json:"questionMaster" gorm:"foreignKey:questionMasterID"`

	AnswerMasterID uint64       `json:"answerMasterID" gorm:"column:answerMasterID"`
	AnswerMaster   AnswerMaster `json:"answerMaster" gorm:"foreignKey:answerMasterID"`

	TestOn       uint64 `json:"testOn" gorm:"column:testOn"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified"`
}
