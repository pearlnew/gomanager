package models

import (
	"time"
)

type User struct {
	UserID     int64  `json:"userid" xorm:"int(11) pk notnull autoincr unique 'UserID' comment('用户ID')"`
	UnionID    string `json:"unionid" xorm:"varchar(100) notnull 'UnionID' comment('微信 UNION ID')"`
	OpenID     string `json:"openid" xorm:"varchar(100) notnull unique 'OpenID' comment('微信 OPEN ID')"`
	UserName   string `json:"username" xorm:"varchar(25) null unique 'UserName' comment('用户名')"`
	NickName   string    `json:"nickName" xorm:"varchar(127) 'NickName' "`
	Sex        int8      `json:"gender" xorm:"int(2) 'Sex' " `
	HeadImgURL string    `json:"avatarUrl" xorm:"varchar(255) 'HeadImgURL' "`
	IsAdmin    int8      `json:"isAdmin" xorm:"tinyint(1) notnull default(0) 'IsAdmin' "`
	RegistTime time.Time `json:"registtime" xorm:"created 'RegistTime'"`
	Province      string    `json:"Province" xorm:"-"`
	City      string    `json:"City" xorm:"-"`
	ExpireTime time.Time `json:"expire" xorm:"-"`
}

func (u *User) TableName() string {
	return "user"
}

func UserAdd(u *User) error {
	return nil
}


func UserSearch(u *User) (*User, error) {
	ret := &User{UserName: "pearl",Sex:2, UserID:1,RegistTime: time.Now(), IsAdmin: 1, NickName: "女王", Province: "上海", City:"上海", HeadImgURL: "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIrv2ibfCvqpGu4PJuZ1hQ1ydkibO7muTBUwoM9TcTBzs8CqHicCxZcLed9nTZ1XABhTtXGfk1o1lQyQ/132"}
	return ret, nil
}

func UserList() (*[]User, error) {
	ret := []User{
		{UserName: "pearl",Sex:2, UserID:1,RegistTime: time.Now(), IsAdmin: 1, NickName: "女王", Province: "上海", City:"上海", HeadImgURL: "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIrv2ibfCvqpGu4PJuZ1hQ1ydkibO7muTBUwoM9TcTBzs8CqHicCxZcLed9nTZ1XABhTtXGfk1o1lQyQ/132"},
		{UserName: "zhang3", Sex:1, UserID:2, RegistTime: time.Now(), IsAdmin: 0, NickName: "张三", Province: "北京", City:"北京", HeadImgURL: "https://thirdwx.qlogo.cn/mmopen/vi_32/VkyFrnwibyhVPh1ejCXDZ9qHebsWBTlEn9wTuvibQGzLPEEWtIdu63uDNrBgsVQVr5RMicGuH16VwYGYWSVb2icJZg/132"},
		{UserName: "li4", Sex:1, UserID:3, RegistTime: time.Now(), IsAdmin: 0, NickName: "里斯", Province: "广东", City:"珠海", HeadImgURL: "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIrv2ibfCvqpGu4PJuZ1hQ1ydkibO7muTBUwoM9TcTBzs8CqHicCxZcLed9nTZ1XABhTtXGfk1o1lQyQ/132"},
		{UserName: "wang5", Sex:1, UserID:4, RegistTime: time.Now(), IsAdmin: 0, NickName: "王五", Province: "天津", City:"天津", HeadImgURL: "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJDdmHm6txEVuKcwEiaqHnQqR9iaKW2axhYkhHpbrH7K4cciaBTAyxWiaS2qvuSO52d0Dv3hwFPMWiaBlw/132"},
	}

	return &ret, nil
}
