package all

import (
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "github/tqtcloud/cmdb/apps/book/impl"
	_ "github/tqtcloud/cmdb/apps/host/impl"
	_ "github/tqtcloud/cmdb/apps/resource/impl"
	_ "github/tqtcloud/cmdb/apps/secret/impl"
	_ "github/tqtcloud/cmdb/apps/task/impl"
)
