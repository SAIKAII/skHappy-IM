package service

import (
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
)

type relationshipService struct {
}

// Greater 比较UserA是否应该排在UserB前
func (r *relationshipService) Greater(userA, userB string) bool {
	return userA >= userB
}

func (r *relationshipService) CreateRelationship(userA, userB string) error {
	rel := &dao.Relationship{}
	// UserA应该权重应该比UserB大，否则交换位置
	if r.Greater(userA, userB) {
		rel.UserA, rel.UserB = userA, userB
	} else {
		rel.UserA, rel.UserB = userB, userA
	}

	res, err := r.get(rel.UserA, rel.UserB)
	if err != nil && err != dao.DAO_ERROR_RECORD_NOT_FOUND {
		base.Logger.Errorln(err)
		return err
	}
	db := base.Database()
	dao := dao.RelationShipDao{DB: db}
	if res != nil {
		// 表中之前有创建过记录
		rel.IsDeleted = 0
		err = dao.Update(rel)
	} else {
		// 表中没有相关记录
		err = dao.Insert(rel)
	}
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (r *relationshipService) DeleteRelationship(userA, userB string) error {
	// UserA应该权重应该比UserB大，否则交换位置
	if !r.Greater(userA, userB) {
		userA, userB = userB, userA
	}

	db := base.Database()
	dao := dao.RelationShipDao{DB: db}
	err := dao.Delete(userA, userB)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (r *relationshipService) GetAllFriends(username string) ([]string, error) {
	db := base.Database()
	relDao := dao.RelationShipDao{DB: db}
	rels, err := relDao.GetAll(username)
	if err != nil {
		base.Logger.Errorln(err)
		return nil, err
	}

	res := make([]string, len(rels))
	for i, r := range rels {
		if r.UserA == username {
			res[i] = r.UserB
		} else {
			res[i] = r.UserA
		}
	}

	return res, nil
}

func (r *relationshipService) IsFriend(userA, userB string) (bool, error) {
	if !r.Greater(userA, userB) {
		userA, userB = userB, userA
	}

	db := base.Database()
	relDao := dao.RelationShipDao{DB: db}
	rel, err := relDao.GetOne(userA, userB)
	if err != nil {
		base.Logger.Errorln(err)
		return false, err
	}

	if rel.IsDeleted == 1 {
		base.Logger.Errorln("两人不是好友")
		return false, nil
	}

	return true, nil
}

func (r *relationshipService) get(userA, userB string) (*dao.Relationship, error) {
	db := base.Database()
	relDao := dao.RelationShipDao{DB: db}
	res, err := relDao.GetOne(userA, userB)
	if err != nil {
		base.Logger.Errorln(err)
		return nil, err
	}

	return res, nil
}
