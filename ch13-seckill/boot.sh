cd $HOME/Downloads/kafka_2.12-3.5.1
./bin/zookeeper-server-start.sh config/zookeeper.properties &

brew services start etcd
docker run --name zipkin -d -p 9411:9411 openzipkin/zipkin

cd $HOME/study/java-config-server
/usr/bin/env /Library/Java/JavaVirtualMachines/jdk-17.jdk/Contents/Home/bin/java @/var/folders/8_/nwjm22012nz8hbr17tpy7qpr0000gn/T/cp_7tmwlrbna27t9va6nihwv6u6j.argfile com.qiushye.configserver.ConfigServerApplication &

cd $HOME/study/goCorrentMS/ch13-seckill
go run user-service/main.go &
go run gateway/main.go &
go run sk-admin/main.go &
go run sk-app/main.go &
go run sk-core/main.go &

cd $HOME/Downloads/apache-jmeter-5.6.2/bin
sh jmeter