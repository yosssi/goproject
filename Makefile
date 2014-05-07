ES_INDEX=http://localhost:9200/goproject
esdelete:
	curl -XDELETE $(ES_INDEX)
escreate:
	curl -XPUT $(ES_INDEX) -d @elasticsearch/goproject.json
esinit:
	make esdelete
	make escreate
