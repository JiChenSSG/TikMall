# 订单文档

## 定时取消实现

在每次调用get相关方法时，会检查订单是否超时，如果超时则取消订单，取消订单的逻辑在`app/order/service/order.go`中的`CancelOrder`方法中实现。

订单超时的时间可在配置文件中配置。