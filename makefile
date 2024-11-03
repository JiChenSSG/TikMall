.PHONY: gen
gen: ## gen client code of {svc}. example: make gen svc=product
	@scripts/gen.sh ${svc}

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=product
	@mkdir -p rpc_gen && cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/cloudwego/biz-demo/gomall/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: ## gen service code of {svc}. example: make gen-server svc=product
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/cloudwego/biz-demo/gomall/app/${svc} --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/checkout_page.proto
