package cache

import (
	"eee/pkg/utils"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

func SendEmailCode(uname, email string) (err error) {
	//判断过期时间内是否已发送过
	stringCmd := RedisCache.Get(uname)
	if stringCmd.Err() != redis.Nil {
		return errors.New("已发送过")
	}
	// 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	//在5min内未发送过则设置kv
	cmd := RedisCache.Set(uname, vcode, 1*time.Minute)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	if err := utils.SendMail(email, vcode); err != nil {
		return err
	}
	return nil
}

func ValidateEmailCode(uname, code string) (err error) {
	stringCmd := RedisCache.Get(uname)
	if stringCmd.Err() != nil {
		if stringCmd.Err().Error() == redis.Nil.Error() {
			return errors.New("code has expired")
		} else {
			return stringCmd.Err()
		}
	}
	tcode := stringCmd.String()
	fmt.Println(strings.Split(tcode, " ")[2])
	if strings.Split(tcode, " ")[2] != code {
		return errors.New("wrong code")
	}
	return nil
}
