/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-11-01 11:43
| Describe  : article
+------------------------------------------------------------------*/

package main

import (
	"time"
)

type Code struct {
	Code  int       `json:"code" db:"code"`
	Msg   string    `json:"msg" db:"msg"`
	Total int       `json:"total" db:"total"`
	Limit int       `json:"limit" db:"limit"`
	Data  []Article `json:"data" db:"data"`
}

type User struct {
	ID          string      `json:"id" db:"id"`
	LastLogin   interface{} `json:"last_login" db:"last_login"`
	Status      bool        `json:"status" db:"status"`
	Sort        int         `json:"sort" db:"sort"`
	Ctime       time.Time   `json:"ctime" db:"ctime"`
	Utime       time.Time   `json:"utime" db:"utime"`
	Username    string      `json:"username" db:"username"`
	Email       string      `json:"email" db:"email"`
	Mobile      string      `json:"mobile" db:"mobile"`
	Qq          string      `json:"qq" db:"qq"`
	Weixin      string      `json:"weixin" db:"weixin"`
	Signature   string      `json:"signature" db:"signature"`
	Gender      int         `json:"gender" db:"gender"`
	Password    string      `json:"password" db:"password"`
	IsSuperuser bool        `json:"is_superuser" db:"is_superuser"`
	IsStaff     bool        `json:"is_staff" db:"is_staff"`
	IsActive    bool        `json:"is_active" db:"is_active"`
	HeadImg     string      `json:"head_img" db:"head_img"`
	UserType    int         `json:"user_type" db:"user_type"`
	Balance     int         `json:"balance" db:"balance"`
}

type Caparent struct {
	ID       string      `json:"id" db:"id"`
	Status   bool        `json:"status" db:"status"`
	Sort     int         `json:"sort" db:"sort"`
	Ctime    time.Time   `json:"ctime" db:"ctime"`
	Utime    time.Time   `json:"utime" db:"utime"`
	Text     string      `json:"text" db:"text"`
	Glyph    int         `json:"glyph" db:"glyph"`
	RowCls   string      `json:"rowCls" db:"row_cls"`
	Block    bool        `json:"block" db:"block"`
	Link     interface{} `json:"link" db:"link"`
	Children interface{} `json:"children" db:"children"`
}

type Category struct {
	ID       string      `json:"id" db:"id"`
	Status   bool        `json:"status" db:"status"`
	Sort     int         `json:"sort" db:"sort"`
	Ctime    time.Time   `json:"ctime" db:"ctime"`
	Utime    time.Time   `json:"utime" db:"utime"`
	Text     string      `json:"text" db:"text"`
	Glyph    int         `json:"glyph" db:"glyph"`
	RowCls   string      `json:"rowCls" db:"row_cls"`
	Block    bool        `json:"block" db:"block"`
	Link     interface{} `json:"link" db:"link"`
	Children string      `json:"children" db:"children"`
}

type Thumbnail struct {
	ID         string    `json:"id" db:"id"`
	Status     bool      `json:"status" db:"status"`
	Sort       int       `json:"sort" db:"sort"`
	Ctime      time.Time `json:"ctime" db:"ctime"`
	Utime      time.Time `json:"utime" db:"utime"`
	Path       string    `json:"path" db:"path"`
	Pixel      string    `json:"pixel" db:"pixel"`
	Proportion string    `json:"proportion" db:"proportion"`
	Md5        string    `json:"md5" db:"md5"`
	Sha1       string    `json:"sha1" db:"sha1"`
	User       string    `json:"user" db:"user"`
}

type Auditor struct {
	ID          string      `json:"id" db:"id"`
	LastLogin   interface{} `json:"last_login" db:"last_login"`
	Status      bool        `json:"status" db:"status"`
	Sort        int         `json:"sort" db:"sort"`
	Ctime       time.Time   `json:"ctime" db:"ctime"`
	Utime       time.Time   `json:"utime" db:"utime"`
	Username    string      `json:"username" db:"username"`
	Email       string      `json:"email" db:"email"`
	Mobile      interface{} `json:"mobile" db:"mobile"`
	Qq          interface{} `json:"qq" db:"qq"`
	Weixin      interface{} `json:"weixin" db:"weixin"`
	Signature   interface{} `json:"signature" db:"signature"`
	Gender      int         `json:"gender" db:"gender"`
	Password    string      `json:"password" db:"password"`
	IsSuperuser bool        `json:"is_superuser" db:"is_superuser"`
	IsStaff     bool        `json:"is_staff" db:"is_staff"`
	IsActive    bool        `json:"is_active" db:"is_active"`
	HeadImg     string      `json:"head_img" db:"head_img"`
	UserType    int         `json:"user_type" db:"user_type"`
	Balance     int         `json:"balance" db:"balance"`
}

type Article struct {
	ID          string    `json:"id" db:"id"`
	Sort        int       `json:"sort" db:"sort"`
	Ctime       time.Time `json:"ctime" db:"ctime"`
	Utime       time.Time `json:"utime" db:"utime"`
	Title       string    `json:"title" db:"title"`
	Synopsis    string    `json:"synopsis" db:"synopsis"`
	Content     string    `json:"content" db:"content"`
	IsArticle   bool      `json:"is_article" db:"is_article"`
	IsResource  bool      `json:"is_resource" db:"is_resource"`
	Encryption  bool      `json:"encryption" db:"encryption"`
	Browse      int       `json:"browse" db:"browse"`
	Collects    int       `json:"collects" db:"collects"`
	Downloads   int       `json:"downloads" db:"downloads"`
	Size        string    `json:"size" db:"size"`
	Price       int       `json:"price" db:"price"`
	Status      int       `json:"status" db:"status"`
	AuditorTime time.Time `json:"auditor_time" db:"auditor_time"`
	User        User      `json:"user" db:"user"`
	Caparent    Caparent  `json:"caparent" db:"caparent"`
	Category    Category  `json:"category" db:"category"`
	Thumbnail   Thumbnail `json:"thumbnail" db:"thumbnail"`
	Auditor     Auditor   `json:"auditor" db:"auditor"`
}

