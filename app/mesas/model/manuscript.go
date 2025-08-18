package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-microservice/common/basic/mysql"
	"gorm.io/gorm"
)

const STATUS_ACTIVE = 1

type (
	ManuscriptModel interface {
		GetItem(manuscriptId int64) (*Manuscript, error)
		GetItems(*ListManuscriptReq) ([]Manuscript, error)
		Create(*Manuscript) (*Manuscript, error)
		Update(manuscriptId int64, params Manuscript) (*Manuscript, error)
		Delete(manuscriptId int64) error
	}

	ListManuscriptReq struct {
		JournalId uint64
		Title     string
		TypeId    uint64
	}

	defaultmanuscriptModel struct {
		conn *gorm.DB
	}

	Manuscript struct {
		Id                           int64   `gorm:"primary_key" json:"id"`
		ProcessType                  uint64  `gorm:"process_type" json:"process_type"`                                         // 流程类型（1普通，2外部编辑）
		Uid                          uint64  `gorm:"uid" json:"uid"`                                                           // 投稿用户ID
		CoMeUid                      uint64  `gorm:"co_me_uid" json:"co_me_uid"`                                               // Co-ME用户ID
		AeUid                        uint64  `gorm:"ae_uid" json:"ae_uid"`                                                     // 助理编辑用户ID
		AcquiringUid                 uint64  `gorm:"acquiring_uid" json:"acquiring_uid"`                                       // 组稿编辑用户ID
		EditorUid                    uint64  `gorm:"editor_uid" json:"editor_uid"`                                             // editor用户ID
		JournalId                    uint64  `gorm:"journal_id" json:"journal_id"`                                             // 期刊ID
		Status                       string  `gorm:"status" json:"status"`                                                     // 状态
		Flag                         uint64  `gorm:"flag" json:"flag"`                                                         // 进度状态
		SerialId                     string  `gorm:"serial_id" json:"serial_id"`                                               // 序列号
		SerialIdRule                 uint64  `gorm:"serial_id_rule" json:"serial_id_rule"`                                     // 序列号规则 1=按年份,2=按天数'
		Title                        string  `gorm:"title" json:"title"`                                                       // 标题
		TypeId                       uint64  `gorm:"type_ids" json:"type_ids"`                                                 // 稿件类型ID
		SourceId                     uint64  `gorm:"source_id" json:"source_id"`                                               // 来源ID
		Abstract                     string  `gorm:"abstract" json:"abstract"`                                                 // 简介
		Highlight                    string  `gorm:"highlight" json:"highlight"`                                               // 高亮内容
		Keywords                     string  `gorm:"keywords" json:"keywords"`                                                 // 关键词
		CompoundPdf                  string  `gorm:"compound_pdf" json:"compound_pdf"`                                         // 附件合成PDF路径
		CompoundPdfForReview         string  `gorm:"compound_pdf_for_review" json:"compound_pdf_for_review"`                   // 审稿人附件合成PDF路径
		CompoundRevisionPdf          string  `gorm:"compound_revision_pdf" json:"compound_revision_pdf"`                       // 反修附件合成PDF路径
		CompoundRevisionPdfForReview string  `gorm:"compound_revision_pdf_for_review" json:"compound_revision_pdf_for_review"` // 反修审稿人附件合成PDF路径
		ManuscriptFile               string  `gorm:"manuscript_file" json:"manuscript_file"`                                   // 手稿
		PublicationFee               float64 `gorm:"publication_fee" json:"publication_fee"`                                   // 出版费(美元)
		IsSendApc                    uint64  `gorm:"is_send_apc" json:"is_send_apc"`                                           // 是否发送APC
		SendApcAt                    uint64  `gorm:"send_apc_at" json:"send_apc_at"`                                           // 发送APC时间
		IsSubmitLicense              uint64  `gorm:"is_submit_license" json:"is_submit_license"`                               // 是否提交授权
		OriginalManuscriptId         uint64  `gorm:"original_manuscript_id" json:"original_manuscript_id"`                     // 原稿件ID
		Evaluation                   string  `gorm:"evaluation" json:"evaluation"`                                             // 评价信息
		Video                        string  `gorm:"video" json:"video"`                                                       // 视频介绍
		SubmissionAt                 uint64  `gorm:"submission_at" json:"submission_at"`                                       // 投稿时间
		AcceptedAt                   uint64  `gorm:"accepted_at" json:"accepted_at"`                                           // 接受时间
		PublishedAt                  uint64  `gorm:"published_at" json:"published_at"`                                         // 发布时间
		OptUid                       uint64  `gorm:"opt_uid" json:"opt_uid"`
		OperatedAt                   uint64  `gorm:"operated_at" json:"operated_at"`     // 最后操作时间
		LastActedAt                  uint64  `gorm:"last_acted_at" json:"last_acted_at"` // 最后审稿操作时间
		IsSyncOrcid                  uint64  `gorm:"is_sync_orcid" json:"is_sync_orcid"` // 是否同步orcid
	}
)

func NewManuscriptModel() ManuscriptModel {
	return &defaultmanuscriptModel{
		conn: mysql.GetDB(),
	}
}

func (u Manuscript) TableName() string {
	return "manuscripts"
}

func (this defaultmanuscriptModel) GetItem(manuscriptId int64) (*Manuscript, error) {

	manuscript := &Manuscript{}

	if err := this.conn.Where(Manuscript{Id: manuscriptId}).First(&manuscript).Error; err != nil {
		return nil, err

	}

	return manuscript, nil
}

func (this defaultmanuscriptModel) GetItems(req *ListManuscriptReq) ([]Manuscript, error) {

	var manuscripts []Manuscript

	model := this.conn.Model(&Manuscript{})

	if req.TypeId != 0 {
		model = model.Where(Manuscript{TypeId: req.TypeId})
	}

	if req.Title != "" {
		model = model.Where(Manuscript{Title: req.Title})
	}

	if req.JournalId != 0 {
		model = model.Where(Manuscript{JournalId: req.JournalId})
	}

	if err := model.Find(&manuscripts).Error; err != nil {
		return nil, err
	}

	return manuscripts, nil
}

func (this defaultmanuscriptModel) Create(manuscript *Manuscript) (*Manuscript, error) {

	if err := this.conn.Create(&manuscript).Error; err != nil {
		return nil, err

	}
	return manuscript, nil
}

func (this defaultmanuscriptModel) Update(manuscriptId int64, params Manuscript) (*Manuscript, error) {

	manuscript, err := this.GetItem(manuscriptId)

	if err != nil {
		logx.Errorf("数据不存在：", err)
		return nil, err
	}

	if params.Title != "" {
		manuscript.Title = params.Title
	}

	if params.Keywords != "" {
		manuscript.Keywords = params.Keywords
	}

	if err := this.conn.Where(Manuscript{Id: manuscriptId}).Save(&manuscript).Error; err != nil {
		return nil, err
	}

	return manuscript, nil
}

func (this defaultmanuscriptModel) Delete(manuscriptId int64) error {

	manuscript, err := this.GetItem(manuscriptId)

	if err != nil {
		logx.Errorf("数据不存在：", err)
		return err
	}

	if err := this.conn.Delete(&manuscript).Error; err != nil {
		return err

	}

	return nil
}
