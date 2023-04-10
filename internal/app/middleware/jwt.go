package middleware

import (
	"RunnerGo-management/internal/pkg/biz/errno"
	"RunnerGo-management/internal/pkg/biz/jwt"
	"RunnerGo-management/internal/pkg/dal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{"code": errno.ErrMustLogin, "message": "must login"})
			c.Abort()
			return
		}

		userID, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": errno.ErrMustLogin, "message": "must login"})
			c.Abort()
			return
		} else if userID == "" {
			c.JSON(http.StatusOK, gin.H{"code": errno.ErrMustLogin, "message": "must login"})
			c.Abort()
			return
		}

		// 查询token里面的用户信息是否存在于数据库
		userTable := dal.GetQuery().User
		_, err = userTable.WithContext(c).Where(userTable.UserID.Eq(userID)).First()
		if err != nil {
			// 把token设置为过期
			_, _, err := jwt.GenerateTokenByTime(userID, 0)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": errno.ErrInvalidToken, "message": "must login"})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, gin.H{"code": errno.ErrMustLogin, "message": "must login"})
			c.Abort()
			return
		}

		// 过滤部分接口的校验条件
		apiPath := c.Request.URL.Path
		//  /management/api/v1/setting/get
		if apiPath != "/management/api/v1/setting/get" {
			// 校验用户默认团队是否正确
			tx := dal.GetQuery().Setting
			teamIDString := c.GetHeader("CurrentTeamID")
			//teamIDInt, err := strconv.ParseInt(teamIDString, 10, 64)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": errno.ErrDefaultTeamFailed, "message": "default team failed--CurrentTeamID类型转换失败"})
				c.Abort()
				return
			}

			if teamIDString == "" { // 默认团队是0，则跳转到去续费
				if apiPath != "/management/api/v1/order/get_team_buy_version_new" &&
					apiPath != "/management/api/v1/order/get_team_trial_expiration_date_new" &&
					apiPath != "/management/api/v1/order/get_team_bug_version_amount_new" &&
					apiPath != "/management/api/v1/order/create_order_new" &&
					apiPath != "/management/api/v1/order/get_order_pay_status_new" &&
					apiPath != "/management/api/v1/order/get_order_pay_detail_new" &&
					apiPath != "/management/api/v1/team/list_new" &&
					apiPath != "/management/api/v1/order/batch_delete_order_new" {
					c.JSON(http.StatusOK, gin.H{"code": errno.ErrNotFundAvailTeam, "message": "没有可用团队"})
					c.Abort()
					return
				}
			}

			_, err = tx.WithContext(c).Where(tx.UserID.Eq(userID)).Where(tx.TeamID.Eq(teamIDString)).First()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": errno.ErrDefaultTeamFailed, "message": "default team failed--查询的默认团队错误"})
				c.Abort()
				return
			}
		}

		c.Set("user_id", userID)

		c.Next()
	}
}
