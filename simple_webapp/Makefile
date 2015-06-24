all :: sles12 sles11sp3
	echo 'DONE'

webapp ::
	cd webapp_demo && go build

sles12 :: webapp
	ln -f webapp_demo/webapp_demo demo-sles12/webapp_demo
	cp -r webapp_demo/web demo-sles12/web
	docker build -t sles12_webapp demo-sles12

sles11sp3 :: webapp
	ln -f webapp_demo/webapp_demo demo-sles11sp3/webapp_demo
	cp -r webapp_demo/web demo-sles11sp3/web
	docker build -t sles11sp3_webapp demo-sles11sp3

clean ::
	rm -rf demo-sles11sp3/web
	rm -rf demo-sles11sp3/webapp_demo
	rm -rf demo-sles12/web
	rm -rf demo-sles12/webapp_demo
	docker rmi sles11sp3_webapp
	docker rmi sles12_webapp
