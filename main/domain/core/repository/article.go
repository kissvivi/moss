package repository

import (
	"gorm.io/gorm"
	"moss/domain/config"
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/repository/gormx"
	"moss/domain/core/utils"
	"moss/infrastructure/general/message"
	"moss/infrastructure/persistent/db"
	"moss/infrastructure/utils/errorx"
)

var Article = new(ArticleRepo)

type ArticleRepo struct {
}

func (r *ArticleRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&entity.ArticleBase{}, &entity.ArticleDetail{})
}

func (r *ArticleRepo) Create(item *entity.Article) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := r.checkPost(tx, item); err != nil {
			return err
		}
		if err := tx.Create(&item.ArticleBase).Error; err != nil {
			return err
		}
		item.ArticleDetail.ArticleID = item.ArticleBase.ID
		return tx.Create(&item.ArticleDetail).Error
	})
}

func (r *ArticleRepo) CreateInBatches(items []entity.Article) (err error) {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		for k := range items {
			if err := r.checkPost(tx, &items[k]); err != nil {
				return err
			}
			if err := tx.Create(&items[k].ArticleBase).Error; err != nil {
				return err
			}
			items[k].ArticleDetail.ArticleID = items[k].ArticleBase.ID
			if e := tx.Create(&items[k].ArticleDetail).Error; err != nil {
				return e
			}
		}
		return nil
	})
}

func (r *ArticleRepo) Update(item *entity.Article) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := r.checkPost(tx, item); err != nil {
			return err
		}
		if err := tx.Select("*").Omit("id").Where("id = ?", item.ArticleBase.ID).Updates(&item.ArticleBase).Error; err != nil {
			return err
		}
		item.ArticleDetail.ArticleID = item.ArticleBase.ID
		return tx.Select("*").Omit("article_id").Where("article_id = ?", item.ArticleDetail.ArticleID).Updates(&item.ArticleDetail).Error
	})
}

func (r *ArticleRepo) checkPost(tx *gorm.DB, item *entity.Article) error {
	var existID int
	if config.Config.More.UniqueTitle {
		if err := tx.Model(&entity.ArticleBase{}).Where("title = ? and id != ?", item.Title, item.ID).Limit(1).Pluck("id", &existID).Error; err != nil {
			return err
		}
		if existID > 0 {
			return message.ErrTitleAlreadyExists
		}
	}
	return nil
}

func (r *ArticleRepo) Delete(id int) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		return errorx.ErrorReturn(
			tx.Delete(&entity.ArticleBase{ID: id}).Error,
			tx.Delete(&entity.ArticleDetail{ArticleID: id}).Error,
		)
	})
}

func (r *ArticleRepo) Get(id int) (*entity.Article, error) {
	var res entity.Article
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.ArticleBase{}).Where("id = ?", id).Find(&res.ArticleBase).Error; err != nil || res.ArticleBase.ID == 0 {
			return err
		}
		return tx.Model(&entity.ArticleDetail{}).Where("article_id = ?", res.ArticleBase.ID).Find(&res.ArticleDetail).Error
	})
	return &res, err
}

// GetBySlug 通过slug获取
func (r *ArticleRepo) GetBySlug(slug string) (*entity.Article, error) {
	var res entity.Article
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.ArticleBase{}).Where("slug = ?", slug).Find(&res.ArticleBase).Error; err != nil || res.ArticleBase.ID == 0 {
			return err
		}
		return tx.Model(&entity.ArticleDetail{}).Where("article_id = ?", res.ArticleBase.ID).Find(&res.ArticleDetail).Error
	})
	return &res, err
}

// UpdateViewsBySlug 通过slug更新views
func (r *ArticleRepo) UpdateViewsBySlug(slug string, n int) error {
	return db.DB.Model(entity.ArticleBase{}).Where("slug = ?", slug).Limit(1).UpdateColumn("views", gorm.Expr("views + ?", n)).Error
}

// CountByWhere 通过where获取统计结果
func (r *ArticleRepo) CountByWhere(where *context.Where) (res int64, err error) {
	err = db.DB.Model(entity.ArticleBase{}).Scopes(gormx.Where(where)).Count(&res).Error
	return
}

// CountTotal 统计文章总数
func (r *ArticleRepo) CountTotal() (res int64, err error) {
	err = db.DB.Model(entity.ArticleBase{}).Count(&res).Error
	return
}

// CountToday 统计今日添加数量
func (r *ArticleRepo) CountToday() (res int64, err error) {
	err = db.DB.Model(entity.ArticleBase{}).Where("create_time >= ?", utils.TodayBeginTime().Unix()).Count(&res).Error
	return
}

// CountYesterday 统计昨日添加数量
func (r *ArticleRepo) CountYesterday() (res int64, err error) {
	today := utils.TodayBeginTime()
	yesterday := today.AddDate(0, 0, -1)
	err = db.DB.Model(entity.ArticleBase{}).Where("create_time >= ? and create_time < ?", yesterday.Unix(), today.Unix()).Count(&res).Error
	return
}

// CountLastFewDays 统计最近几日的数据
func (r *ArticleRepo) CountLastFewDays(n int) (res int64, err error) {
	today := utils.TodayBeginTime()
	days := today.AddDate(0, 0, -n)
	err = db.DB.Model(entity.ArticleBase{}).Where("create_time >= ?", days.Unix()).Count(&res).Error
	return
}

func (r *ArticleRepo) GetIdByTitle(title string) (id int, err error) {
	err = db.DB.Model(entity.ArticleBase{}).Where("title = ?", title).Limit(1).Pluck("id", &id).Error
	return
}

func (r *ArticleRepo) GetIdBySlug(slug string) (id int, err error) {
	err = db.DB.Model(entity.ArticleBase{}).Where("slug = ?", slug).Limit(1).Pluck("id", &id).Error
	return
}

// MaxID 获取最大ID
func (r *ArticleRepo) MaxID() (res int, err error) {
	err = db.DB.Model(entity.ArticleBase{}).Limit(1).Order("id desc").Limit(1).Pluck("id", &res).Error
	return
}

// List 调用文章列表
func (r *ArticleRepo) List(ctx *context.Context) (res []entity.ArticleBase, err error) {
	err = db.DB.Model(&entity.ArticleBase{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListExistThumbnail 调用文章列表
func (r *ArticleRepo) ListExistThumbnail(ctx *context.Context) (res []entity.ArticleBase, err error) {
	err = db.DB.Model(&entity.ArticleBase{}).Where("thumbnail !=''").Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByIds 根据id调用文章列表
func (r *ArticleRepo) ListByIds(ctx *context.Context, ids []int) (res []entity.ArticleBase, err error) {
	err = db.DB.Model(&entity.ArticleBase{}).Scopes(gormx.WhereIds(ids), gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByCategoryIds 根据分类ID调用文章列表
func (r *ArticleRepo) ListByCategoryIds(ctx *context.Context, categoryIds []int) (res []entity.ArticleBase, err error) {
	err = db.DB.Model(&entity.ArticleBase{}).Scopes(gormx.WhereCategoryIds(categoryIds), gormx.Context(ctx)).Find(&res).Error
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (r *ArticleRepo) ListAfterCreateTime(ctx *context.Context, t int64) (res []entity.ArticleBase, err error) {
	err = db.DB.Model(&entity.ArticleBase{}).Scopes(gormx.WhereCreateTimeAfter(t), gormx.Context(ctx)).Find(&res).Error
	return
}

// ListBeforeCreateTime 根据创建时间调用列表
func (r *ArticleRepo) ListBeforeCreateTime(ctx *context.Context, t int64) (res []entity.ArticleBase, err error) {
	err = db.DB.Model(&entity.ArticleBase{}).Scopes(gormx.WhereCreateTimeBefore(t), gormx.Context(ctx)).Find(&res).Error
	return
}

// ListDetail 调用详情表文章列表
func (r *ArticleRepo) ListDetail(ctx *context.Context) (res []entity.ArticleDetail, err error) {
	err = db.DB.Model(&entity.ArticleDetail{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

func (r *ArticleRepo) ListDetailByIds(ctx *context.Context, ids []int) (res []entity.ArticleDetail, err error) {
	err = db.DB.Model(&entity.ArticleDetail{}).Where("article_id in ?", ids).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// CountByCategoryID 模板分类ID统计
func (r *ArticleRepo) CountByCategoryID(id int) (res int64, err error) {
	err = db.DB.Model(entity.ArticleBase{}).Scopes(gormx.WhereCategoryID(id)).Count(&res).Error
	return
}

// BatchSetCategory 批量设置分类
func (r *ArticleRepo) BatchSetCategory(categoryID int, ids []int) error {
	return db.DB.Model(&entity.ArticleBase{}).Where("id in ?", ids).UpdateColumn("category_id", categoryID).Error
}
