package service

import (
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/cache"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/jinzhu/gorm"
)

type groupService struct {
}

func (g *groupService) CreateGroup(req *pb.CreateGroupReq) (uint64, error) {
	grp := &dao.Group{
		GroupName:    req.Group.GroupName,
		CreateUser:   req.Group.CreateUser,
		Owner:        req.Group.Owner,
		Announcement: req.Group.Announcement,
		UserNum:      0,
		IsDeleted:    0,
	}
	var (
		groupId uint64
		err     error
	)
	db := base.Database()
	err = db.Transaction(func(tx *gorm.DB) error {
		groupDao := &dao.GroupDao{DB: tx}
		groupId, err = groupDao.InsertOne(grp)
		if err != nil {
			return err
		}

		// 创建完群组后要把创建者加入群组
		groupUserDao := &dao.GroupUserDao{DB: tx}
		groupUser := &dao.GroupUser{
			GroupId:   groupId,
			Username:  grp.CreateUser,
			IsDeleted: 0,
			CreatedAt: nil,
			UpdatedAt: nil,
		}
		err = groupUserDao.InsertOne(groupUser)
		if err != nil {
			return err
		}

		err = groupDao.UpdateNum(groupId, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return groupId, nil
}

func (g *groupService) DeleteGroup(groupId uint64) error {
	db := base.Database()
	users, err := g.ListGroupMember(groupId)
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		groupUserDao := &dao.GroupUserDao{DB: tx}
		// 解散群组的同时要把群员移出该群
		for _, v := range users {
			v.IsDeleted = 1
			err = groupUserDao.UpdateOne(v)
			if err != nil {
				return err
			}
		}

		groupDao := &dao.GroupDao{DB: tx}
		err = groupDao.DeleteOne(groupId)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	key := cache.GroupUserCache.Key(groupId)
	err = cache.GroupUserCache.Del(key)
	if err != nil {
		return err
	}

	return nil
}

func (g *groupService) AddGroupMember(req *pb.AddGroupMemberReq) error {
	db := base.Database()
	err := db.Transaction(func(tx *gorm.DB) error {
		groupUserDao := &dao.GroupUserDao{DB: tx}
		groupUser := &dao.GroupUser{
			GroupId:  req.GroupId,
			Username: req.Username,
		}
		err := groupUserDao.InsertOne(groupUser)
		if err != nil {
			if err == dao.DAO_ERROR_DUPLICATE_RECORD {
				// 数据库中已有该记录
				groupUser.IsDeleted = 0
				err = groupUserDao.UpdateOne(groupUser)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		groupDao := &dao.GroupDao{DB: tx}
		num, err := groupDao.UserNum(req.GroupId)
		if err != nil {
			return err
		}
		num++
		err = groupDao.UpdateNum(req.GroupId, num)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	key := cache.GroupUserCache.Key(req.GroupId)
	err = cache.GroupUserCache.Del(key)
	if err != nil {
		return err
	}

	return nil
}

func (g *groupService) DelGroupMember(req *pb.DelGroupMemberReq) error {
	db := base.Database()
	err := db.Transaction(func(tx *gorm.DB) error {
		groupUserDao := &dao.GroupUserDao{DB: tx}
		groupUser := &dao.GroupUser{
			GroupId:   req.GroupId,
			Username:  req.Username,
			IsDeleted: 1,
		}
		err := groupUserDao.UpdateOne(groupUser)
		if err != nil {
			return err
		}

		groupDao := &dao.GroupDao{DB: tx}
		num, err := groupDao.UserNum(req.GroupId)
		if err != nil {
			return err
		}

		num--
		if num == 0 {
			// 如果群组已经没人
			err = groupDao.DeleteOne(req.GroupId)
			if err != nil {
				return err
			}
		} else {
			err = groupDao.UpdateNum(req.GroupId, num)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	key := cache.GroupUserCache.Key(req.GroupId)
	err = cache.GroupUserCache.Del(key)
	if err != nil {
		return err
	}

	return nil
}

func (g *groupService) ListGroupMember(groupId uint64) ([]*dao.GroupUser, error) {
	key := cache.GroupUserCache.Key(groupId)
	u, err := cache.GroupUserCache.Get(key)
	if err == nil && u != nil {
		return u, nil
	}

	db := base.Database()
	groupUserDao := &dao.GroupUserDao{DB: db}
	users, err := groupUserDao.GetAll(groupId)
	if err != nil {
		return nil, err
	}

	// 存储从数据库中取到的数据到缓存中，下次直接从缓存中取
	cache.GroupUserCache.Set(key, users)

	return users, nil
}

func (g *groupService) IsMember(groupId uint64, username string) (bool, error) {
	db := base.Database()
	groupUserDao := &dao.GroupUserDao{DB: db}
	groupUser, err := groupUserDao.GetOne(groupId, username)
	if err != nil {
		return false, err
	}

	if groupUser.IsDeleted == 1 {
		return false, nil
	}
	return true, nil
}
