// 代码生成时间: 2025-09-21 12:04:33
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
    "github.com/markbates/willie"
    "github.com/markbates/willie/environment"
    "github.com/markbates/willie/middleware"
    "go.uber.org/zap"
    "time"
)

// CacheService 是缓存服务的接口
type CacheService interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{}, duration time.Duration)
    MustGet(key string) interface{}
    Invalidate(key string)
}

// cacheService 是 CacheService 接口的实现
type cacheService struct {
    *willie.Willie
}

// NewCacheService 创建一个新的缓存服务
func NewCacheService() CacheService {
    willie := willie.New()
    willie.WithOptions(willie.Options{
        Environment: environment.Get(),
        Logger:      zap.NewNop().Sugar(), // 使用无操作日志记录器
    })
    return &cacheService{
        Willie: willie,
    }
}

// Get 从缓存中获取值
func (c *cacheService) Get(key string) (interface{}, bool) {
    value, err := c.Willie.Get(key)
    if err != nil {
        // 处理错误，可能记录日志
        return nil, false
    }
    return value, true
}

// Set 将值设置到缓存
func (c *cacheService) Set(key string, value interface{}, duration time.Duration) {
    if err := c.Willie.Set(key, value, duration); err != nil {
        // 处理错误，可能记录日志
    }
}

// MustGet 必须从缓存中获取值，如果值不存在则返回错误
func (c *cacheService) MustGet(key string) interface{} {
    value, err := c.Willie.Get(key)
    if err != nil {
        panic(err) // 或者根据需要处理错误
    }
    return value
}

// Invalidate 从缓存中使值失效
func (c *cacheService) Invalidate(key string) {
    c.Willie.Del(key)
}

// main函数，Buffalo应用的入口点
func main() {
    app := buffalo.Automatic()

    // 设置缓存中间件
    app.Use(middleware.NewLogger(zap.NewNop().Sugar()))
    app.Use(popmw.Transactioner{})

    // 创建缓存服务
    cache := NewCacheService()

    // 定义路由和处理函数，例如：
    app.GET("/cache/:key", func(c buffalo.Context) error {
        key := c.Param("key")
        value, ok := cache.Get(key)
        if !ok {
            return buffalo.NewError("Value not found in cache", 404)
        }
        return c.Render(200, buffalo.JSON(map[string]interface{}{"value": value}))
    })

    // 启动应用
    app.Serve()
}
