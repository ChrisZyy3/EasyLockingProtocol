package initialize

// func OtherInit() {
// 	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
// 	if err != nil {
// 		panic(err)
// 	}

// 	global.BlackCache = local_cache.NewCache(
// 		local_cache.SetDefaultExpire(dr),
// 	)
// 	file, err := os.Open("go.mod")
// 	if err == nil && global.GVA_CONFIG.AutoCode.Module == "" {
// 		scanner := bufio.NewScanner(file)
// 		scanner.Scan()
// 		global.GVA_CONFIG.AutoCode.Module = strings.TrimPrefix(scanner.Text(), "module ")
// 	}
// }
