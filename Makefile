.PHONY: protogen unbindPort3000

protogen:
	@sh ./scripts/protoc-gen.sh

unbindPort9090:
	kill -9 $$(lsof -i :9090 | grep main | awk '{ print $$2}' | xargs) ||: