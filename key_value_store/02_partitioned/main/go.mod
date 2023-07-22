module github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/main

go 1.20

replace github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/common => ../common

replace github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/key_value_worker => ../key_value_worker

replace github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/web_server => ../web_server

require (
	github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/common v0.0.0-00010101000000-000000000000
	github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/key_value_worker v0.0.0-00010101000000-000000000000
	github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/web_server v0.0.0-00010101000000-000000000000
)
