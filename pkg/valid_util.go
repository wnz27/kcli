// Package pkg
package pkg

/**
 * @project ads-serving
 * @Author 27
 * @Description //TODO
 * @Date 2021/11/20 11:17 11月
 **/

// ValidServiceName TODO 这里支持服务 可以 从注册的获得
func ValidServiceName(serviceName string) {
	if len(serviceName) == 0 ||
		(serviceName != "control" &&
			serviceName != "delivery" &&
			serviceName != "collector" &&
			serviceName != "dsp-caller") {
		panic("[not support this service: " + serviceName + "]")
	}
}
