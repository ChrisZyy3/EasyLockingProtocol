package internal

import (
	"os"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
}

// NewZapCore 创建并初始化一个ZapCore实例。
// 参数:
//
//	level zapcore.Level - 日志级别，只有等于该级别的日志将会被处理。
//
// 返回值:
//
//	*ZapCore - 初始化后的ZapCore实例指针。
func NewZapCore(level zapcore.Level) *ZapCore {
	// 创建ZapCore实例并设置其级别。
	entity := &ZapCore{level: level}

	// 获取并设置实体的WriteSyncer，这是为了确保日志消息可以被写入和同步。
	syncer := entity.WriteSyncer()

	// 定义一个仅允许特定级别日志通过的LevelEnabler。
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})

	// 使用全局配置、WriteSyncer和LevelEnabler来初始化ZapCore的Core。
	entity.Core = zapcore.NewCore(global.GVA_CONFIG.Zap.Encoder(), syncer, levelEnabler)

	// 返回初始化后的ZapCore实例。
	return entity
}

func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		global.GVA_CONFIG.Zap.Director,
		z.level.String(),
		global.GVA_CONFIG.Zap.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	if global.GVA_CONFIG.Zap.LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiSyncer)
	}
	return zapcore.AddSync(cutter)
}

func (z *ZapCore) Enabled(level zapcore.Level) bool {
	return z.level == level
}

func (z *ZapCore) With(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *ZapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *ZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	for i := 0; i < len(fields); i++ {
		if fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory" {
			syncer := z.WriteSyncer(fields[i].String)
			z.Core = zapcore.NewCore(global.GVA_CONFIG.Zap.Encoder(), syncer, z.level)
		}
	}
	return z.Core.Write(entry, fields)
}

func (z *ZapCore) Sync() error {
	return z.Core.Sync()
}
