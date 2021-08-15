Prometheus  	  – сбор и хранение метрик (http://:9090)

Grafana    		  - визуализация метри  (http://:3000)

AlertManger 	  - отправка сообщений  (http://:9093)

Caddy			  - обратный прокис-сервер с обеспечением Basic-аутентификации для Prometheus, AlertManager, Pushgateway

Node-exporter     – сборщик метрик c нод/хостов

Cadvisor 		  - сборщик метрик c контейнеров

Pushgateway       - сборщик метрик

____

Проверка Pushgateway:

echo "Temperature +20" | curl   -X POST -H  "Content-Type: text/plain" --data-binary @-  http://127.0.0.1:9091/metrics/job/temperature_metrics/instance/localhost
____

Проверка созданных правил:

docker exec -it 1111111 promtool  check rules /etc/prometheus/alert.rules
