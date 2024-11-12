// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goshop/internal/dao/internal"
)

// internalCommentInfoDao is internal type for wrapping internal DAO implements.
type internalCommentInfoDao = *internal.CommentInfoDao

// commentInfoDao is the data access object for table comment_info.
// You can define custom methods on it to extend its functionality as you wish.
type commentInfoDao struct {
	internalCommentInfoDao
}

var (
	// CommentInfo is globally public accessible object for table comment_info operations.
	CommentInfo = commentInfoDao{
		internal.NewCommentInfoDao(),
	}
)

// Fill with you ideas below.
