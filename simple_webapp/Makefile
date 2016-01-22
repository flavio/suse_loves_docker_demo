all :: sles11sp3 sles11sp4 sles12 sles12sp1
	@echo 'DONE'

webapp ::
	cd webapp_demo; make

sles12 :: webapp
	@echo Building SLE12 demo image
	echo docker build -f Dockerfile-sles12 -t sles12_webapp .

sles12sp1 :: webapp
	@echo Building SLE12SP1 demo image
	docker build -f Dockerfile-sles12sp1 -t sles12sp1_webapp .

sles11sp3 :: webapp
	@echo Building SLE11SP3 demo image
	docker build -f Dockerfile-sles11sp3 -t sles11sp3_webapp .

sles11sp4 :: webapp
	@echo Building SLE11SP4 demo image
	docker build -f Dockerfile-sles11sp4 -t sles11sp4_webapp .

clean ::
	rm webapp_demo/webapp_demo
	docker rmi sles11sp3_webapp
	docker rmi sles11sp4_webapp
	docker rmi sles12_webapp
	docker rmi sles12sp1_webapp
