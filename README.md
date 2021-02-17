# enterprise-wechat-push-restful-api

send wechat notification via http restful api

## how to use?

1. Configure according to config.yaml
2. Run the program through `e-wpra -c config.yaml`
3. Use nginx to reverse the 8090 port to achieve diversion

## api

1. ../verify wechat verify handler
2. ../send send wechat notification handler (post `title=%s&body=%s&touser=%s`)