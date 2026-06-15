up:
	docker-compose -p velonetics_otel up -d
.PHONY: up

down:
	docker-compose -p velonetics_otel down --volumes --remove-orphans
.PHONY: down

stop:
	docker-compose -p velonetics_otel stop
.PHONY: stop

local_up:
	docker-compose -p velonetics_otel_local -f compose.local.yaml up -d
.PHONY: up

local_down:
	docker-compose -p velonetics_otel_local -f compose.local.yaml down --volumes --remove-orphans
.PHONY: down

local_stop:
	docker-compose -p velonetics_otel_local -f compose.local.yaml stop
.PHONY: stop
