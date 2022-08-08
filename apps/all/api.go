package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github/tqtcloud/cmdb/apps/book/api"
	_ "github/tqtcloud/cmdb/apps/host/api"
	_ "github/tqtcloud/cmdb/apps/resource/api"
	_ "github/tqtcloud/cmdb/apps/secret/api"
	_ "github/tqtcloud/cmdb/apps/task/api"
)
